package login

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
	logger = logrus.WithField("component", "login")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	logger.Info("login service startup")

	// register game handler
	comps := &component.Components{}
	comps.Register(&ServiceLogin{})

	nano.Listen(viper.GetString("login.listen"),
		nano.WithAdvertiseAddr(viper.GetString("master.listen")),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithDebugMode(),
	)

	return nil
}
