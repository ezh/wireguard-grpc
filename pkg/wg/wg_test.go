package wg_test

import (
	"errors"
	"os"
	osExec "os/exec"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ezh/wireguard-grpc/config"
	"github.com/ezh/wireguard-grpc/pkg/app"
	"github.com/ezh/wireguard-grpc/pkg/wg"
	mock "github.com/ezh/wireguard-grpc/test/mock"
)

const (
	wgOutputExample = "wireguard-tools v1.0.20210914 - https://git.zx2c4.com/wireguard-tools/"
	wgVersion       = "v1.0.20210914"
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
})
