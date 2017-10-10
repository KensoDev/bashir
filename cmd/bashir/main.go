package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/kensodev/bashir"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"

	var configFileLocation string

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Run commands from a bashir configuration file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Load configuration from `FILE`",
					Destination: &configFileLocation,
				},
			},
			Action: func(c *cli.Context) error {

				parser := bashir.NewParser(configFileLocation)
				config, err := parser.ParseConfigurationFile()

				if err != nil {
					color.Red(err.Error())
					return err
				}

				runner := bashir.NewRunner(config)
				runner.RunCommands()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
