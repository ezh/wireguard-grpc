package zap_test

import (
	"bytes"

	"github.com/ezh/wireguard-grpc/internal/logger/zap"
	"github.com/ezh/wireguard-grpc/pkg/logger"
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
		err, logr := zap.New(0, logger.WithOutput(logFile))
		Expect(err).To(Succeed())
		logr.Info("zap message", "string", "details", "number", 1)
		Expect(logFile.String()).To(ContainSubstring(",\"M\":\"zap message\",\"string\":\"details\",\"number\":1}\n"))
	})
})
