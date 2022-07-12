package zap

import (
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates Zap logger
func New(logLevel logger.LogLevel, options ...logger.Option) (logr.Logger, error) {
	lo := logger.Get(options...)
	config := zap.NewDevelopmentConfig()
	switch logLevel {
	case logger.LogInfo:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case logger.LogDebug:
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case logger.LogTrace:
		config.Level = zap.NewAtomicLevelAt(-127)
	case logger.LogOff:
		config.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.AddSync(lo.Output), config.Level)
	zapLog := zap.New(core, lo.ZapOptions...)
	return zapr.NewLogger(zapLog), nil
}
