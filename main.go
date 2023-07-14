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
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "count",
						Value:   false,
						Usage:   "print serial numbers",
						Aliases: []string{"c"},
					},
				},
			},
			{
				Name:   "add",
				Usage:  "add new ssh config",
				Action: commands.Add,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
