package app_test

import (
	"testing"

	"github.com/ezh/wireguard-grpc/test/utilities"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var logger = utilities.NewGinkgoLogger()

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}
