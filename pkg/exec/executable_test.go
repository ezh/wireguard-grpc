package exec_test

import (
	"errors"
	osExec "os/exec"

	"github.com/ezh/wireguard-grpc/pkg/exec"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Executable", func() {
	Specify("Run", func() {
		testExec := exec.Executable{
			Cmd: "bash",
		}
		stdout, stderr, err := testExec.Run("-c", "echo 123")
		Expect(err).To(Succeed())
		Expect(stdout).To(Equal("123"))
		Expect(stderr).To(BeEmpty())
	})
	Specify("RunCombined", func() {
		testExec := exec.Executable{
			Cmd: "bash",
		}
		out, err := testExec.RunCombined("-c", "echo 123;echo 321 >&2")
		Expect(err).To(Succeed())
		Expect(out).To(Equal("123\n321"))
	})
	Specify("Exit code", func() {
		var exerr *osExec.ExitError
		testExec := exec.Executable{
			Cmd: "false",
		}
		stdout, stderr, err := testExec.Run()
		Expect(err).To(BeAssignableToTypeOf(&osExec.ExitError{}))
		Expect(stdout).To(BeEmpty())
		Expect(stderr).To(BeEmpty())
		if errors.As(err, &exerr) {
			Expect(exerr.ExitCode()).To(Equal(1))
		}
	})
})
