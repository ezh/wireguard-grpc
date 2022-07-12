package wgquick

import (
	"strings"

	"github.com/ezh/wireguard-grpc/pkg/exec"
	"github.com/go-logr/logr"
)

type Exec struct {
	exec.Executor
}

func New(rawCmd string) *Exec {
	return &Exec{Executor: exec.New(rawCmd)}
}

func (exe *Exec) Verify(l *logr.Logger) bool {
	stdout, stderr, err := exe.Run(l, "-h")
	if err != nil || !strings.HasPrefix(stderr, "Usage: wg-quick") {
		l.Error(err, "wg-quick failed", "stdout", stdout, "stderr", stderr)
		return false
	}
	return true
}
