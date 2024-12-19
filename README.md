# WebAssembly for Proxies (Go SDK) [![Build](https://github.com/proxy-wasm/proxy-wasm-go-sdk/workflows/Test/badge.svg)](https://github.com/proxy-wasm/proxy-wasm-go-sdk/actions) [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

The Go SDK for [Proxy-Wasm](https://github.com/proxy-wasm/spec), enabling
developers to write Proxy-Wasm plugins in Go.

Note: This SDK targets the upstream Go compiler, version 1.24 or later. This is
different than the forked github.com/tetratelabs/proxy-wasm-go-sdk, which
targets the TinyGo compiler.

## Project Status

This SDK is based off of github.com/tetratelabs/proxy-wasm-go-sdk; however, it
is effectively a new SDK targeting a completely different toolchain. It relies
on the not-yet-released Go 1.24 and hasn't seen extensive prod testing by
end-users. This SDK is an alpha product.

## Getting Started

-   [examples](examples) directory contains the example codes on top of this
    SDK.
-   [OVERVIEW.md](doc/OVERVIEW.md) the overview of Proxy-Wasm, the API of this
    SDK, and the things you should know when writing plugins.

## Requirements

-   \[Required] [Go](https://go.dev/): v1.24+ - This SDK leverages Go 1.24's
    support for [WASI](https://github.com/WebAssembly/WASI) (WebAssembly System
    Interface) reactors. You can grab a release candidate from the
    [Go unstable releases page](https://go.dev/dl/#unstable). A stable release
    of Go 1.24 is
    [expected in February 2025](https://tip.golang.org/doc/go1.24).
-   \[Required] A host environment supporting this toolchain - This SDK
    leverages additional host imports added to the proxy-wasm-cpp-host in
    [PR#427](https://github.com/proxy-wasm/proxy-wasm-cpp-host/pull/427). This
    has yet to be merged, let alone make its way downstream to Envoy, ATS,
    nginx, or managed environments.
-   \[Optional] [Envoy](https://www.envoyproxy.io) - To run end-to-end tests,
    you need to have an Envoy binary. You can use [func-e](https://func-e.io) as
    an easy way to get started with Envoy or follow
    [the official instruction](https://www.envoyproxy.io/docs/envoy/latest/start/install).

## Installation

```
go get github.com/proxy-wasm/proxy-wasm-go-sdk
```

## Minimal Example Plugin

A minimal `main.go` plugin defining VM and Plugin contexts:

```go
package main

import (
    "github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm"
    "github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {}
func init() {
    proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
    types.DefaultVMContext
}
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
    return &pluginContext{}
}

type pluginContext struct {
    types.DefaultPluginContext
}
// pluginContext should implement OnPluginStart, NewHttpContext, NewTcpContext, etc
```

Compile the plugin as follows:

```bash
$ go mod init main.go
$ go mod tidy
$ env GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o my-plugin.wasm main.go
```

## Build and run Examples

```bash
# Build all examples.
make build.examples

# Build a specific example using the `go` tool.
cd examples/helloworld
env GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o main.wasm main.go

# Test the built wasm module using `go test`.
go test

# Build and test a specific example from the repo root using `make`.
make build.example name=helloworld
make run name=helloworld
```

## Compatible Envoy builds

Envoy is the first host side implementation of Proxy-Wasm ABI, and we run
end-to-end tests with multiple versions of Envoy and Envoy-based
[istio/proxy](https://github.com/istio/proxy) in order to verify Proxy-Wasm Go
SDK works as expected.

Please refer to [workflow.yaml](.github/workflows/workflow.yaml) for which
version is used for End-to-End tests.

## Build tags

The following build tags can be used to customize the behavior of the built
plugin. Build tags
[can be specified in the `go build` command via the `-tags` flag](https://pkg.go.dev/cmd/go#:~:text=tags):

-   `proxywasm_timing`: Enables logging of time spent in invocation of the
    plugin's exported functions. This can be useful for debugging performance
    issues.

## Contributing

We welcome contributions from the community! See
[CONTRIBUTING.md](doc/CONTRIBUTING.md) for how to contribute to this repository.

## External links

-   [WebAssembly for Proxies (ABI specification)](https://github.com/proxy-wasm/spec)
-   [WebAssembly for Proxies (C++ SDK)](https://github.com/proxy-wasm/proxy-wasm-cpp-sdk)
-   [WebAssembly for Proxies (Rust SDK)](https://github.com/proxy-wasm/proxy-wasm-rust-sdk)
-   [WebAssembly for Proxies (AssemblyScript SDK)](https://github.com/solo-io/proxy-runtime)
