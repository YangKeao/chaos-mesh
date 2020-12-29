// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package timer

// Error represents error while getting time from the timer
// +thaterror:error=failed to get time from timer output: Err: {{.Err}}
// +thaterror:wrap=*ParseIntFailed
// +thaterror:wrap=*IOError
// +thaterror:wrap=*ProcessStartError
type Error struct {
	Err error
}

// ParseIntFailed represents error from `ParseInt`
// +thaterror:error=failed to parse int: S: {{.S}}, Base: {{.Base}}, Bitsize: {{.Bitsize}}
type ParseIntFailed struct {
	S       string
	Base    int
	Bitsize int
}

// IOError represents error while reading/scanning from the output
// +thaterror:transparent
type IOError struct {
	Err error
}

// ProcessStartError represents error while starting the process and getting the stdout
// +thaterror:transparent
type ProcessStartError struct {
	Err error
}
