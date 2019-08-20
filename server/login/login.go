package login

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
	logger = logrus.WithField("component", "login")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	logger.Info("login service startup")

	// register game handler
	comps := &component.Components{}

	addr := fmt.Sprintf("%s:%d", viper.GetString("login.host"), viper.GetInt("login.port"))
	nano.Listen(addr,
		nano.WithHeartbeatInterval(time.Duration(viper.GetInt("core.heartbeat"))*time.Second),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)

	return nil
}
