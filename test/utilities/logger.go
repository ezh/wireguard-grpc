package utilities

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	ginkgo "github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewGinkgoLogger returns new logger with ginkgo backend.
func NewGinkgoLogger() *logr.Logger {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(ginkgo.GinkgoWriter), zap.DebugLevel)
	zaplog := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
	logger := zapr.NewLogger(zaplog)
	return &logger
}
