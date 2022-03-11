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

package controller

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha2"
)

func makeTestPodKill(creationTime time.Time, duration *string, desiredPhase v1alpha2.DesiredPhase, records []*v1alpha2.Record) v1alpha2.InnerObject {
	return &v1alpha2.PodChaos{
		ObjectMeta: v1.ObjectMeta{
			CreationTimestamp: v1.Time{
				Time: creationTime,
			},
		},
		Spec: v1alpha2.PodChaosSpec{
			Action:   v1alpha2.PodKillAction,
			Duration: duration,
		},
		Status: v1alpha2.PodChaosStatus{
			ChaosStatus: v1alpha2.ChaosStatus{
				Experiment: v1alpha2.ExperimentStatus{
					DesiredPhase: desiredPhase,
					Records:      records,
				},
			},
		},
	}
}
func makeTestNetworkChaos(creationTime time.Time, duration *string, desiredPhase v1alpha2.DesiredPhase, records []*v1alpha2.Record) v1alpha2.InnerObject {
	return &v1alpha2.NetworkChaos{
		ObjectMeta: v1.ObjectMeta{
			CreationTimestamp: v1.Time{
				Time: creationTime,
			},
		},
		Spec: v1alpha2.NetworkChaosSpec{
			Duration: duration,
		},
		Status: v1alpha2.NetworkChaosStatus{
			ChaosStatus: v1alpha2.ChaosStatus{
				Experiment: v1alpha2.ExperimentStatus{
					DesiredPhase: desiredPhase,
					Records:      records,
				},
			},
		},
	}
}

func TestIsChaosFinished(t *testing.T) {
	g := NewGomegaWithT(t)

	type testCase struct {
		chaos v1alpha2.InnerObject
		now   time.Time

		expected bool
	}

	beginTime := time.Now()
	cases := []testCase{
		{
			chaos: makeTestNetworkChaos(beginTime, pointer.StringPtr("20s"), v1alpha2.RunningPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.Injected,
				},
			}),
			now: beginTime.Add(10 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestNetworkChaos(beginTime, pointer.StringPtr("20s"), v1alpha2.RunningPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.NotInjected,
				},
			}),
			now: beginTime.Add(10 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestNetworkChaos(beginTime, pointer.StringPtr("20s"), v1alpha2.RunningPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.NotInjected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestNetworkChaos(beginTime, pointer.StringPtr("20s"), v1alpha2.StoppedPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.NotInjected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: true,
		},
		{
			chaos: makeTestNetworkChaos(beginTime, nil, v1alpha2.RunningPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.NotInjected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestNetworkChaos(beginTime, nil, v1alpha2.RunningPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.Injected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: false,
		},
		// The chaos is paused, but not recovered yet
		{
			chaos: makeTestNetworkChaos(beginTime, nil, v1alpha2.StoppedPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.Injected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestPodKill(beginTime, pointer.StringPtr("20s"), v1alpha2.StoppedPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.NotInjected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: false,
		},
		{
			chaos: makeTestPodKill(beginTime, pointer.StringPtr("20s"), v1alpha2.StoppedPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.Injected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: true,
		},
		{
			chaos: makeTestPodKill(beginTime, nil, v1alpha2.StoppedPhase, []*v1alpha2.Record{
				{
					Id:          "some",
					SelectorKey: "some",
					Phase:       v1alpha2.Injected,
				},
			}),
			now: beginTime.Add(30 * time.Second),

			expected: true,
		},
	}

	for index, c := range cases {
		if index == 5 {
			fmt.Println("some")
		}
		fmt.Println(index)
		g.Expect(IsChaosFinished(c.chaos, c.now)).To(Equal(c.expected))
	}
}
