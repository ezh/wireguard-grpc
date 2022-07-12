package zap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zap Suite")
}
