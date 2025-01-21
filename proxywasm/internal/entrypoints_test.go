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

import (
	"testing"

	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/stretchr/testify/require"
)

type testPluginContext struct {
	types.DefaultPluginContext
	contextID uint32
}

func TestSetPluginContext_ReturnsPluginContext(t *testing.T) {
	SetPluginContext(func(contextID uint32) types.PluginContext {
		return &testPluginContext{contextID: contextID}
	})

	pluginContext := currentState.vmContext.NewPluginContext(4321)

	require.IsType(t, pluginContext, &testPluginContext{})
	require.Equal(t, pluginContext.(*testPluginContext).contextID, uint32(4321))
}

type testPluginContextB struct {
	types.DefaultPluginContext
}

func TestSetPluginContext_Reentrant(t *testing.T) {
	SetPluginContext(func(uint32) types.PluginContext {
		return &testPluginContext{}
	})
	SetPluginContext(func(uint32) types.PluginContext {
		return &testPluginContextB{}
	})

	require.IsType(t, currentState.vmContext.NewPluginContext(1), &testPluginContextB{})
}

type (
	testHttpContext struct {
		types.DefaultHttpContext
		contextID uint32
	}
	testTcpContext struct{ types.DefaultTcpContext }
)

func TestSetHttpContext(t *testing.T) {
	SetHttpContext(func(contextID uint32) types.HttpContext {
		return &testHttpContext{contextID: contextID}
	})

	pluginContext := currentState.vmContext.NewPluginContext(4321).NewHttpContext(1234)

	require.IsType(t, pluginContext, &testHttpContext{})
	require.Equal(t, pluginContext.(*testHttpContext).contextID, uint32(1234))
}

func TestSetTcpContext(t *testing.T) {
	SetTcpContext(func(uint32) types.TcpContext {
		return &testTcpContext{}
	})

	pluginContext := currentState.vmContext.NewPluginContext(4321).NewTcpContext(1234)

	require.IsType(t, pluginContext, &testTcpContext{})
}

func TestSetTcpContext_ClearsSetHttpContext(t *testing.T) {
	SetHttpContext(func(contextID uint32) types.HttpContext {
		return &testHttpContext{contextID: contextID}
	})
	SetTcpContext(func(uint32) types.TcpContext {
		return &testTcpContext{}
	})

	pluginContext := currentState.vmContext.NewPluginContext(4321).NewTcpContext(1234)

	require.IsType(t, pluginContext, &testTcpContext{})
}
