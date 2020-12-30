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

package commonerror

import v1 "k8s.io/api/core/v1"

// ParseIntError represents error from `ParseInt`
// +thaterror:error=failed to parse int: S: {{.S}}, Base: {{.Base}}, Bitsize: {{.Bitsize}}
type ParseIntError struct {
	S       string
	Base    int
	Bitsize int
}

// ParseDurationError represents error from `time.ParseDuration`
// +thaterror:error=failed to parse duration: {{.S}}
type ParseDurationError struct {
	S string
}

// IOError represents error while io operation
// +thaterror:transparent
type IOError struct {
	Err error
}

// UnmarshalError represents error during unmarshal
// +thaterror:transparent
type UnmarshalError struct {
	Err error
}

// +thaterror:transparent
type KubernetesError struct {
	Err error
}

// ContainerStatusEmpty means the ContainerStatuses is empty
// +thaterror:error=the container statuses of {{.Name}}/{{.Namespace}} is empty
type ContainerStatusEmpty struct {
	Name      string
	Namespace string
	Pod       v1.Pod
}
