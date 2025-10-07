package cmd

import (
	"context"

	"github.com/twinsnes/cligen/internal/config"
	"github.com/twinsnes/cligen/internal/prompt"

	"github.com/urfave/cli/v3"
)

func newCmd() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Create a new CLI app in the current directory",
		Flags: []cli.Flag{},
		Action: func(c context.Context, cmd *cli.Command) error {

			conf, err := config.LoadConfig()

			if err != nil {
				return err
			}

			userPrompt := prompt.NewUserPrompt(conf)

			templateOptions, err := userPrompt.Run()

			if err != nil {
				return err
			}

			err = templateOptions.RenderTemplate()
			if err != nil {
				return err
			}
			return nil
		},
	}
}
