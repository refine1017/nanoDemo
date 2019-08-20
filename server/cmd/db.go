package cmd

import (
	"github.com/urfave/cli"
	"server/game"
)

var DB = cli.Command{
	Name:   "db",
	Usage:  "db server",
	Flags:  []cli.Flag{},
	Action: action(runDB),
}

func runDB(c *cli.Context) error {
	return game.Startup()
}
