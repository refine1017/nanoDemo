package game

import (
	"math/rand"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = logrus.WithField("component", "game")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	logger.Info("game service startup")

	// register handler
	comps := &component.Components{}
	comps.Register(&ServiceGame{})

	nano.Listen(viper.GetString("game.listen"),
		nano.WithAdvertiseAddr(viper.GetString("master.listen")),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithDebugMode(),
	)

	return nil
}
