package zerolog

import (
	"github.com/ezh/wireguard-grpc/internal/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func New(options ...logger.Option) logr.Logger {
	lo := logger.Get(options...)
	zl := zerolog.New(lo.Output)
	return zerologr.New(&zl)
}
