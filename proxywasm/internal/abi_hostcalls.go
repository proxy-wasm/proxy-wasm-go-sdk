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

//go:build wasm

package internal

import "unsafe"

//go:wasmimport env proxy_log
func ProxyLog(logLevel LogLevel, messageData *byte, messageSize int32) Status

//go:wasmimport env proxy_send_local_response
func ProxySendLocalResponse(statusCode uint32, statusCodeDetailData *byte, statusCodeDetailsSize int32,
	bodyData *byte, bodySize int32, headersData *byte, headersSize int32, grpcStatus int32) Status

//go:wasmimport env proxy_get_shared_data
func ProxyGetSharedData(keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32, returnCas *uint32) Status

//go:wasmimport env proxy_set_shared_data
func ProxySetSharedData(keyData *byte, keySize int32, valueData *byte, valueSize int32, cas uint32) Status

//go:wasmimport env proxy_register_shared_queue
func ProxyRegisterSharedQueue(nameData *byte, nameSize int32, returnID *uint32) Status

//go:wasmimport env proxy_resolve_shared_queue
func ProxyResolveSharedQueue(vmIDData *byte, vmIDSize int32, nameData *byte, nameSize int32, returnID *uint32) Status

//go:wasmimport env proxy_dequeue_shared_queue
func ProxyDequeueSharedQueue(queueID uint32, returnValueData unsafe.Pointer, returnValueSize *int32) Status

//go:wasmimport env proxy_enqueue_shared_queue
func ProxyEnqueueSharedQueue(queueID uint32, valueData *byte, valueSize int32) Status

//go:wasmimport env proxy_get_header_map_value
func ProxyGetHeaderMapValue(mapType MapType, keyData *byte, keySize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status

//go:wasmimport env proxy_add_header_map_value
func ProxyAddHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status

//go:wasmimport env proxy_replace_header_map_value
func ProxyReplaceHeaderMapValue(mapType MapType, keyData *byte, keySize int32, valueData *byte, valueSize int32) Status

//go:wasmimport env proxy_remove_header_map_value
func ProxyRemoveHeaderMapValue(mapType MapType, keyData *byte, keySize int32) Status

//go:wasmimport env proxy_get_header_map_pairs
func ProxyGetHeaderMapPairs(mapType MapType, returnValueData unsafe.Pointer, returnValueSize *int32) Status

//go:wasmimport env proxy_set_header_map_pairs
func ProxySetHeaderMapPairs(mapType MapType, mapData *byte, mapSize int32) Status

//go:wasmimport env proxy_get_buffer_bytes
func ProxyGetBufferBytes(bufferType BufferType, start int32, maxSize int32, returnBufferData unsafe.Pointer, returnBufferSize *int32) Status

//go:wasmimport env proxy_set_buffer_bytes
func ProxySetBufferBytes(bufferType BufferType, start int32, maxSize int32, bufferData *byte, bufferSize int32) Status

//go:wasmimport env proxy_continue_stream
func ProxyContinueStream(streamType StreamType) Status

//go:wasmimport env proxy_close_stream
func ProxyCloseStream(streamType StreamType) Status

//go:wasmimport env proxy_http_call
func ProxyHttpCall(upstreamData *byte, upstreamSize int32, headerData *byte, headerSize int32,
	bodyData *byte, bodySize int32, trailersData *byte, trailersSize int32, timeout uint32, calloutIDPtr *uint32,
) Status

//go:wasmimport env proxy_call_foreign_function
func ProxyCallForeignFunction(funcNamePtr *byte, funcNameSize int32, paramPtr *byte, paramSize int32, returnData unsafe.Pointer, returnSize *int32) Status

//go:wasmimport env proxy_set_tick_period_milliseconds
func ProxySetTickPeriodMilliseconds(period uint32) Status

//go:wasmimport env proxy_set_effective_context
func ProxySetEffectiveContext(contextID uint32) Status

//go:wasmimport env proxy_done
func ProxyDone() Status

//go:wasmimport env proxy_define_metric
func ProxyDefineMetric(metricType MetricType, metricNameData *byte, metricNameSize int32, returnMetricIDPtr *uint32) Status

//go:wasmimport env proxy_increment_metric
func ProxyIncrementMetric(metricID uint32, offset int64) Status

//go:wasmimport env proxy_record_metric
func ProxyRecordMetric(metricID uint32, value uint64) Status

//go:wasmimport env proxy_get_metric
func ProxyGetMetric(metricID uint32, returnMetricValue *uint64) Status

//go:wasmimport env proxy_get_property
func ProxyGetProperty(pathData *byte, pathSize int32, returnValueData unsafe.Pointer, returnValueSize *int32) Status

//go:wasmimport env proxy_set_property
func ProxySetProperty(pathData *byte, pathSize int32, valueData *byte, valueSize int32) Status
