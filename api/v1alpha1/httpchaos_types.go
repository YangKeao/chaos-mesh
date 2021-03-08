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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +chaos-mesh:base

// HTTPChaos is the Schema for the HTTPchaos API
type HTTPChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPChaosSpec   `json:"spec,omitempty"`
	Status HTTPChaosStatus `json:"status,omitempty"`
}

// HTTPChaosAction represents the chaos action about HTTP.
type HTTPChaosAction string

const (
	HTTPDelayAction HTTPChaosAction = "delay"
	HTTPAbortAction                 = "abort"
	HTTPMixedAction                 = "mixed"
)

type Matcher struct {
	Name           string  `json:"name"`
	ExactMatch     *string `json:"exact_match,omitempty"`
	RegexMatch     *string `json:"regex_match,omitempty"`
	SafeRegexMatch *string `json:"safe_regex_match,omitempty"`
	RangeMatch     *string `json:"range_match,omitempty"`
	PresentMatch   *string `json:"present_match,omitempty"`
	PrefixMatch    *string `json:"prefix_match,omitempty"`
	SuffixMatch    *string `json:"suffix_match,omitempty"`
	InvertMatch    *string `json:"invert_match,omitempty"`
}

type HTTPChaosSpec struct {
	PodSelector `json:",inline"`

	// Action defines the specific pod chaos action.
	// Supported action: delay | abort | mixed
	// Default action: delay
	// +kubebuilder:validation:Enum=delay;abort;mixed
	Action HTTPChaosAction `json:"action"`

	// Duration represents the duration of the chaos action.
	// It is required when the action is `PodFailureAction`.
	// A duration string is a possibly signed sequence of
	// decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	Duration *string `json:"duration,omitempty"`

	// Percent defines the percentage of injection errors and provides a number from 0-100.
	// default: 100.
	// +optional
	Percent string `json:"percent,omitempty"`

	// Specifies how the header match will be performed to route the request.
	Headers []Matcher `json:"headers,omitempty"`
}

func (in *HTTPChaosSpec) GetHeaders() []Matcher {
	return in.Headers
}

type HTTPChaosStatus struct {
	ChaosStatus `json:",inline"`
}
