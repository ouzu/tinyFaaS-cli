package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tinyFaaS-cli",
		Usage: "a tinyFaaS command line client",
		Commands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"l"},
				Usage:     "list all functions",
				Action:    listCommand,
				ArgsUsage: "",
			},
			{
				Name:      "delete",
				Aliases:   []string{"d"},
				Usage:     "delete a function",
				Action:    deleteCommand,
				ArgsUsage: "[function]",
			},
			{
				Name:      "upload",
				Aliases:   []string{"u"},
				Usage:     "upload a function",
				Action:    uploadCommand,
				ArgsUsage: "[directory] [name] [threads]",
			},
			{
				Name:      "wipe",
				Action:    wipeCommand,
				Usage:     "wipe all functions",
				ArgsUsage: "",
			},
			{
				Name:      "logs",
				Usage:     "print function logs",
				Action:    logsCommand,
				ArgsUsage: "",
			},
			{
				Name:      "run",
				Aliases:   []string{"r"},
				Usage:     "run a function",
				Action:    runCommand,
				ArgsUsage: "[function] [arguments...]",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "config",
				Usage:     "toml file to read configuration from",
				TakesFile: true,
			},
		},
		EnableBashCompletion: true,
		Authors:              []*cli.Author{{Name: "Marek Wallich"}},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
