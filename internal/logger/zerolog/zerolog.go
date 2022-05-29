package zerolog

import (
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func New(logLevel int, options ...logger.Option) (error, logr.Logger) {
	lo := logger.Get(options...)
	zl := zerolog.New(lo.Output)
	return nil, zerologr.New(&zl).V(logLevel)
}
