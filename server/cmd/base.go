package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"sync"
)

func action(callback func(ctx *cli.Context) error) func(ctx *cli.Context) error {
	return func(c *cli.Context) error {
		viper.SetConfigType("toml")
		viper.SetConfigFile(c.GlobalString("config"))

		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
		if viper.GetBool("core.debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			if err := callback(c); err != nil {
				panic(err)
			}
		}()

		wg.Wait()
		return nil
	}
}
