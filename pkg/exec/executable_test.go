package exec_test

import (
	"errors"
	osExec "os/exec"

	"github.com/ezh/wireguard-grpc/pkg/exec"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Executable Interface", func() {
	Specify("Test Run behavior", func() {
		testExec := exec.Executable{
			Cmd: "bash",
		}
		stdout, stderr, err := testExec.Run(logger, "-c", "echo 123")
		Expect(err).To(Succeed())
		Expect(stdout).To(Equal("123"))
		Expect(stderr).To(BeEmpty())
	})
	Specify("Test RunCombined behavior", func() {
		testExec := exec.Executable{
			Cmd: "bash",
		}
		out, err := testExec.RunCombined(logger, "-c", "echo 123;echo 321 >&2")
		Expect(err).To(Succeed())
		Expect(out).To(Equal("123\n321"))
	})
	Specify("Test Run Exit code with 'false'", func() {
		var exerr *osExec.ExitError
		testExec := exec.Executable{
			Cmd: "false",
		}
		stdout, stderr, err := testExec.Run(logger)
		Expect(err).To(BeAssignableToTypeOf(&osExec.ExitError{}))
		Expect(stdout).To(BeEmpty())
		Expect(stderr).To(BeEmpty())
		if errors.As(err, &exerr) {
			Expect(exerr.ExitCode()).To(Equal(1))
		}
	})
})
