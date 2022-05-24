package zap

import (
	"github.com/ezh/wireguard-grpc/internal/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapOption zap.Option

func New(options ...logger.Option) logr.Logger {
	lo := logger.Get(options...)
	config := zap.NewDevelopmentConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.AddSync(lo.Output), config.Level)
	zapLog := zap.New(core, lo.ZapOptions...)
	return zapr.NewLogger(zapLog)
}
