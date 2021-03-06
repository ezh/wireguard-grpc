package utilities

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	ginkgo "github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewGinkgoLogger returns new logger with ginkgo backend.
func NewGinkgoLogger(tags ...string) *logr.Logger {
	if len(tags)%2 != 0 {
		panic("Odd number of elements in args for NewGinkgoLogger")
	}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(ginkgo.GinkgoWriter), zap.DebugLevel)
	zaplog := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel), zap.AddCallerSkip(1))
	var i int
	if len(tags) > 0 {
		for ; i < len(tags); i += 2 {
			zaplog = zaplog.WithOptions(zap.Fields(zap.String(tags[i], tags[i+1])))
		}
	}
	logger := zapr.NewLogger(zaplog)
	return &logger
}
