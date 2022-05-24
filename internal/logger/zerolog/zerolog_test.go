package zerolog_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ezh/wireguard-grpc/internal/logger"
	"github.com/ezh/wireguard-grpc/internal/logger/zerolog"
)

var _ = Describe("Zerolog", func() {
	It("should produce log message", func() {
		logFile := new(bytes.Buffer)
		logr := zerolog.New(logger.WithOutput(logFile))
		logr.Info("zero message", "string", "details", "number", 1)
		Expect(logFile.String()).To(Equal("{\"level\":\"info\",\"v\":0,\"string\":\"details\",\"number\":1,\"message\":\"zero message\"}\n"))
	})
})
