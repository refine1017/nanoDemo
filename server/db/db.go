package db

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
	logger = logrus.WithField("component", "db")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	if err := initOrm(); err != nil {
		logger.Error("orm init with err: %v", err)
		return err
	}

	logger.Info("game service startup")

	// register game handler
	comps := &component.Components{}
	comps.Register(playerManager)

	addr := fmt.Sprintf("%s:%d", viper.GetString("db.host"), viper.GetInt("db.port"))
	nano.Listen(addr,
		nano.WithHeartbeatInterval(time.Duration(viper.GetInt("core.heartbeat"))*time.Second),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)

	return nil
}
