package cmd

import (
	"github.com/urfave/cli"
	"server/game"
)

var Login = cli.Command{
	Name:   "login",
	Usage:  "login server",
	Flags:  []cli.Flag{},
	Action: action(runLogin),
}

func runLogin(c *cli.Context) error {
	return game.Startup()
}
