// Copyright 2021 Chaos Mesh Authors.
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
//

package v1alpha2

import (
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="action",type=string,JSONPath=`.spec.action`
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment
// NetworkChaos is the Schema for the networkchaos API
type NetworkChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod chaos experiment
	Spec v1alpha1.NetworkChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the chaos experiment about pods
	Status v1alpha1.NetworkChaosStatus `json:"status"`
}

func (*NetworkChaos) Hub() {}

var _ v1alpha1.InnerObjectWithCustomStatus = (*NetworkChaos)(nil)
var _ v1alpha1.InnerObjectWithSelector = (*NetworkChaos)(nil)
var _ v1alpha1.InnerObject = (*NetworkChaos)(nil)

func (obj *NetworkChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".":       &obj.Spec.PodSelector,
		".Target": obj.Spec.Target,
	}
}

func (obj *NetworkChaos) GetCustomStatus() interface{} {
	return &obj.Status.Instances
}
