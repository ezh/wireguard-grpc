package wgquick

import (
	"github.com/ezh/wireguard-grpc/pkg/exec"
	"github.com/go-logr/logr"
)

type Exec struct {
	exec.Executable
}

func New(rawCmd string) *Exec {
	return &Exec{Executable: exec.New(rawCmd)}
}

func (exe *Exec) Verify(l *logr.Logger) bool {
	out, err := exe.RunCombined("-h")
	if err != nil {
		l.Error(err, "wg-quick failed", "output", out)
		return false
	}
	return true
}
