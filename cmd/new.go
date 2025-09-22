package cmd

import (
	"context"

	"github.com/manifoldco/promptui"
	"github.com/twinsnes/cligen/internal/gen"
	"github.com/urfave/cli/v3"
)

func newCmd() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Create a new CLI app in the specified directory",
		ArgsUsage: "[path]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "force", Aliases: []string{"f"}, Usage: "Overwrite existing files if they already exist"},
		},
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
	prompt := promptui.Select{
		Label: "Select GoLang version",
		Items: []string{
			"1.25",
			"1.24",
			"1.23",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForAppName() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Enter App Name",
		Default: "mycli",
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForModuleName() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Enter Module Name",
		Default: "github.com/twinsnes/cligen",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}
