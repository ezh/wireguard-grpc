package local_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Local Integration Tests", func() {
	It("Test CLI -h", func() {
		var stderr, stdout bytes.Buffer
		command := exec.Command(pathToCLI, "-h")
		session, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session, executableTimeout).Should(gexec.Exit(0))
		fmt.Fprint(GinkgoWriter, stdout.String())
		Expect(stderr.String()).To(BeEmpty())
		Expect(stdout.String()).To(ContainSubstring("wireguard-grpc is a wireguard GRPC API"))
	})

	It("Test CLI diag with broken wg", func() {
		var stderr, stdout bytes.Buffer
		command := exec.Command(pathToCLI, "diag")
		command.Env = os.Environ()
		command.Env = append(command.Env, "WG_EXE=false")
		session, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session, executableTimeout).Should(gexec.Exit(1))
	})
})
