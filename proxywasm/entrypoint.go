// Copyright 2020-2024 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxywasm

import (
	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/internal"
	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"
)

// SetVMContext is one possible entrypoint for setting up the entire Wasm VM.
//
// Subsequent calls to any entrypoint overwrite previous calls to any
// entrypoint. Be sure to call exactly one entrypoint during `init()`,
// otherwise the VM fails.
func SetVMContext(ctx types.VMContext) {
	internal.SetVMContext(ctx)
}

// SetPluginContext is one possible entrypoint for setting up the Wasm VM.
//
// Subsequent calls to any entrypoint overwrite previous calls to any
// entrypoint. Be sure to call exactly one entrypoint during `init()`,
// otherwise the VM fails.
//
// Using SetPluginContext instead of SetVmContext is suitable iff the plugin
// does not make use of the VM configuration provided during `VmContext`'s
// `OnVmStart` call (plugin configuration data is still provided during
// `PluginContext`'s `OnPluginStart` call).
func SetPluginContext(newPluginContext func(contextID uint32) types.PluginContext) {
	internal.SetPluginContext(newPluginContext)
}

// SetHttpContext is one possible entrypoint for setting up the Wasm VM. It
// allows plugin authors to provide an Http context implementation without
// writing a VmContext or PluginContext.
//
// Subsequent calls to any entrypoint overwrite previous calls to any
// entrypoint. Be sure to call exactly one entrypoint during `init()`,
// otherwise the VM fails.
//
// SetHttpContext is suitable for stateless plugins that share no state between
// HTTP requests, do not process TCP streams, have no expensive shared setup
// requiring execution during `OnPluginStart`, and do not access the plugin
// configuration data.
func SetHttpContext(newHttpContext func(contextID uint32) types.HttpContext) {
	internal.SetHttpContext(newHttpContext)
}

// SetTcpContext is one possible entrypoint for setting up the Wasm VM. It
// allows plugin authors to provide a TCP context implementation without
// writing a VmContext or PluginContext.
//
// Subsequent calls to any entrypoint overwrite previous calls to any
// entrypoint. Be sure to call exactly one entrypoint during `init()`,
// otherwise the VM fails.
//
// SetTcpContext is suitable for stateless plugins that share no state between
// TCP streams, do not process HTTP requests, have no expensive shared setup
// requiring execution during `OnPluginStart`, and do not access the plugin
// configuration data.
func SetTcpContext(newTcpContext func(contextID uint32) types.TcpContext) {
	internal.SetTcpContext(newTcpContext)
}
