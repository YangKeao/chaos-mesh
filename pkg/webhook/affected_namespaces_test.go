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

package webhook

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha2"
)

func TestAffectedNamespaces(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	_, namespaces := affectedNamespaces(&v1alpha2.Schedule{
		Spec: v1alpha2.ScheduleSpec{
			ScheduleItem: v1alpha2.ScheduleItem{
				EmbedChaos: v1alpha2.EmbedChaos{
					PodChaos: &v1alpha2.PodChaosSpec{
						ContainerSelector: v1alpha2.ContainerSelector{
							PodSelector: v1alpha2.PodSelector{
								Selector: v1alpha2.PodSelectorSpec{
									GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
										Namespaces: []string{"ns1", "ns2"},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	g.Expect(namespaces).To(gomega.Equal(map[string]struct{}{
		"ns1": {},
		"ns2": {},
	}))

	_, namespaces = affectedNamespaces(&v1alpha2.Workflow{
		Spec: v1alpha2.WorkflowSpec{
			Templates: []v1alpha2.Template{
				{
					EmbedChaos: &v1alpha2.EmbedChaos{
						NetworkChaos: &v1alpha2.NetworkChaosSpec{
							Target: &v1alpha2.PodSelector{
								Selector: v1alpha2.PodSelectorSpec{
									GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
										Namespaces: []string{"ns1", "ns2"},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	g.Expect(namespaces).To(gomega.Equal(map[string]struct{}{
		"ns1": {},
		"ns2": {},
	}))

	clusterScoped, _ := affectedNamespaces(&v1alpha2.NetworkChaos{
		Spec: v1alpha2.NetworkChaosSpec{
			Target: &v1alpha2.PodSelector{},
		},
	})
	g.Expect(clusterScoped).To(gomega.BeTrue())

	clusterScoped, _ = affectedNamespaces(&v1alpha2.NetworkChaos{})
	g.Expect(clusterScoped).To(gomega.BeTrue())

	_, namespaces = affectedNamespaces(&v1alpha2.Workflow{
		Spec: v1alpha2.WorkflowSpec{
			Templates: []v1alpha2.Template{
				{
					EmbedChaos: &v1alpha2.EmbedChaos{
						NetworkChaos: &v1alpha2.NetworkChaosSpec{
							Target: &v1alpha2.PodSelector{
								Selector: v1alpha2.PodSelectorSpec{
									GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
										Namespaces: []string{"ns1", "ns2"},
									},
								},
							},
						},
					},
				},
				{
					EmbedChaos: &v1alpha2.EmbedChaos{
						NetworkChaos: &v1alpha2.NetworkChaosSpec{
							Target: &v1alpha2.PodSelector{
								Selector: v1alpha2.PodSelectorSpec{
									GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
										Namespaces: []string{"ns3", "ns4"},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	g.Expect(namespaces).To(gomega.Equal(map[string]struct{}{
		"ns1": {},
		"ns2": {},
		"ns3": {},
		"ns4": {},
	}))

	_, namespaces = affectedNamespaces(&v1alpha2.NetworkChaos{
		Spec: v1alpha2.NetworkChaosSpec{
			Target: &v1alpha2.PodSelector{
				Selector: v1alpha2.PodSelectorSpec{
					Pods: map[string][]string{
						"ns1": {"pod1", "pod2"},
						"ns2": {"pod3", "pod4"},
					},
				},
			},
		},
	})
	g.Expect(namespaces).To(gomega.Equal(map[string]struct{}{
		"ns1": {},
		"ns2": {},
	}))
}
