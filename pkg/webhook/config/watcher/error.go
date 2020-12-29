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

package watcher

// Error means fail to operate the fuse device
// +thaterror:error=fail to validate the parameter configuration
// +thaterror:wrap=*EmptyTemplateLabels
// +thaterror:wrap=*EmptyConfigMapLabels
// +thaterror:wrap=*EmptyTargetNamespace
type Error struct {
	Err error
}

// EmptyTemplateLabels represents that "TEMPLATE_LABELS" is empty
// +thaterror:error=envconfig:"TEMPLATE_LABELS" template labels must be set
type EmptyTemplateLabels struct{}

// EmptyConfigMapLabels represents that "CONFIGMAP_LABELS" is empty
// +thaterror:error=envconfig:"CONFIGMAP_LABELS" template labels must be set
type EmptyConfigMapLabels struct{}

// EmptyTargetNamespace represents that "TARGET_NAMESPACE" is empty
// +thaterror:error=envconfig:"TARGET_NAMESPACE" conf labels must be set while CLUSTER_SCOPED is false
type EmptyTargetNamespace struct{}
