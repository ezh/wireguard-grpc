package utilities

import (
	"fmt"

	ginkgo "github.com/onsi/ginkgo/v2"
)

type mockGinkgoT struct {
	ginkgo.GinkgoTInterface
}

// Fatal copies GoMock error message to console in case of panic
func (t *mockGinkgoT) Fatal(args ...interface{}) {
	t.Helper()
	fmt.Fprint(ginkgo.GinkgoWriter, args...)
	t.GinkgoTInterface.Fatal(args...)
}

// Fatalf copies GoMock error message to console in case of panic
func (t *mockGinkgoT) Fatalf(format string, args ...interface{}) {
	t.Helper()
	fmt.Fprintf(ginkgo.GinkgoWriter, fmt.Sprintf("Fatal: %s", format), args...)
	t.GinkgoTInterface.Fatalf(format, args...)
}

func MockGinkgoT() ginkgo.GinkgoTInterface {
	return &mockGinkgoT{
		GinkgoTInterface: ginkgo.GinkgoT(),
	}
}
