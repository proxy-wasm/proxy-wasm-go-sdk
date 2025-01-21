// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import "github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"

// elidedVmContext is registered when the plugin author uses a
// SetPluginContext, SetHttpContext, or SetTcpContext entrypoint. It indicates
// the author did not register a VmContext.
//
// elidedVmContext's primary responsibility is calling the author-provided (or
// elided) thunk to create a new PluginContext.
type elidedVmContext struct {
	types.DefaultVMContext
	newPluginContext types.PluginContextFactory
}

func (ctx *elidedVmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return ctx.newPluginContext(contextID)
}

// elidedPluginContext is registered when the plugin author uses the
// SetHttpContext or SetTcpContext entrypoints. It indicates the author did not
// register a VmContext or PluginContext.
//
// elidedVmContext's primary responsibility is calling the author-provided (or
// elided) thunk to create a new HttpContext or TcpContext.
type elidedPluginContext struct {
	types.DefaultPluginContext
	newHttpContext types.HttpContextFactory
	newTcpContext  types.TcpContextFactory
}

func (ctx *elidedPluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return ctx.newHttpContext(contextID)
}

func (ctx *elidedPluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return ctx.newTcpContext(contextID)
}

func SetPluginContext(newPluginContext types.PluginContextFactory) {
	SetVMContext(&elidedVmContext{newPluginContext: newPluginContext})
}

func SetHttpContext(newHttpContext types.HttpContextFactory) {
	SetVMContext(&elidedVmContext{
		newPluginContext: func(uint32) types.PluginContext {
			return &elidedPluginContext{newHttpContext: newHttpContext}
		},
	})
}

func SetTcpContext(newTcpContext types.TcpContextFactory) {
	SetVMContext(&elidedVmContext{
		newPluginContext: func(uint32) types.PluginContext {
			return &elidedPluginContext{newTcpContext: newTcpContext}
		},
	})
}
