package exec

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/go-logr/logr"
)

type Executor interface {
	Run(l *logr.Logger, args ...string) (string, string, error)
	RunCombined(l *logr.Logger, args ...string) (string, error)
	GetCmd() (string, []string)
}

type Executable struct {
	Cmd     string
	CmdArgs []string
}

var _ Executor = (*Executable)(nil)

func New(rawCmd string) *Executable {
	args := strings.Fields(rawCmd)
	if len(args) == 0 {
		return &Executable{}
	}
	return &Executable{
		Cmd:     args[0],
		CmdArgs: args[1:],
	}
}

func (exe *Executable) GetCmd() (string, []string) {
	return exe.Cmd, exe.CmdArgs
}

func (exe *Executable) Run(l *logr.Logger, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	err := exe.run(l, &stdout, &stderr, args...)
	return strings.TrimSpace(stdout.String()),
		strings.TrimSpace(stderr.String()), err
}

func (exe *Executable) RunCombined(l *logr.Logger, args ...string) (string, error) {
	var buf bytes.Buffer
	err := exe.run(l, &buf, &buf, args...)
	return strings.TrimSpace(buf.String()), err
}

func (exe *Executable) run(l *logr.Logger, stdout, stderr *bytes.Buffer, args ...string) error {
	var cmd *exec.Cmd

	argsForCmd := make([]string, len(exe.CmdArgs)+len(args))
	copy(argsForCmd, exe.CmdArgs)
	copy(argsForCmd[len(exe.CmdArgs):], args)
	l.V(1).Info("execute command", "cmd", exe.Cmd, "args", argsForCmd)
	cmd = exec.Command(exe.Cmd, argsForCmd...) // #nosec G204
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	return err
}
