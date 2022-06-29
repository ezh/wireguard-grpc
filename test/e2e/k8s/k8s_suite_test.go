//go:build integration
// +build integration

package p2p_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/dlespiau/kube-test-harness/logger"
	"github.com/ezh/wireguard-grpc/test/utilities/kube"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	v1 "k8s.io/api/core/v1"
)

const (
	// Keep K8S namespace and gexec.Build artifact, dump K8S objects
	DEBUG             = false
	executableTimeout = 10 * time.Second
	pollingInterval   = time.Second / 2
)

var kubeTest *kube.TestEx
var pathToCLI string
var podA v1.Pod
var podB v1.Pod

func podCmd(pod v1.Pod, cmd string) *exec.Cmd {
	var args []string
	args = append(args, "exec", "-n", kubeTest.Namespace, pod.Name, "--")
	args = append(args, strings.Fields(cmd)...)
	return exec.Command("kubectl", args...)
}

func getPodCmd(pod v1.Pod, cmd string) string {
	return fmt.Sprintf("kubectl exec -n %s %s -- %s", kubeTest.Namespace, pod.Name, cmd)
}

var _ = BeforeSuite(func() {
	var err error
	By("bootstrapping test environment")
	kubeTest, err = kube.NewTest(logger.Debug)
	Expect(err).NotTo(HaveOccurred())

	wgA := kubeTest.LoadDeployment("wireguard-a-deployment.yaml")
	kubeTest.CreateConfigMapFromFile(kubeTest.Namespace, "wireguard-a-configmap.yaml")
	kubeTest.CreateServiceFromFile(kubeTest.Namespace, "wireguard-a-service.yaml")
	kubeTest.CreateDeployment(kubeTest.Namespace, wgA)

	wgB := kubeTest.LoadDeployment("wireguard-b-deployment.yaml")
	kubeTest.CreateConfigMapFromFile(kubeTest.Namespace, "wireguard-b-configmap.yaml")
	kubeTest.CreateDeployment(kubeTest.Namespace, wgB)

	kubeTest.WaitForDeploymentReady(wgA, 10*time.Minute)
	kubeTest.WaitForDeploymentReady(wgB, 10*time.Minute)

	podsA := kubeTest.ListPodsFromDeployment(wgA)
	Expect(len(podsA.Items)).To(Equal(1))
	podsB := kubeTest.ListPodsFromDeployment(wgB)
	Expect(len(podsB.Items)).To(Equal(1))

	podA = podsA.Items[0]
	podB = podsB.Items[0]

	By("environment is ready")
	// Give 1s for wg setup
	time.Sleep(time.Second)
	Eventually(func() error {
		return InterceptGomegaFailure(func() {
			// Ping in pod A
			pingA1 := podCmd(podA, "ping -c 1 -i 0.1 10.255.255.1")
			pingA1s, err := gexec.Start(pingA1, GinkgoWriter, GinkgoWriter)
			Expect(err).To(Succeed())
			pingA2 := podCmd(podA, "ping -c 1 -i 0.1 10.255.255.2")
			pingA2s, err := gexec.Start(pingA2, GinkgoWriter, GinkgoWriter)
			Expect(err).To(Succeed())

			// Ping in pod B
			pingB1 := podCmd(podB, "ping -c 1 10.255.255.1")
			pingB1s, err := gexec.Start(pingB1, GinkgoWriter, GinkgoWriter)
			Expect(err).To(Succeed())
			pingB2 := podCmd(podB, "ping -c 1 10.255.255.2")
			pingB2s, err := gexec.Start(pingB2, GinkgoWriter, GinkgoWriter)
			Expect(err).To(Succeed())

			Eventually(pingA1s).Should(gexec.Exit(0))
			Eventually(pingA2s).Should(gexec.Exit(0))
			Eventually(pingB1s).Should(gexec.Exit(0))
			Eventually(pingB2s).Should(gexec.Exit(0))
		})
	}, executableTimeout, pollingInterval).Should(Succeed())
	pathToCLI, err = gexec.Build("github.com/ezh/wireguard-grpc/cmd", "-race")
	Expect(err).ShouldNot(HaveOccurred())
	By(fmt.Sprintf("CLI is ready at %s", pathToCLI))

})

var _ = AfterEach(func() {
	if DEBUG && CurrentSpecReport().Failed() {
		kubeTest.Fail()
	}
})

var _ = AfterSuite(func() {
	if kubeTest != nil {
		kubeTest.Close()
	}
	gexec.KillAndWait()
	// Debug github.com/ezh/wireguard-grpc/cmd
	if !DEBUG && (kubeTest == nil || !kubeTest.Failed()) {
		gexec.CleanupBuildArtifacts()
	}
})

func TestP2p(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P2P Suite")
}
