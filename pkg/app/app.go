package app

import (
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
)

// Run creates objects via constructors.
func Run(logBuilder logger.LogBuilder, cfg *config.Config, verbosity int) error {
	if verbosity == -1 {
		logBuilder = func(logLevel logger.LogLevel, options ...logger.Option) (error, logr.Logger) {
			return nil, logr.Discard()
		}
	}
	logLevel := logger.ParseLogLevel(cfg.LogLevel)
	err, l := logBuilder(logger.LogLevel(int(logLevel) + verbosity))
	if err != nil {
		return err
	}
	l.V(0).Info("HELLO")
	l.V(1).Info("HELLO")
	l.V(2).Info("HELLO")
	l.V(3).Info("HELLO")
	return nil
}
