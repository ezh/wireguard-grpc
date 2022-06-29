//go:build integration
// +build integration

package p2p_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Kubernetes Integration Tests", func() {
	FIt("Test CLI diag", func() {
		var stderr, stdout bytes.Buffer
		command := exec.Command(pathToCLI, "diag")
		command.Env = os.Environ()
		command.Env = append(command.Env, fmt.Sprintf("WG_EXE=%s", getPodCmd(podA, "wg")))
		command.Env = append(command.Env, fmt.Sprintf("WGQUICK_EXE=%s", getPodCmd(podA, "wg-quick")))
		session, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session, executableTimeout).Should(gexec.Exit(0))
		fmt.Fprintf(GinkgoWriter, "STDOUT:\n%s\n", stdout.String())
		fmt.Fprintf(GinkgoWriter, "STDERR:\n%s\n", stderr.String())
	})
	It("Test API behavior 3", func() {
		Expect(false).To(BeTrue())
	})
})
