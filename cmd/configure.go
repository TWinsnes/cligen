package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/twinsnes/cligen/internal/config"

	"github.com/manifoldco/promptui"
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
			return configurePrompt()
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

func configurePrompt() error {

	conf, err := config.LoadConfig()

	if err != nil {
		return err
	}

	enabled, err := promptForHomebrewEnabled()

	if err != nil {
		return err
	}

	conf.HomebrewConfig.Enabled = enabled

	username, err := promtForGithubUsername(conf.HomebrewConfig.GithubUsername)
	if err != nil {
		return err
	}
	conf.HomebrewConfig.GithubUsername = username

	repo, err := promptForHomebrewRepo(conf.HomebrewConfig.Repo)
	if err != nil {
		return err
	}
	conf.HomebrewConfig.Repo = repo

	err = conf.SaveConfig()
	if err != nil {
		return err
	}

	path, _ := config.GetConfigPath()

	fmt.Println("Configuration saved to: ", path)
	return nil
}

func promptForHomebrewEnabled() (bool, error) {
	prompt := promptui.Prompt{
		Label:     "Enable Homebrew tap?",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrAbort) {
			return false, nil
		}
		fmt.Println("Prompt failed:", err)
		return false, err
	}

	return true, nil
}

func promtForGithubUsername(def string) (string, error) {
	prompt := promptui.Prompt{
		Label:   "Homebrew Github username",
		Default: def,
	}

	return prompt.Run()
}

func promptForHomebrewRepo(def string) (string, error) {
	prompt := promptui.Prompt{
		Label:   "Homebrew repo name",
		Default: def,
	}

	return prompt.Run()
}
