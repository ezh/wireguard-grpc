//go:build integration
// +build integration

package p2p_test

import (
	"time"

	"github.com/ezh/wireguard-grpc/test/utilities/kube"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("P2P", func() {
	It("", func() {
		kubeTest, err := kube.NewTest()
		defer kubeTest.Close()
		Expect(err).NotTo(HaveOccurred())

		d := kubeTest.LoadDeployment("wireguard-deployment.yaml")
		kubeTest.CreateDeployment(kubeTest.Namespace, d)
		kubeTest.WaitForDeploymentReady(d, 10*time.Minute)
		pods := kubeTest.ListPodsFromDeployment(d)
		Expect(len(pods.Items)).To(Equal(2))
	})
})
