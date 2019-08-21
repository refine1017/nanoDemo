package db

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
	comps.Register(&ServiceDB{})

	nano.Listen(viper.GetString("db.listen"),
		nano.WithAdvertiseAddr(viper.GetString("master.listen")),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithDebugMode(),
	)

	return nil
}
