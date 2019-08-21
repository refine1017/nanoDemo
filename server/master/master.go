package master

import (
	"github.com/lonng/nano/examples/cluster/master"
	"github.com/lonng/nano/session"
	"math/rand"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = logrus.WithField("component", "master")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	logger.Info("master service startup")

	// register handler
	comps := &component.Components{}

	session.Lifetime.OnClosed(master.OnSessionClosed)

	nano.Listen(viper.GetString("master.listen"),
		nano.WithMaster(),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithDebugMode(),
	)

	return nil
}
