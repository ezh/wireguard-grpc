package local_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const (
	executableTimeout = 10 * time.Second
)

var pathToCLI string

func TestLocal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Local Suite")
}

var _ = BeforeSuite(func() {
	var err error

	pathToCLI, err = gexec.Build("github.com/ezh/wireguard-grpc/cmd", "-race")
	Expect(err).ShouldNot(HaveOccurred())
	By(fmt.Sprintf("CLI is ready at %s", pathToCLI))
})
