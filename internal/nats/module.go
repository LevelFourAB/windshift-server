package nats

import (
	"github.com/levelfourab/sprout-go"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module that provides a connection to a NATS server.
var Module = fx.Module(
	"nats",
	fx.Provide(sprout.Logger("nats"), fx.Private),
	fx.Provide(sprout.Config("NATS", &Config{}), fx.Private),
	fx.Provide(newNats),
)

type Config struct {
	// URL is the URL of the NATS server to connect to.
	URL string `env:"URL,required"`
}

func newNats(logger *zap.Logger, config *Config) (*nats.Conn, error) {
	return nats.Connect(
		config.URL,
		nats.RetryOnFailedConnect(true),
		nats.ConnectHandler(func(c *nats.Conn) {
			logger.Info("Connected to NATS server", zap.String("url", config.URL))
		}),
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			logger.Info("Disconnected from NATS server", zap.String("url", config.URL), zap.Error(err))
		}),
		nats.ReconnectHandler(func(c *nats.Conn) {
			logger.Info("Reconnected to NATS server", zap.String("url", config.URL))
		}),
		nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
			logger.Error("Error processing incoming message", zap.Error(err))
		}),
	)
}
