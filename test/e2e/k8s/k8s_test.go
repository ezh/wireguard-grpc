//go:build integration
// +build integration

package p2p_test

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ = Describe("Kubernetes Integration Tests", func() {

	It("Test CLI diag", func() {
		var stderr, stdout bytes.Buffer

		GinkgoT().Setenv("WG_EXE", getPodCmd(podA, "wg"))
		GinkgoT().Setenv("WGQUICK_EXE", getPodCmd(podA, "wg-quick"))

		command := exec.Command(pathToCLI, "diag", "-v")
		session, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session, executableTimeout).Should(gexec.Exit(0))
		fmt.Fprintf(GinkgoWriter, "STDOUT:\n%s\n", stdout.String())
		fmt.Fprintf(GinkgoWriter, "STDERR:\n%s\n", stderr.String())
	})
	It("Test CLI dump", func() {
		var stderr, stdout bytes.Buffer

		socket := filepath.Join(GinkgoT().TempDir(), "test.sock")
		GinkgoT().Setenv("WG_EXE", getPodCmd(podA, "wg"))
		GinkgoT().Setenv("WGQUICK_EXE", getPodCmd(podA, "wg-quick"))
		GinkgoT().Setenv("LISTEN", socket)

		command := exec.Command(pathToCLI, "server", "-v")
		_, err := gexec.Start(command, &stdout, &stderr)
		Expect(err).ShouldNot(HaveOccurred())

		// Client
		cCtx, cCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cCancel()
		cConn, err := grpc.DialContext(cCtx, "unix://"+socket,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		Expect(err).To(Succeed())
		defer cConn.Close()
		client := wireguardv1.NewWireGuardServiceClient(cConn)

		Eventually(socket).Should(BeAnExistingFile())

		Eventually(func() error {
			return InterceptGomegaFailure(func() {
				resPing, err := client.Ping(cCtx, &emptypb.Empty{})
				Expect(resPing).To(BeAssignableToTypeOf(&emptypb.Empty{}))
				Expect(err).To(Succeed())
			})
		}, executableTimeout, pollingInterval).Should(Succeed())

		resDump, err := client.Dump(cCtx, &wireguardv1.DumpRequest{})
		Expect(resDump).To(BeAssignableToTypeOf(&wireguardv1.DumpResponse{}))
		Expect(err).To(Succeed())
		Expect(resDump.GetInterfaces()).To(HaveLen(1))

		iFace := resDump.GetInterfaces()[0]
		peers := iFace.GetPeers()
		Expect(peers).To(HaveLen(1))

		fmt.Fprintf(GinkgoWriter, "Dump: %v\n", resDump.Interfaces)
		fmt.Fprintf(GinkgoWriter, "STDOUT:\n%s\n", stdout.String())
		fmt.Fprintf(GinkgoWriter, "STDERR:\n%s\n", stderr.String())
	})
})
