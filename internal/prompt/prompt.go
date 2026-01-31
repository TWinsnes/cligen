package prompt

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/twinsnes/cligen/internal/config"
	"github.com/twinsnes/cligen/internal/gen"
	"github.com/twinsnes/cligen/internal/git"
)

type UserPrompt struct {
	conf              *config.Config
	AvailableFeatures map[string]FeaturePrompt
}

type FeaturePrompt interface {
	RunPrompt() error
	IsDefaultSelected() bool

	gen.Feature
}

func NewUserPrompt(conf *config.Config) *UserPrompt {

	homebrewFeature := NewHomebrewFeature(conf)
	docsFeature := NewDocsFeature(conf)

	return &UserPrompt{
		conf: conf,
		AvailableFeatures: map[string]FeaturePrompt{
			homebrewFeature.GetName(): homebrewFeature,
			docsFeature.GetName():     docsFeature,
		},
	}
}

func (p *UserPrompt) Run() (gen.TemplateOptions, error) {

	options, err := getBaseSettings(p.conf)

	if err != nil {
		return gen.TemplateOptions{}, err
	}

	selectedFeatures, err := runSelectFeatures(p.AvailableFeatures)

	if err != nil {
		return gen.TemplateOptions{}, err
	}

	var features []gen.Feature
	for _, feature := range selectedFeatures {
		err := feature.RunPrompt()
		if err != nil {
			return gen.TemplateOptions{}, err
		}

		features = append(features, feature)
	}

	options.Features = features

	return options, nil
}

func getBaseSettings(conf *config.Config) (gen.TemplateOptions, error) {
	var appName string
	var golangVersion string
	var goModuleName string
	var templateType string

	defaultAppName := getDefaultAppName()
	defaultModule := getDefaultModule()
	golangVersion = getDefaultGoVersion()

	templateTypeOptions, err := getTemplateTypeOptions()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	goVersionOptions := getGolangVersionOptions(golangVersion)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("App name").
				Description("The name of the CLI app.").
				Placeholder(defaultAppName).
				Validate(validateAppName).
				Value(&appName),
			huh.NewInput().
				Title("Go Module name").
				Description("The name of the Go module.").
				Placeholder(defaultModule).
				Value(&goModuleName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Golang version").
				Description("The version of Go to use.").
				Options(goVersionOptions...).
				Value(&golangVersion),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Template").
				Description("The template to generate the app from.").
				Options(templateTypeOptions...).
				Value(&templateType),
		),
	)

	err = form.Run()
	if err != nil {
		return gen.TemplateOptions{}, err
	}

	if appName == "" {
		appName = defaultAppName
	}

	if goModuleName == "" {
		goModuleName = defaultModule
	}

	return gen.TemplateOptions{
		AppName:       appName,
		GolangVersion: golangVersion,
		ModuleName:    goModuleName,
		TemplateType:  templateType,
		DryRun:        false,
	}, nil
}

func getGolangVersionOptions(golangVersion string) []huh.Option[string] {
	goVersions := []string{"1.25", "1.24", "1.23"}
	hasDefault := false
	for _, v := range goVersions {
		if v == golangVersion {
			hasDefault = true
			break
		}
	}

	if !hasDefault {
		goVersions = append(goVersions, golangVersion)
	}

	slices.SortFunc(goVersions, func(a, b string) int {
		return strings.Compare(b, a)
	})

	goVersionOptions := make([]huh.Option[string], len(goVersions))
	for i, v := range goVersions {
		goVersionOptions[i] = huh.NewOption(v, v)
	}
	return goVersionOptions
}

func validateAppName(appName string) error {
	matched, err := regexp.Match("^[a-zA-Z0-9_-]*$", []byte(appName))

	if err != nil {
		slog.Debug("error matching app name",
			slog.Any("error", err),
		)

		return fmt.Errorf("internal error when validating app name, enable debug logging for more details")
	}

	if !matched {
		return fmt.Errorf("app name must only contain letters, numbers, hyphens or underscores")
	}
	return nil
}

func getDefaultAppName() string {
	cwd, err := os.Getwd()
	if err == nil {
		return filepath.Base(cwd)
	}
	return "cligen"
}

func getDefaultModule() string {
	m := git.ModuleFromGit()

	if m == "" {
		m = "github.com/twinsnes/cligen"
	}

	return m
}

func getTemplateTypeOptions() ([]huh.Option[string], error) {
	templates, err := gen.ListTemplates()
	if err != nil {
		return []huh.Option[string]{}, err
	}
	var options []huh.Option[string]
	for _, template := range templates {
		options = append(options, huh.Option[string]{Key: template, Value: template})
	}
	return options, nil
}

func getDefaultGoVersion() string {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		return "1.25"
	}
	// Output is usually: "go version go1.25.1 darwin/arm64"
	fields := strings.Fields(string(out))
	if len(fields) < 3 {
		return "1.25"
	}
	version := strings.TrimPrefix(fields[2], "go")
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return "1.25"
}
