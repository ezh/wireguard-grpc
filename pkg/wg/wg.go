package wg

import (
	"github.com/ezh/wireguard-grpc/pkg/exec"
	"github.com/go-logr/logr"
)

type Exec struct {
	exec.Executable
}

func (exe *Exec) Verify(l *logr.Logger) bool {
	out, err := exe.RunCombined("show")
	if err != nil {
		l.Error(err, "wg failed", "output", out)
		return false
	}
	return true
}
