//go:build integration
// +build integration

package p2p_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestP2p(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "P2p Suite")
}
