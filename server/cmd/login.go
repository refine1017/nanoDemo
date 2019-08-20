package cmd

import (
	"github.com/urfave/cli"
	"server/login"
)

var Login = cli.Command{
	Name:   "login",
	Usage:  "login server",
	Flags:  []cli.Flag{},
	Action: action(runLogin),
}

func runLogin(c *cli.Context) error {
	return login.Startup()
}
