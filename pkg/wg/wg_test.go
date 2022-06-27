package wg_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ezh/wireguard-grpc/pkg/wg"
)

const wgOutputExample = "wireguard-tools v1.0.20210914 - https://git.zx2c4.com/wireguard-tools/"
const wgVersion = "v1.0.20210914"

var _ = Describe("WG", func() {
	Specify("Version", func() {
		wg := wg.New("")
		version, err := wg.Version(nil)
		Expect(err).To(Succeed())
		Expect(version).To(Equal(wgVersion))
	})
})
