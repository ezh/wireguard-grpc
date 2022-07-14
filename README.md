[![Build Status](https://github.com/ezh/wireguard-grpc/actions/workflows/go/badge.svg)](https://github.com/ezh/wireguard-grpc/actions?query=workflow%3Ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/ezh/wireguard-grpc)](https://goreportcard.com/report/github.com/ezh/wireguard-grpc)

# Generate

`go generate -v`

# Test

## Unit tests

`ginkgo ./...`

## Integration tests

`ginkgo -tags=integration -v test/...`

# Build

`go build cmd/main.go`
