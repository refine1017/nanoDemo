package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"os"
	"server/game"
	"sync"
)

func main() {
	app := cli.NewApp()

	app.Name = "server"
	app.Author = "refine1017"
	app.Version = "0.0.1"
	app.Usage = "server"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: ".\\bin\\config\\config.toml",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Action = serve
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func serve(c *cli.Context) error {
	viper.SetConfigType("toml")
	viper.SetConfigFile(c.String("config"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	if viper.GetBool("core.debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() { defer wg.Done(); game.Startup() }()

	wg.Wait()
	return nil
}
