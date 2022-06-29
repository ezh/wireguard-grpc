package wg_test

import (
	"testing"

	"github.com/ezh/wireguard-grpc/test/utilities"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var logger = utilities.NewGinkgoLogger()

func TestWg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wg Suite")
}
