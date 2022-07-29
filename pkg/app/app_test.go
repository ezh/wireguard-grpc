package app_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"time"

	wireguardv1 "github.com/ezh/wireguard-grpc/api/wireguard/v1"
	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/ezh/wireguard-grpc/test/mock"
	"github.com/ezh/wireguard-grpc/test/utilities"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var wgDumpOutputExample = "" +
	"wg0\tKOnxWiTPGxFW9AFBCI0NSLDTuZtmYNSKoM5Tb4auvlc=\tSFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw=\t51820\toff\n" +
	"wg0\tGQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE=\t(none)\t10.42.0.29:51820\t10.255.255.2/32\t1656746383\t24732\t22420\t5\n" // nolint

var _ = Describe("App", func() {
	var (
		mockCtrl     *gomock.Controller
		mockExecutor *mock.MockExecutor
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(utilities.MockGinkgoT())
		mockExecutor = mock.NewMockExecutor(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Specify("Test GRPC dump behavior", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		lis := utilities.NewPipeListener()
		app := app.New(&config.Config{})
		app.WG.Executor = mockExecutor
		app.WGQuick.Executor = mockExecutor
		mockExecutor.EXPECT().Run("show").
			Return("wg stdout text", "", nil).Times(1)
		mockExecutor.EXPECT().Run("-h").
			Return("", "Usage: wg-quick [ up | down | save | strip ] [ CONFIG_FILE | INTERFACE ]", nil).Times(1)
		mockExecutor.EXPECT().Run("show", "all", "dump").
			Return(wgDumpOutputExample, "", nil).Times(1)

		// Server
		go func() {
			defer GinkgoRecover()
			err := app.RunServer(ctx, lis)
			Expect(err).To(Succeed())
		}()

		// Client
		cCtx, cCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cCancel()
		cConn, err := grpc.DialContext(cCtx, "",
			grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(lis.Dialer()))
		Expect(err).To(Succeed())
		defer cConn.Close()
		client := wireguardv1.NewWireGuardServiceClient(cConn)

		resPing, err := client.Ping(cCtx, &emptypb.Empty{})
		Expect(resPing).To(BeAssignableToTypeOf(&emptypb.Empty{}))
		Expect(err).To(Succeed())

		resDump, err := client.Dump(cCtx, &wireguardv1.DumpRequest{})
		Expect(resDump).To(BeAssignableToTypeOf(&wireguardv1.DumpResponse{}))
		Expect(err).To(Succeed())
		Expect(resDump.GetInterfaces()).To(HaveLen(1))

		iFace := resDump.GetInterfaces()[0]
		Expect(iFace.GetFirewallMark()).To(BeEquivalentTo(0))
		Expect(iFace.GetListenPort()).To(BeEquivalentTo(51820))
		Expect(base64.StdEncoding.EncodeToString(iFace.GetPublicKey())).
			To(Equal("SFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw="))
		Expect(iFace.GetWgIfName()).To(Equal("wg0"))

		peers := iFace.GetPeers()
		Expect(peers).To(HaveLen(1))
		peer := peers[0]
		Expect(peer.GetAllowedIps()).To(Equal([]string{"10.255.255.2/32"}))
		Expect(peer.GetEndpointIp()).To(Equal("10.42.0.29"))
		Expect(peer.GetEndpontPort()).To(BeEquivalentTo(51820))
		Expect(peer.GetFlags()).To(BeZero())
		Expect(peer.GetLastHandshakeTime()).To(BeEquivalentTo(1656746383))
		Expect(peer.GetPersistentKeepalive()).To(BeEquivalentTo(5))
		Expect(base64.StdEncoding.EncodeToString(peer.GetPresharedKey())).
			To(Equal("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="))
		Expect(base64.StdEncoding.EncodeToString(peer.GetPublicKey())).
			To(Equal("GQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE="))
		Expect(peer.GetReceiveBytes()).To(BeEquivalentTo(24732))
		Expect(peer.GetTransmitBytes()).To(BeEquivalentTo(22420))
	})
	Specify("Test RunDiag behavior", func() {
		var out bytes.Buffer
		app := app.New(&config.Config{})
		app.WG.Executor = mockExecutor
		app.WGQuick.Executor = mockExecutor
		mockExecutor.EXPECT().Run("-h").
			Return("", "Usage: wg-quick [ up | down | save | strip ] [ CONFIG_FILE | INTERFACE ]", nil).Times(1)
		mockExecutor.EXPECT().GetCmd().
			Return("wg-quick", []string{})
		mockExecutor.EXPECT().Run("-v").
			Return("wireguard-tools v1.0.20210914 - https://git.zx2c4.com/wireguard-tools/\n", "", nil).Times(1)
		mockExecutor.EXPECT().GetCmd().
			Return("wg", []string{})

		Expect(app.RunDiag(&out)).To(Succeed())
		Expect(out.String()).To(ContainSubstring("wg version: v1.0.20210914"))
	})
})
