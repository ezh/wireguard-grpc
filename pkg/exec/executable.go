package exec

import (
	"bytes"
	"os/exec"
	"strings"
)

type Executable struct {
	Cmd  string
	Sudo bool
}

func (exe *Executable) Run(args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	err := exe.run(&stdout, &stderr, args...)
	return strings.TrimSpace(stdout.String()),
		strings.TrimSpace(stderr.String()), err
}

func (exe *Executable) RunCombined(args ...string) (string, error) {
	var buf bytes.Buffer
	err := exe.run(&buf, &buf, args...)
	return strings.TrimSpace(buf.String()), err
}

func (exe *Executable) run(stdout, stderr *bytes.Buffer, args ...string) error {
	var cmd *exec.Cmd

	if exe.Sudo {
		argsForSudo := make([]string, len(args)+1)
		argsForSudo[0] = exe.Cmd
		copy(argsForSudo[1:], args)
		cmd = exec.Command("sudo", argsForSudo...) // #nosec G204
	} else {
		cmd = exec.Command(exe.Cmd, args...) // #nosec G204
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return err
}
