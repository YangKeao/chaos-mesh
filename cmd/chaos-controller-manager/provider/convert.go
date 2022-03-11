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

package provider

import (
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"go.uber.org/fx"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

var oldObjs = []runtime.Object{
	&v1alpha1.AWSChaos{},
	&v1alpha1.DNSChaos{},
	&v1alpha1.HTTPChaos{},
	&v1alpha1.IOChaos{},
	&v1alpha1.KernelChaos{},
	&v1alpha1.JVMChaos{},
	&v1alpha1.NetworkChaos{},
	&v1alpha1.PodChaos{},
	&v1alpha1.StressChaos{},
	&v1alpha1.TimeChaos{},
	&v1alpha1.GCPChaos{},
	&v1alpha1.PhysicalMachineChaos{},
	&v1alpha1.BlockChaos{},
}

var RegisterConvert = fx.Invoke(func(mgr ctrl.Manager) error {
	for _, obj := range oldObjs {
		err := ctrl.NewWebhookManagedBy(mgr).
			For(obj).
			Complete()
		if err != nil {
			return err
		}
	}

	return nil
})
