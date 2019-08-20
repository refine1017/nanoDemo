package main

import (
	"github.com/urfave/cli"
	"os"
	"server/cmd"
)

func main() {
	app := cli.NewApp()

	app.Name = "server"
	app.Author = "refine1017"
	app.Version = "0.0.1"
	app.Usage = "server"

	app.Commands = []cli.Command{
		cmd.Gate,
		cmd.Login,
		cmd.Game,
		cmd.DB,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: `./bin/config/config.toml`,
			Usage: "load configuration from `FILE`",
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
