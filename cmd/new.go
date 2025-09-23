package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/twinsnes/cligen/internal/gen"
	"github.com/twinsnes/cligen/internal/git"
	"github.com/urfave/cli/v3"
)

func newCmd() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Create a new CLI app in the current directory",
		Flags: []cli.Flag{},
		Action: func(c context.Context, cmd *cli.Command) error {
			templateOptions, err := promptForOptions()

			if err != nil {
				return err
			}

			err = gen.RenderTemplate(templateOptions)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func promptForOptions() (gen.TemplateOptions, error) {
	templateOptions := gen.TemplateOptions{
		OutputPathPrefix: ".",
	}

	result, err := promptForGolangVersion()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	templateOptions.GolangVersion = result

	template, err := promptForTemplate()

	if err != nil {
		return gen.TemplateOptions{}, err
	}

	templateOptions.TemplateType = template

	appName, err := promptForAppName()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	templateOptions.AppName = appName

	module, err := promptForModuleName()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	templateOptions.ModuleName = module

	return templateOptions, nil
}

func promptForTemplate() (string, error) {
	templates, err := gen.GetTemplates()

	if err != nil {
		return "", err
	}

	prompt := promptui.Select{
		Label: "Select template",
		Items: templates,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil

}

func promptForGolangVersion() (string, error) {
	items := []string{
		"1.25",
		"1.24",
		"1.23",
	}

	prompt := promptui.Select{
		Label: "Select GoLang version",
		Items: items,
	}

	want := getCurrentEnvGoMinor()
	start := 0
	for i, v := range items {
		if v == want {
			start = i
			break
		}
	}

	_, result, err := prompt.RunCursorAt(start, 0)

	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForAppName() (string, error) {

	cliName := "mycli"

	cwd, err := os.Getwd()
	if err == nil {
		cliName = filepath.Base(cwd)
	}

	prompt := promptui.Prompt{
		Label:   "Enter App Name",
		Default: cliName,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForModuleName() (string, error) {
	def := getDefaultModuleName()

	prompt := promptui.Prompt{
		Label:   "Enter Module Name",
		Default: def,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func getDefaultModuleName() string {
	m := git.ModuleFromGit()

	if m == "" {
		m = "github.com/twinsnes/cligen"
	}

	return m
}

func getCurrentEnvGoMinor() string {
	// Try to get the toolchain version
	out, err := exec.Command("go", "env", "GOVERSION").Output()
	if err == nil {
		return normalizeMajorMinor(strings.TrimSpace(string(out)))
	}
	// Fallback to the build version
	return normalizeMajorMinor(runtime.Version())
}

func normalizeMajorMinor(goVersion string) string {
	// Accept inputs like "go1.23.1", "1.23.1", or "devel go1.24-abcdef"
	s := strings.TrimSpace(goVersion)

	// If it contains spaces (like devel output), take the first token that starts with "go"
	fields := strings.Fields(s)
	for _, f := range fields {
		if strings.HasPrefix(f, "go") {
			s = f
			break
		}
	}

	s = strings.TrimPrefix(s, "go")
	if i := strings.IndexByte(s, ' '); i >= 0 { // just in case
		s = s[:i]
	}
	parts := strings.Split(s, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1] // major.minor
	}
	return s
}
