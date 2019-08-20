package game

import (
	"fmt"
	"math/rand"
	"server/db"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
	"github.com/lonng/nano/serialize/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = logrus.WithField("component", "game")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	if err := db.Startup(); err != nil {
		logger.Error("db startup with err: %v", err)
		return err
	}

	logger.Info("game service startup")

	// register game handler
	comps := &component.Components{}
	comps.Register(playerManager)

	// crypto
	c := newCrypto()
	pip := pipeline.New()
	pip.Inbound().PushBack(c.inbound)
	pip.Outbound().PushBack(c.outbound)

	addr := fmt.Sprintf("%s:%d", viper.GetString("game.host"), viper.GetInt("game.port"))
	nano.Listen(addr,
		nano.WithPipeline(pip),
		nano.WithHeartbeatInterval(time.Duration(viper.GetInt("core.heartbeat"))*time.Second),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)

	return nil
}
