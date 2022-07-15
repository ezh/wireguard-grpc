[![Build Status](https://github.com/ezh/wireguard-grpc/actions/workflows/go.yml/badge.svg)](https://github.com/ezh/wireguard-grpc/actions?query=workflow%3Ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/ezh/wireguard-grpc)](https://goreportcard.com/report/github.com/ezh/wireguard-grpc)
[![codecov](https://codecov.io/gh/ezh/wireguard-grpc/branch/main/graph/badge.svg)](https://codecov.io/gh/ezh/wireguard-grpc)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc?ref=badge_shield)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

It is a secure implementation. Exposing private keys through a network is prohibited explicitly.

As a standalone application, it implements the following commands:
* diag: Test wireguard-grpc configuration.
* env: Prints environment variables.
* server: Run GRPC server. It provides the next API:
  - Ping: returns empty result if configuration is valid
  - Dump: returns a list of devices with all endpoints
* version: Prints the program’s version information.

As an embedded library:
* it can be seamlessly integrated with any logging implementation supported by [logr](https://github.com/go-logr/logr), like:
  - **a function** (can bridge to non-structured libraries): [funcr](https://github.com/go-logr/logr/tree/master/funcr)
  - **a testing.T** (for use in Go tests, with JSON-like output): [testr](https://github.com/go-logr/logr/tree/master/testr)
  - **github.com/google/glog**: [glogr](https://github.com/go-logr/glogr)
  - **k8s.io/klog** (for Kubernetes): [klogr](https://git.k8s.io/klog/klogr)
  - **a testing.T** (with klog-like text output): [ktesting](https://git.k8s.io/klog/ktesting)
  - **go.uber.org/zap**: [zapr](https://github.com/go-logr/zapr)
  - **log** (the Go standard library logger): [stdr](https://github.com/go-logr/stdr)
  - **github.com/sirupsen/logrus**: [logrusr](https://github.com/bombsimon/logrusr)
  - **github.com/wojas/genericr**: [genericr](https://github.com/wojas/genericr) (makes it easy to implement your own backend)
  - **logfmt** (Heroku style [logging](https://www.brandur.org/logfmt)): [logfmtr](https://github.com/iand/logfmtr)
  - **github.com/rs/zerolog**: [zerologr](https://github.com/go-logr/zerologr)
  - **github.com/go-kit/log**: [gokitlogr](https://github.com/tonglil/gokitlogr) (also compatible with github.com/go-kit/kit/log since v0.12.0)
  - **bytes.Buffer** (writing to a buffer): [bufrlogr](https://github.com/tonglil/buflogr) (useful for ensuring values were logged, like during testing)
  - more?
* it provides mock for [api/wireguard/v1/WireGuardServiceClient](api/wireguard/v1/wireguard_service.proto)


# Quick start

`wireguard-grpc server`

# Generate

`go generate -v`

# Test

## Unit tests

`ginkgo ./...`

## Integration tests

`ginkgo -tags=integration -v test/...`

# Build

`go build cmd/main.go`

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc?ref=badge_large)
