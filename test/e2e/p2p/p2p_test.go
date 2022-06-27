//go:build integration
// +build integration

package p2p_test

import (
	"bytes"
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("P2P", func() {
	It("Test -h", func() {
		var stderr, stdout bytes.Buffer
		command := exec.Command(pathToCLI, "-h")
		session, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session, executableTimeout).Should(gexec.Exit(0))
		fmt.Fprint(GinkgoWriter, stdout.String())
		Expect(stderr.String()).To(BeEmpty())
		Expect(stdout.String()).To(ContainSubstring("wireguard-grpc is a wireguard GRPC API"))
	})
	It("Test API behavior 2", func() {
		Expect(true).To(BeTrue())
	})
	It("Test API behavior 3", func() {
		Expect(false).To(BeTrue())
	})
})
