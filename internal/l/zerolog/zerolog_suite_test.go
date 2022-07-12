package zerolog_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZerolog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zerolog Suite")
}
