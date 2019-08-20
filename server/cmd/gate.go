package cmd

import (
	"github.com/urfave/cli"
	"server/game"
)

var Gate = cli.Command{
	Name:   "gate",
	Usage:  "gate server",
	Flags:  []cli.Flag{},
	Action: action(runGate),
}

func runGate(c *cli.Context) error {
	return game.Startup()
}
