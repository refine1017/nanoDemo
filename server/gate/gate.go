package gate

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/pipeline"
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

	// register game handler
	comps := &component.Components{}

	// crypto
	c := newCrypto()
	pip := pipeline.New()
	pip.Inbound().PushBack(c.inbound)
	pip.Outbound().PushBack(c.outbound)

	addr := fmt.Sprintf("%s:%d", viper.GetString("gate.host"), viper.GetInt("gate.port"))
	nano.Listen(addr,
		nano.WithPipeline(pip),
		nano.WithHeartbeatInterval(time.Duration(viper.GetInt("core.heartbeat"))*time.Second),
		nano.WithLogger(logger),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithComponents(comps),
	)

	return nil
}
