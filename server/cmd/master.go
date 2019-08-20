package cmd

import (
	"github.com/urfave/cli"
	"server/master"
)

var Master = cli.Command{
	Name:   "master",
	Usage:  "master server",
	Flags:  []cli.Flag{},
	Action: action(runMaster),
}

func runMaster(c *cli.Context) error {
	return master.Startup()
}
