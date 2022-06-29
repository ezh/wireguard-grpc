//go:build tools
// +build tools

// go mod tidy acts as if all build tags are enabled,
// so it will consider platform-specific source files
// and files that require custom build tags,
// even if those source files wouldn't normally be built.
//
// There is one exception: the ignore build tag is not enabled,
// so a file with the build constraint + build ignore will not be considered.

// run 'go generate --tags=tools -x'

package wireguardgrpc

//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
//go:generate go install github.com/bufbuild/buf/cmd/buf@v1.5.0
//go:generate go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@v1.5.0
//go:generate go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v1.5.0
//go:generate go install github.com/golang/mock/mockgen@v1.6.0
