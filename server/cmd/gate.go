package cmd

import (
	"github.com/urfave/cli"
	"server/gate"
)

var Gate = cli.Command{
	Name:   "gate",
	Usage:  "gate server",
	Flags:  []cli.Flag{},
	Action: action(runGate),
}

func runGate(c *cli.Context) error {
	return gate.Startup()
}
