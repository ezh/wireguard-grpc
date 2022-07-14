package exec_test

import (
	"testing"

	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/ezh/wireguard-grpc/test/utilities"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApp(t *testing.T) {
	app.RegisterLogger(utilities.NewGinkgoLogger())
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exec Suite")
}
