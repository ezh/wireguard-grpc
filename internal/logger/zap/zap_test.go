package zap_test

import (
	"bytes"

	"github.com/ezh/wireguard-grpc/internal/logger"
	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
)

func NewTestCore(buf *bytes.Buffer) zapcore.Core {
	config := zapcore.EncoderConfig{
		MessageKey: "M",
	}
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(buf),
		zapcore.DebugLevel,
	)
}

var _ = Describe("Zap", func() {
	It("should produce log message", func() {
		logFile := new(bytes.Buffer)
		logr := zap.New(logger.WithOutput(logFile))
		logr.Info("zap message", "string", "details", "number", 1)
		Expect(logFile.String()).To(ContainSubstring(",\"M\":\"zap message\",\"string\":\"details\",\"number\":1}\n"))
	})
})
