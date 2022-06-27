package wg_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wg Suite")
}
