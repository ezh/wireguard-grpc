package zap

import (
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates Zap logger
func New(logLevel logger.LogLevel, options ...logger.Option) (error, logr.Logger) {
	lo := logger.Get(options...)
	config := zap.NewDevelopmentConfig()
	switch logLevel {
	case logger.LOG_INFO:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case logger.LOG_DEBUG:
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(-127)
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.AddSync(lo.Output), config.Level)
	zapLog := zap.New(core, lo.ZapOptions...)
	return nil, zapr.NewLogger(zapLog)
}
