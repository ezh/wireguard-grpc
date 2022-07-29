package utilities_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ezh/wireguard-grpc/test/utilities"
)

var _ = Describe("Logger", func() {
	Specify("NewGinkgoLogger tags", func() {
		var buf bytes.Buffer
		GinkgoWriter.TeeTo(&buf)
		l := utilities.NewGinkgoLogger("abc", "def", "123", "321")
		l.Info("123")
		Expect(buf.String()).To(ContainSubstring(`"abc":"def","123":"321"`))
	})
})
