//go:build integration
// +build integration

package p2p_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kubernetes Integration Tests", func() {
	FIt("Test API behavior 2", func() {
		// var stderr, stdout bytes.Buffer
		// command := exec.Command(pathToCLI, "diag")
		// command.Env = os.Environ()
		// command.Env = append(command.Env, "WG_EXE=false")
		// session, err := gexec.Start(command, &stdout, &stderr)
		// Expect(err).ShouldNot(HaveOccurred())
		// Eventually(session, executableTimeout).Should(gexec.Exit(0))
	})
	It("Test API behavior 3", func() {
		Expect(false).To(BeTrue())
	})
})
