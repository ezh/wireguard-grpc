//go:build integration
// +build integration

package p2p_test

import (
	"os/exec"
	"time"

	"github.com/dlespiau/kube-test-harness/logger"
	"github.com/ezh/wireguard-grpc/test/utilities/kube"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const (
	executableTimeout = 10 * time.Second
)

var _ = Describe("P2P", func() {
	It("Test API behavior", func() {
		kubeTest, err := kube.NewTest(logger.Debug)
		defer kubeTest.Close()
		Expect(err).NotTo(HaveOccurred())

		wgA := kubeTest.LoadDeployment("wireguard-a-deployment.yaml")
		kubeTest.CreateConfigMapFromFile(kubeTest.Namespace, "wireguard-a-configmap.yaml")
		kubeTest.CreateServiceFromFile(kubeTest.Namespace, "wireguard-a-service.yaml")
		kubeTest.CreateDeployment(kubeTest.Namespace, wgA)
		kubeTest.WaitForDeploymentReady(wgA, 10*time.Minute)

		wgB := kubeTest.LoadDeployment("wireguard-b-deployment.yaml")
		kubeTest.CreateConfigMapFromFile(kubeTest.Namespace, "wireguard-b-configmap.yaml")
		kubeTest.CreateDeployment(kubeTest.Namespace, wgB)
		kubeTest.WaitForDeploymentReady(wgB, 10*time.Minute)

		podsA := kubeTest.ListPodsFromDeployment(wgA)
		Expect(len(podsA.Items)).To(Equal(1))
		podsB := kubeTest.ListPodsFromDeployment(wgB)
		Expect(len(podsB.Items)).To(Equal(1))

		// Ping in pod A
		pingA1 := exec.Command("kubectl", "exec", "-n", kubeTest.Namespace,
			podsA.Items[0].Name, "--", "ping", "-q", "-c", "1", "10.255.255.1")
		pingA1s, err := gexec.Start(pingA1, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Succeed())
		Eventually(pingA1s, executableTimeout).Should(gexec.Exit(0))
		pingA2 := exec.Command("kubectl", "exec", "-n", kubeTest.Namespace,
			podsA.Items[0].Name, "--", "ping", "-q", "-c", "1", "10.255.255.2")
		pingA2s, err := gexec.Start(pingA2, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Succeed())
		Eventually(pingA2s, executableTimeout).Should(gexec.Exit(0))

		// Ping in pod B
		pingB1 := exec.Command("kubectl", "exec", "-n", kubeTest.Namespace,
			podsB.Items[0].Name, "--", "ping", "-q", "-c", "1", "10.255.255.1")
		pingB1s, err := gexec.Start(pingB1, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Succeed())
		Eventually(pingB1s, executableTimeout).Should(gexec.Exit(0))
		pingB2 := exec.Command("kubectl", "exec", "-n", kubeTest.Namespace,
			podsB.Items[0].Name, "--", "ping", "-q", "-c", "1", "10.255.255.2")
		pingB2s, err := gexec.Start(pingB2, GinkgoWriter, GinkgoWriter)
		Expect(err).To(Succeed())
		Eventually(pingB2s, executableTimeout).Should(gexec.Exit(0))
	})
})
