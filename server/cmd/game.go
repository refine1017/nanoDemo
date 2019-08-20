package cmd

import (
	"github.com/urfave/cli"
	"server/game"
)

var Game = cli.Command{
	Name:   "game",
	Usage:  "game server",
	Flags:  []cli.Flag{},
	Action: action(runGame),
}

func runGame(c *cli.Context) error {
	return game.Startup()
}
