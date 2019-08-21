package gate

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logger = logrus.WithField("component", "gate")
)

func Startup() error {
	rand.Seed(time.Now().Unix())

	logger.Info("gate service startup")

	// register handler
	comps := &component.Components{}
	comps.Register(&ServiceGate{})

	nano.Listen(viper.GetString("gate.listen"),
		nano.WithAdvertiseAddr(viper.GetString("master.listen")),
		nano.WithClientAddr(viper.GetString("gate.client_listen")),
		nano.WithHeartbeatInterval(time.Duration(viper.GetInt("core.heartbeat"))*time.Second),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithIsWebsocket(true),
		nano.WithWSPath("/nano"),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithDebugMode(),
	)

	return nil
}
