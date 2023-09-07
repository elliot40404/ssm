package main

import (
	"log"
	"os"

	"github.com/elliot40404/ssm/pkg/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "ssm",
		Usage:   "simple ssh manager",
		Version: "0.2.0",
		Authors: []*cli.Author{
			{
				Name:  "Elliot",
				Email: "admin@elliot404.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "filter",
						Aliases: []string{"f"},
						Usage:   "filter ssh configs",
					},
				},
				Usage:  "list available ssh configs",
				Action: commands.List,
			},
			{
				Name:   "add",
				Usage:  "add new ssh config",
				Action: commands.Add,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "dry-run",
						Aliases: []string{"d"},
						Usage:   "dry run",
					},
				},
			},
			{
				Name:      "print",
				Usage:     "print ssh config",
				Action:    commands.Print,
				ArgsUsage: "<config name>",
			},
			{
				Name:      "edit",
				Aliases:   []string{"e"},
				Usage:     "edit ssh config",
				Action:    commands.Edit,
				ArgsUsage: "<config name>",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
