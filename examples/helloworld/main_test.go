// The framework emulates the expected behavior of Envoyproxy, and you can test your extensions without running Envoy and with
// the standard Go CLI. To run tests, simply run
// go test ./...

package main

import (
	"os"
	"testing"

	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/proxytest"
	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld_OnTick(t *testing.T) {
	vmTest(t, func(t *testing.T, pc types.PluginContextFactory) {
		opt := proxytest.NewEmulatorOption().WithPluginContext(pc)
		host, reset := proxytest.NewHostEmulator(opt)
		defer reset()

		// Call OnPluginStart.
		require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())
		require.Equal(t, tickMilliseconds, host.GetTickPeriod())

		// Call OnTick.
		host.Tick()

		// Check Envoy logs.
		logs := host.GetInfoLogs()
		require.Contains(t, logs, "OnTick called")
	})
}

func TestHelloWorld_OnPluginStart(t *testing.T) {
	vmTest(t, func(t *testing.T, pc types.PluginContextFactory) {
		opt := proxytest.NewEmulatorOption().WithPluginContext(pc)
		host, reset := proxytest.NewHostEmulator(opt)
		defer reset()

		// Call OnPluginStart.
		require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())

		// Check Envoy logs.
		logs := host.GetInfoLogs()
		require.Contains(t, logs, "OnPluginStart from Go!")
		require.Equal(t, tickMilliseconds, host.GetTickPeriod())
	})
}

// vmTest executes f twice, once with a types.VMContext that executes plugin code directly
// in the host, and again by executing the plugin code within the compiled main.wasm binary.
// Execution with main.wasm will be skipped if the file cannot be found.
func vmTest(t *testing.T, f func(*testing.T, types.PluginContextFactory)) {
	t.Helper()

	t.Run("go", func(t *testing.T) {
		f(t, func(uint32) types.PluginContext { return &helloWorld{} })
	})

	t.Run("wasm", func(t *testing.T) {
		wasm, err := os.ReadFile("main.wasm")
		if err != nil {
			t.Skip("wasm not found")
		}
		v, err := proxytest.NewWasmVMContext(wasm)
		require.NoError(t, err)
		defer v.Close()
		f(t, v.NewPluginContext)
	})
}
