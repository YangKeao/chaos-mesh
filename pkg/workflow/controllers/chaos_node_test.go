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

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha2"
)

// integration tests
var _ = Describe("Workflow", func() {
	var ns string
	BeforeEach(func() {
		ctx := context.TODO()
		newNs := corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "chaos-mesh-",
			},
			Spec: corev1.NamespaceSpec{},
		}
		Expect(kubeClient.Create(ctx, &newNs)).To(Succeed())
		ns = newNs.Name
		By(fmt.Sprintf("create new namespace %s", ns))
	})

	AfterEach(func() {
		ctx := context.TODO()
		nsToDelete := corev1.Namespace{}
		Expect(kubeClient.Get(ctx, types.NamespacedName{Name: ns}, &nsToDelete)).To(Succeed())
		Expect(kubeClient.Delete(ctx, &nsToDelete)).To(Succeed())
		By(fmt.Sprintf("cleanup namespace %s", ns))
	})

	Context("one chaos node", func() {
		It("could spawn one chaos", func() {
			ctx := context.TODO()
			now := time.Now()
			duration := 5 * time.Second

			By("create simple chaos node with pod chaos")
			startTime := metav1.NewTime(now)
			deadline := metav1.NewTime(now.Add(duration))
			workflowNode := v1alpha2.WorkflowNode{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns,
					GenerateName: "chaos-node-with-chaos-",
				},
				Spec: v1alpha2.WorkflowNodeSpec{
					TemplateName: "",
					WorkflowName: "",
					Type:         v1alpha2.TypePodChaos,
					StartTime:    &startTime,
					Deadline:     &deadline,
					EmbedChaos: &v1alpha2.EmbedChaos{
						PodChaos: &v1alpha2.PodChaosSpec{
							ContainerSelector: v1alpha2.ContainerSelector{
								PodSelector: v1alpha2.PodSelector{
									Selector: v1alpha2.PodSelectorSpec{
										GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
											Namespaces: []string{ns},
										},
									},
									Mode: v1alpha2.AllMode,
								},
							},
							Action: v1alpha2.PodKillAction,
						},
					},
				},
			}
			Expect(kubeClient.Create(ctx, &workflowNode)).To(Succeed())
			Eventually(func() bool {
				podChaosList := v1alpha2.PodChaosList{}
				Expect(kubeClient.List(ctx, &podChaosList, &client.ListOptions{Namespace: ns})).To(Succeed())
				if len(podChaosList.Items) == 0 {
					return false
				}
				return strings.HasPrefix(podChaosList.Items[0].Name, "chaos-node-with-chaos-")
			}, 10*time.Second, time.Second).Should(BeTrue())
		})

		It("could spawn one schedule", func() {
			ctx := context.TODO()
			now := time.Now()
			duration := 5 * time.Second

			By("create simple chaos node with schedule")
			startTime := metav1.NewTime(now)
			deadline := metav1.NewTime(now.Add(duration))
			node := v1alpha2.WorkflowNode{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:    ns,
					GenerateName: "chaos-node-schedule-",
				},
				Spec: v1alpha2.WorkflowNodeSpec{
					WorkflowName: "",
					Type:         v1alpha2.TypeSchedule,
					StartTime:    &startTime,
					Deadline:     &deadline,
					Schedule: &v1alpha2.ScheduleSpec{
						Schedule:                "@every 1s",
						StartingDeadlineSeconds: nil,
						ConcurrencyPolicy:       v1alpha2.AllowConcurrent,
						HistoryLimit:            5,
						Type:                    v1alpha2.ScheduleTypePodChaos,
						ScheduleItem: v1alpha2.ScheduleItem{
							EmbedChaos: v1alpha2.EmbedChaos{
								PodChaos: &v1alpha2.PodChaosSpec{
									ContainerSelector: v1alpha2.ContainerSelector{
										PodSelector: v1alpha2.PodSelector{
											Selector: v1alpha2.PodSelectorSpec{
												GenericSelectorSpec: v1alpha2.GenericSelectorSpec{
													Namespaces: []string{ns},
													LabelSelectors: map[string]string{
														"app": "not-actually-exist",
													},
												},
											},
											Mode: v1alpha2.AllMode,
										},
										ContainerNames: nil,
									},
									Action: v1alpha2.PodKillAction,
								},
							},
						},
					},
				},
			}
			Expect(kubeClient.Create(ctx, &node)).To(Succeed())
			Eventually(func() bool {
				scheduleList := v1alpha2.ScheduleList{}
				Expect(kubeClient.List(ctx, &scheduleList, &client.ListOptions{Namespace: ns})).To(Succeed())
				if len(scheduleList.Items) == 0 {
					return false
				}
				return strings.HasPrefix(scheduleList.Items[0].Name, "chaos-node-schedule-")
			}, 10*time.Second, time.Second).Should(BeTrue())
		})
	})
})
