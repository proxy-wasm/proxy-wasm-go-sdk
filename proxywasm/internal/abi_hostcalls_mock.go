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

//go:build !wasm

// TODO: Auto generate this file from abi_hostcalls.go.

package internal

import (
	"sync"
	"unsafe"
)

var (
	currentHost ProxyWasmHost
	mutex       = &sync.Mutex{}
)

func RegisterMockWasmHost(host ProxyWasmHost) (release func()) {
	mutex.Lock()
	currentHost = host
	return func() {
		mutex.Unlock()
	}
}

type ProxyWasmHost interface {
	ProxyLog(logLevel LogLevel, messageData *byte, messageSize int32) Status
	ProxySetProperty(pathData *byte, pathSize int32, valueData *byte, valueSize int32) Status
	ProxyGetProperty(pathData *byte, pathSize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status
	ProxySendLocalResponse(statusCode uint32, statusCodeDetailData *byte, statusCodeDetailsSize int32, bodyData *byte, bodySize int32, headersData *byte, headersSize int32, grpcStatus int32) Status
	ProxyGetSharedData(keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32, returnCas *uint32) Status
	ProxySetSharedData(keyData *byte, keySize int32, valueData *byte, valueSize int32, cas uint32) Status
	ProxyRegisterSharedQueue(nameData *byte, nameSize int32, returnID *uint32) Status
	ProxyResolveSharedQueue(vmIDData *byte, vmIDSize int32, nameData *byte, nameSize int32, returnID *uint32) Status
	ProxyDequeueSharedQueue(queueID uint32, returnValueData unsafe.Pointer, returnValueSize *int32) Status
	ProxyEnqueueSharedQueue(queueID uint32, valueData *byte, valueSize int32) Status
	ProxyGetHeaderMapValue(mapType MapType, keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status
	ProxyAddHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status
	ProxyReplaceHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status
	ProxyContinueStream(streamType StreamType) Status
	ProxyCloseStream(streamType StreamType) Status
	ProxyRemoveHeaderMapValue(mapType MapType, keyData *byte, keySize int32) Status
	ProxyGetHeaderMapPairs(mapType MapType, returnValueData unsafe.Pointer, returnValueSize *int32) Status
	ProxySetHeaderMapPairs(mapType MapType, mapData *byte, mapSize int32) Status
	ProxyGetBufferBytes(bufferType BufferType, start int32, maxSize int32, returnBufferData unsafe.Pointer, returnBufferSize *int32) Status
	ProxySetBufferBytes(bufferType BufferType, start int32, maxSize int32, bufferData *byte, bufferSize int32) Status
	ProxyHttpCall(upstreamData *byte, upstreamSize int32, headerData *byte, headerSize int32, bodyData *byte, bodySize int32, trailersData *byte, trailersSize int32, timeout uint32, calloutIDPtr *uint32) Status
	ProxyCallForeignFunction(funcNamePtr *byte, funcNameSize int32, paramPtr *byte, paramSize int32, returnData unsafe.Pointer, returnSize *int32) Status
	ProxySetTickPeriodMilliseconds(period uint32) Status
	ProxySetEffectiveContext(contextID uint32) Status
	ProxyDone() Status
	ProxyDefineMetric(metricType MetricType, metricNameData *byte, metricNameSize int32, returnMetricIDPtr *uint32) Status
	ProxyIncrementMetric(metricID uint32, offset int64) Status
	ProxyRecordMetric(metricID uint32, value uint64) Status
	ProxyGetMetric(metricID uint32, returnMetricValue *uint64) Status
}

type DefaultProxyWAMSHost struct{}

var _ ProxyWasmHost = DefaultProxyWAMSHost{}

func (d DefaultProxyWAMSHost) ProxyLog(logLevel LogLevel, messageData *byte, messageSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySetProperty(pathData *byte, pathSize int32, valueData *byte, valueSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyGetProperty(pathData *byte, pathSize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySendLocalResponse(statusCode uint32, statusCodeDetailData *byte, statusCodeDetailsSize int32, bodyData *byte, bodySize int32, headersData *byte, headersSize int32, grpcStatus int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyGetSharedData(keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32, returnCas *uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySetSharedData(keyData *byte, keySize int32, valueData *byte, valueSize int32, cas uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyRegisterSharedQueue(nameData *byte, nameSize int32, returnID *uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyResolveSharedQueue(vmIDData *byte, vmIDSize int32, nameData *byte, nameSize int32, returnID *uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyDequeueSharedQueue(queueID uint32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyEnqueueSharedQueue(queueID uint32, valueData *byte, valueSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyGetHeaderMapValue(mapType MapType, keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyAddHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyReplaceHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyContinueStream(streamType StreamType) Status { return 0 }
func (d DefaultProxyWAMSHost) ProxyCloseStream(streamType StreamType) Status    { return 0 }
func (d DefaultProxyWAMSHost) ProxyRemoveHeaderMapValue(mapType MapType, keyData *byte, keySize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyGetHeaderMapPairs(mapType MapType, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySetHeaderMapPairs(mapType MapType, mapData *byte, mapSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyGetBufferBytes(bufferType BufferType, start int32, maxSize int32, returnBufferData unsafe.Pointer, returnBufferSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySetBufferBytes(bufferType BufferType, start int32, maxSize int32, bufferData *byte, bufferSize int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyHttpCall(upstreamData *byte, upstreamSize int32, headerData *byte, headerSize int32, bodyData *byte, bodySize int32, trailersData *byte, trailersSize int32, timeout uint32, calloutIDPtr *uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyCallForeignFunction(funcNamePtr *byte, funcNameSize int32, paramPtr *byte, paramSize int32, returnData unsafe.Pointer, returnSize *int32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxySetTickPeriodMilliseconds(period uint32) Status { return 0 }
func (d DefaultProxyWAMSHost) ProxySetEffectiveContext(contextID uint32) Status    { return 0 }
func (d DefaultProxyWAMSHost) ProxyDone() Status                                   { return 0 }
func (d DefaultProxyWAMSHost) ProxyDefineMetric(metricType MetricType, metricNameData *byte, metricNameSize int32, returnMetricIDPtr *uint32) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyIncrementMetric(metricID uint32, offset int64) Status {
	return 0
}
func (d DefaultProxyWAMSHost) ProxyRecordMetric(metricID uint32, value uint64) Status { return 0 }
func (d DefaultProxyWAMSHost) ProxyGetMetric(metricID uint32, returnMetricValue *uint64) Status {
	return 0
}

func ProxyLog(logLevel LogLevel, messageData *byte, messageSize int32) Status {
	return currentHost.ProxyLog(logLevel, messageData, messageSize)
}

func ProxySetProperty(pathData *byte, pathSize int32, valueData *byte, valueSize int32) Status {
	return currentHost.ProxySetProperty(pathData, pathSize, valueData, valueSize)
}

func ProxyGetProperty(pathData *byte, pathSize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return currentHost.ProxyGetProperty(pathData, pathSize, returnValueData, returnValueSize)
}

func ProxySendLocalResponse(statusCode uint32, statusCodeDetailData *byte,
	statusCodeDetailsSize int32, bodyData *byte, bodySize int32, headersData *byte, headersSize int32, grpcStatus int32) Status {
	return currentHost.ProxySendLocalResponse(statusCode,
		statusCodeDetailData, statusCodeDetailsSize, bodyData, bodySize, headersData, headersSize, grpcStatus)
}

func ProxyGetSharedData(keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32, returnCas *uint32) Status {
	return currentHost.ProxyGetSharedData(keyData, keySize, returnValueData, returnValueSize, returnCas)
}

func ProxySetSharedData(keyData *byte, keySize int32, valueData *byte, valueSize int32, cas uint32) Status {
	return currentHost.ProxySetSharedData(keyData, keySize, valueData, valueSize, cas)
}

func ProxyRegisterSharedQueue(nameData *byte, nameSize int32, returnID *uint32) Status {
	return currentHost.ProxyRegisterSharedQueue(nameData, nameSize, returnID)
}

func ProxyResolveSharedQueue(vmIDData *byte, vmIDSize int32, nameData *byte, nameSize int32, returnID *uint32) Status {
	return currentHost.ProxyResolveSharedQueue(vmIDData, vmIDSize, nameData, nameSize, returnID)
}

func ProxyDequeueSharedQueue(queueID uint32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return currentHost.ProxyDequeueSharedQueue(queueID, returnValueData, returnValueSize)
}

func ProxyEnqueueSharedQueue(queueID uint32, valueData *byte, valueSize int32) Status {
	return currentHost.ProxyEnqueueSharedQueue(queueID, valueData, valueSize)
}

func ProxyGetHeaderMapValue(mapType MapType, keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return currentHost.ProxyGetHeaderMapValue(mapType, keyData, keySize, returnValueData, returnValueSize)
}

func ProxyAddHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status {
	return currentHost.ProxyAddHeaderMapValue(mapType, keyData, keySize, valueData, valueSize)
}

func ProxyReplaceHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status {
	return currentHost.ProxyReplaceHeaderMapValue(mapType, keyData, keySize, valueData, valueSize)
}

func ProxyContinueStream(streamType StreamType) Status {
	return currentHost.ProxyContinueStream(streamType)
}

func ProxyCloseStream(streamType StreamType) Status {
	return currentHost.ProxyCloseStream(streamType)
}
func ProxyRemoveHeaderMapValue(mapType MapType, keyData *byte, keySize int32) Status {
	return currentHost.ProxyRemoveHeaderMapValue(mapType, keyData, keySize)
}

func ProxyGetHeaderMapPairs(mapType MapType, returnValueData unsafe.Pointer, returnValueSize *int32) Status {
	return currentHost.ProxyGetHeaderMapPairs(mapType, returnValueData, returnValueSize)
}

func ProxySetHeaderMapPairs(mapType MapType, mapData *byte, mapSize int32) Status {
	return currentHost.ProxySetHeaderMapPairs(mapType, mapData, mapSize)
}

func ProxyGetBufferBytes(bufferType BufferType, start int32, maxSize int32, returnBufferData unsafe.Pointer, returnBufferSize *int32) Status {
	return currentHost.ProxyGetBufferBytes(bufferType, start, maxSize, returnBufferData, returnBufferSize)
}

func ProxySetBufferBytes(bufferType BufferType, start int32, maxSize int32, bufferData *byte, bufferSize int32) Status {
	return currentHost.ProxySetBufferBytes(bufferType, start, maxSize, bufferData, bufferSize)
}

func ProxyHttpCall(upstreamData *byte, upstreamSize int32, headerData *byte, headerSize int32, bodyData *byte,
	bodySize int32, trailersData *byte, trailersSize int32, timeout uint32, calloutIDPtr *uint32) Status {
	return currentHost.ProxyHttpCall(upstreamData, upstreamSize,
		headerData, headerSize, bodyData, bodySize, trailersData, trailersSize, timeout, calloutIDPtr)
}

func ProxyCallForeignFunction(funcNamePtr *byte, funcNameSize int32, paramPtr *byte, paramSize int32, returnData unsafe.Pointer, returnSize *int32) Status {
	return currentHost.ProxyCallForeignFunction(funcNamePtr, funcNameSize, paramPtr, paramSize, returnData, returnSize)
}

func ProxySetTickPeriodMilliseconds(period uint32) Status {
	return currentHost.ProxySetTickPeriodMilliseconds(period)
}

func ProxySetEffectiveContext(contextID uint32) Status {
	return currentHost.ProxySetEffectiveContext(contextID)
}

func ProxyDone() Status {
	return currentHost.ProxyDone()
}

func ProxyDefineMetric(metricType MetricType,
	metricNameData *byte, metricNameSize int32, returnMetricIDPtr *uint32) Status {
	return currentHost.ProxyDefineMetric(metricType, metricNameData, metricNameSize, returnMetricIDPtr)
}

func ProxyIncrementMetric(metricID uint32, offset int64) Status {
	return currentHost.ProxyIncrementMetric(metricID, offset)
}

func ProxyRecordMetric(metricID uint32, value uint64) Status {
	return currentHost.ProxyRecordMetric(metricID, value)
}

func ProxyGetMetric(metricID uint32, returnMetricValue *uint64) Status {
	return currentHost.ProxyGetMetric(metricID, returnMetricValue)
}
