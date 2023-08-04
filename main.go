package main

import (
	"log"
	"os"

	"github.com/elliot40404/ssm/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "ssm",
		Usage:   "ssh config manager",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list available ssh configs",
				Action:  commands.List,
			},
			{
				Name:   "add",
				Usage:  "add new ssh config",
				Action: commands.Add,
			},
			{
				Name:      "print",
				Usage:     "print ssh config",
				Action:    commands.Print,
				ArgsUsage: "<config name>",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
