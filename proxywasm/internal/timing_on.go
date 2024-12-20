// Copyright 2020-2022 Tetrate
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

//go:build proxywasm_timing

package internal

import (
	"fmt"
	"time"
)

// When the build tag is specified, we record timing information so set this to true.
const recordTiming = true

func logTiming(msg string, start time.Time) {
	if !recordTiming {
		panic("BUG: logTiming should not be called when timing is disabled")
	}
	f := fmt.Sprintf("%s took %s", msg, time.Since(start))
	ProxyLog(LogLevelDebug, StringBytePtr(f), int32(len(f)))
}
