package cmd

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/twinsnes/cligen/internal/config"

	"github.com/urfave/cli/v3"
)

func configureCmd() *cli.Command {
	return &cli.Command{
		Name:  "configure",
		Usage: "Configure default options for new CLI apps.",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			configureFileLocationCmd(),
		},
		Action: func(c context.Context, cmd *cli.Command) error {
			return runConfigurePrompt()
		},
	}
}

func configureFileLocationCmd() *cli.Command {
	return &cli.Command{
		Name:  "location",
		Usage: "Display the location of the configuration file.",
		Action: func(c context.Context, cmd *cli.Command) error {
			path, err := config.GetConfigPath()
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		},
	}
}

func runConfigurePrompt() error {

	conf, err := config.LoadConfig()

	if err != nil {
		return err
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Homebrew").
				Description("Enable Homebrew tap?").
				Affirmative("yes").
				Negative("no").
				Value(&conf.HomebrewConfig.Enabled),
			huh.NewInput().
				Title("Homebrew Github username").
				Description("The username of the homebrew repo owner.").
				Value(&conf.HomebrewConfig.GithubUsername),
			huh.NewInput().
				Title("Homebrew repo name").
				Description("The name of the homebrew repo.").
				Value(&conf.HomebrewConfig.Repo),
		),
	)

	err = form.Run()

	if err != nil {
		return err
	}

	err = conf.SaveConfig()
	if err != nil {
		return err
	}

	path, _ := config.GetConfigPath()

	fmt.Println("Configuration saved to: ", path)
	return nil
}
