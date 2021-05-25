// Copyright 2021 Chaos Mesh Authors.
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

package recorder

type FinalizerInited struct {
}

func (p FinalizerInited) Type() string {
	return "Normal"
}

func (p FinalizerInited) Reason() string {
	return "FinalizerInited"
}

func (p FinalizerInited) Message() string {
	return "Finalizer has been inited"
}

func (p FinalizerInited) Parse(message string) ChaosEvent {
	if message == "Finalizer has been inited" {
		return FinalizerInited{}
	}

	return nil
}

type FinalizerRemoved struct {
}

func (p FinalizerRemoved) Type() string {
	return "Normal"
}

func (p FinalizerRemoved) Reason() string {
	return "FinalizerInited"
}

func (p FinalizerRemoved) Message() string {
	return "Finalizer has been removed"
}

func init() {
	register(FinalizerInited{}, FinalizerRemoved{})
}