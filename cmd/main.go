package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func Run() {
	app := &cli.Command{
		Name:    "cligen",
		Usage:   "Generate a new Go CLI app scaffold powered by urfave/cli",
		Version: "0.1.0",
		Commands: []*cli.Command{
			newCmd(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
