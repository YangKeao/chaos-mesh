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

package config

// Error means fail to inject config
// +thaterror:wrap=*NotFound
// +thaterror:wrap=*MissingTemplateName
// +thaterror:wrap=*MissingName
// +thaterror:wrap="pkg/commonerror".*IOError
// +thaterror:wrap="pkg/commonerror".*UnmarshalError
type Error struct {
	Err error
}

// NotFound means injection config not found
// +thaterror:error=injection config not found  {{if not .Key}}for key {{.Key}} {{end}} at namespace {{.Namespace}}
type NotFound struct {
	Key       string
	Namespace string
}

// +thaterror:error=template field is required for template args config
type MissingTemplateName struct{}

// +thaterror:error=name field is required for template args config
type MissingName struct{}
