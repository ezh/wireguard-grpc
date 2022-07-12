package wgquick

import (
	"strings"

	"github.com/ezh/wireguard-grpc/internal/l"
	"github.com/ezh/wireguard-grpc/pkg/exec"
)

type Exec struct {
	exec.Executor
}

func New(rawCmd string) *Exec {
	return &Exec{Executor: exec.New(rawCmd)}
}

func (exe *Exec) Verify() bool {
	stdout, stderr, err := exe.Run("-h")
	if err != nil || !strings.HasPrefix(stderr, "Usage: wg-quick") {
		l.Error(err, "wg-quick failed", "stdout", stdout, "stderr", stderr)
		return false
	}
	return true
}
