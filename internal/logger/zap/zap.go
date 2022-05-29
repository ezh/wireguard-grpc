package zap

import (
	"github.com/ezh/wireguard-grpc/pkg/logger"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(logLevel int, options ...logger.Option) (error, logr.Logger) {
	lo := logger.Get(options...)
	config := zap.NewDevelopmentConfig()
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig),
		zapcore.AddSync(lo.Output), config.Level)
	zapLog := zap.New(core, lo.ZapOptions...)
	return nil, zapr.NewLogger(zapLog).V(logLevel)
}
