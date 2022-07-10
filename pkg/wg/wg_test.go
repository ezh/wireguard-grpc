package wg_test

import (
	"errors"
	"net"
	"os"
	osExec "os/exec"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	mock "github.com/ezh/wireguard-grpc/test/mock"
)

const (
	wgOutputExample      = "wireguard-tools v1.0.20210914 - https://git.zx2c4.com/wireguard-tools/"
	wgDumpOutputExample1 = "" +
		"wg0\tKOnxWiTPGxFW9AFBCI0NSLDTuZtmYNSKoM5Tb4auvlc=\tSFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw=\t51820\toff\n" +
		"wg0\tGQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE=\t(none)\t10.42.0.29:51820\t10.255.255.2/32\t1656746383\t24732\t22420\t5\n" // nolint
	wgDumpOutputExample2 = "" +
		"wg0\tKOnxWiTPGxFW9AFBCI0NSLDTuZtmYNSKoM5Tb4auvlc=\tSFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw=\t51820\toff\n" +
		"wg0\tGQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE=\t(none)\t10.42.0.64:51820\t10.255.255.2/32,192.168.0.1/32,192.168.0.2/32\t1657479339\t1576\t1592\t5" // nolint
	wgVersion = "v1.0.20210914"
)

var _ = Describe("Test WG package", func() {
	var (
		mockCtrl     *gomock.Controller
		mockExecutor *mock.MockExecutor
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockExecutor = mock.NewMockExecutor(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Specify("Test Version behavior", func() {
		wg := wg.New("wg")
		wg.Executor = mockExecutor
		mockExecutor.EXPECT().Run(logger, "-v").
			Return(wgOutputExample, "", nil).Times(1)
		version, err := wg.Version(logger)
		Expect(err).To(Succeed())
		Expect(version).To(Equal(wgVersion))
	})
	Specify("Test Version behavior with broken wg output", func() {
		wg := wg.New("wg")
		wg.Executor = mockExecutor
		mockExecutor.EXPECT().Run(logger, "-v").
			Return("wireguard-toolsv1.0.20210914 ", "", errors.New("unable to get wg version")).Times(1)
		_, err := wg.Version(logger)
		Expect(err).ToNot(Succeed())
	})
	Specify("Test RunDiag returns exit code 1 for broken wg", func() {
		var exerr *osExec.ExitError

		os.Setenv("WG_EXE", "false")
		err := app.RunDiag(logger, config.ReadConfig(), GinkgoWriter)
		Expect(err).To(BeAssignableToTypeOf(&osExec.ExitError{}))
		if errors.As(err, &exerr) {
			Expect(exerr.ExitCode()).To(Equal(1))
		}
	})
	Specify("Test Dump behavior", func() { // nolint:dupl
		wg := wg.New("wg")
		wg.Executor = mockExecutor
		mockExecutor.EXPECT().Run(logger, "show", "all", "dump").
			Return(wgDumpOutputExample1, "", nil).Times(1)
		devices, err := wg.Dump(logger)
		Expect(err).To(Succeed())
		Expect(devices).Should(HaveLen(1))
		Expect(devices[0].PrivateKey.String()).To(Equal("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="))
		Expect(devices[0].PublicKey.String()).To(Equal("SFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw="))
		Expect(devices[0].ListenPort).To(Equal(51820))
		Expect(devices[0].FirewallMark).To(Equal(0))
		Expect(devices[0].Peers).To(HaveLen(1))
		peer := devices[0].Peers[0]
		Expect(peer.PublicKey.String()).To(Equal("GQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE="))
		Expect(peer.PresharedKey.String()).To(Equal("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="))
		Expect(*peer.Endpoint).To(Equal(net.UDPAddr{
			IP:   net.ParseIP("10.42.0.29"),
			Port: 51820,
		}))
		Expect(peer.PersistentKeepaliveInterval).To(Equal(time.Second * 5))
		Expect(peer.LastHandshakeTime).To(Equal(time.Unix(1656746383, 0)))
		Expect(peer.ReceiveBytes).To(Equal(int64(24732)))
		Expect(peer.TransmitBytes).To(Equal(int64(22420)))
		Expect(peer.AllowedIPs).To(HaveLen(1))
	})
	Specify("Test Dump behavior with multiple allowed IPs", func() { // nolint:dupl
		wg := wg.New("wg")
		wg.Executor = mockExecutor
		mockExecutor.EXPECT().Run(logger, "show", "all", "dump").
			Return(wgDumpOutputExample2, "", nil).Times(1)
		devices, err := wg.Dump(logger)
		Expect(err).To(Succeed())
		Expect(devices).Should(HaveLen(1))
		Expect(devices[0].PrivateKey.String()).To(Equal("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="))
		Expect(devices[0].PublicKey.String()).To(Equal("SFPxov7YbLKXOPWuluBfm6RnITaWAjN2S67TNuSsMRw="))
		Expect(devices[0].ListenPort).To(Equal(51820))
		Expect(devices[0].FirewallMark).To(Equal(0))
		Expect(devices[0].Peers).To(HaveLen(1))
		peer := devices[0].Peers[0]
		Expect(peer.PublicKey.String()).To(Equal("GQ+LsM9LFtXYXv+tVXWwWVa2QPexzrEekABPvgKUHRE="))
		Expect(peer.PresharedKey.String()).To(Equal("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="))
		Expect(*peer.Endpoint).To(Equal(net.UDPAddr{
			IP:   net.ParseIP("10.42.0.64"),
			Port: 51820,
		}))
		Expect(peer.PersistentKeepaliveInterval).To(Equal(time.Second * 5))
		Expect(peer.LastHandshakeTime).To(Equal(time.Unix(1657479339, 0)))
		Expect(peer.ReceiveBytes).To(Equal(int64(1576)))
		Expect(peer.TransmitBytes).To(Equal(int64(1592)))
		Expect(peer.AllowedIPs).To(HaveLen(3))
	})
})
