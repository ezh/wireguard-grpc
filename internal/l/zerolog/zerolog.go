package zerolog

import (
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

// New creates Zerolog logger
func New(logLevel int, options ...logger.Option) (logr.Logger, error) {
	lo := logger.Get(options...)
	zl := zerolog.New(lo.Output)
	return zerologr.New(&zl).V(logLevel), nil
}
