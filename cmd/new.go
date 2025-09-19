package cmd

import (
	"context"
	"errors"
	"fmt"

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

			fmt.Printf("%q\n", templateOptions)

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
		OutputPathPrefix: "tmp",
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
		fmt.Printf("Prompt failed %v\n", err)
		return "", errors.New("prompt failed")
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
		fmt.Printf("Prompt failed %v\n", err)
		return "", errors.New("prompt failed")
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
