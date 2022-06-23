//go:build integration
// +build integration

package kube

import (
	"fmt"
	"os"
	"path"
	"strings"

	harness "github.com/dlespiau/kube-test-harness"
	"github.com/dlespiau/kube-test-harness/logger"
	"github.com/ezh/wireguard-grpc/test/utilities"
	ginkgo "github.com/onsi/ginkgo/v2"
)

type tHelper struct{ ginkgo.GinkgoTInterface }

func (t *tHelper) Helper()      {}
func (t *tHelper) Name() string { return "wireguard-grpc-test" }

// NewTest creates a new test harness to more easily run integration tests against the provided Kubernetes cluster.
func NewTest(kubeconfigPath ...string) (*harness.Test, error) {
	t := &tHelper{ginkgo.GinkgoT()}
	l := &logger.TestLogger{}
	h := harness.New(harness.Options{
		Logger:            l.ForTest(t),
		ManifestDirectory: path.Join(utilities.TestGetRoot(), "data"),
	})
	if err := h.Setup(); err != nil {
		return nil, err
	}
	if err := h.SetKubeconfig(
		strings.Join(kubeconfigPath, string(os.PathSeparator))); err != nil {
		return nil, err
	}
	test := h.NewTest(t)
	test.Setup()
	fmt.Fprintf(ginkgo.GinkgoWriter, "using namespace %s\n", test.Namespace)
	return test, nil
}
