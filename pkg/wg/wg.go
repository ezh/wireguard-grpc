package wg

import (
	"errors"
	"regexp"

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
	out, err := exe.RunCombined(l, "show")
	if err != nil {
		l.Error(err, "wg failed", "output", out)
		return false
	}
	return true
}

func (exe *Exec) Version(l *logr.Logger) (string, error) {
	stdout, _, err := exe.Run(l, "-v")
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`\bv\d+[[:graph:]]+`)
	version := re.FindString(stdout)
	if len(version) == 0 {
		return "", errors.New("unable to get wg version")
	}
	return version, nil
}
