package exec

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Executor interface {
	Run(args ...string) (string, string, error)
	RunCombined(args ...string) (string, error)
}

type Executable struct {
	Cmd     string
	CmdArgs []string
}

var _ Executor = (*Executable)(nil)

func New(rawCmd string) Executable {
	args := strings.Fields(rawCmd)
	if len(args) == 0 {
		panic(fmt.Errorf("unable to create Executable for '%s'", rawCmd))
	}
	return Executable{
		Cmd:     args[0],
		CmdArgs: args[1:],
	}
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

	argsForCmd := make([]string, len(exe.CmdArgs)+len(args))
	copy(argsForCmd, exe.CmdArgs)
	copy(argsForCmd[len(exe.CmdArgs):], args)
	cmd = exec.Command(exe.Cmd, argsForCmd...) // #nosec G204
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return err
}
