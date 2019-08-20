package game

import (
	"fmt"
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

	// register game handler
	comps := &component.Components{}
	comps.Register(playerManager)

	addr := fmt.Sprintf("%s:%d", viper.GetString("game.host"), viper.GetInt("game.port"))
	nano.Listen(addr,
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)

	return nil
}
