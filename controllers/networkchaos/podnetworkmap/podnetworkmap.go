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

package podnetworkmap

import (
	"context"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PodNetworkMap will save all the related podnetworkchaos
type PodNetworkMap struct {
	Name string
	Log  logr.Logger
	client.Client
	Resources map[types.NamespacedName]*v1alpha1.PodNetworkChaos
}

// New creates a new PodNetworkMap
func New(name string, logger logr.Logger, client client.Client) *PodNetworkMap {
	return &PodNetworkMap{
		Name:      name,
		Log:       logger,
		Client:    client,
		Resources: make(map[types.NamespacedName]*v1alpha1.PodNetworkChaos),
	}
}

// Get will get podnetworkmap from kubernetes
func (m *PodNetworkMap) Get(ctx context.Context, key types.NamespacedName) (*v1alpha1.PodNetworkChaos, error) {
	chaos, ok := m.Resources[key]
	if ok {
		return chaos, nil
	}

	chaos = &v1alpha1.PodNetworkChaos{}
	err := m.Client.Get(ctx, key, chaos)
	if err != nil {
		return nil, err
	}

	m.Resources[key] = chaos

	return chaos, nil
}

// GetAndClear gets and removes all related resources if it was the first time to get it
func (m *PodNetworkMap) GetAndClear(ctx context.Context, key types.NamespacedName) (*v1alpha1.PodNetworkChaos, error) {
	chaos, ok := m.Resources[key]
	if ok {
		return chaos, nil
	}

	chaos, err := m.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	ipsets := []v1alpha1.RawIPSet{}
	for _, ipset := range chaos.Spec.IPSets {
		if ipset.Source != m.Name {
			ipsets = append(ipsets, ipset)
		}
	}
	chaos.Spec.IPSets = ipsets

	chains := []v1alpha1.RawIptables{}
	for _, chain := range chaos.Spec.Iptables {
		if chain.Source != m.Name {
			chains = append(chains, chain)
		}
	}
	chaos.Spec.Iptables = chains

	qdiscs := []v1alpha1.RawQdisc{}
	for _, qdisc := range chaos.Spec.Qdiscs {
		if qdisc.Source != m.Name {
			qdiscs = append(qdiscs, qdisc)
		}
	}
	chaos.Spec.Qdiscs = qdiscs

	return chaos, nil
}

// Commit will update all modifications to the cluster
func (m *PodNetworkMap) Commit(ctx context.Context) error {

	// TODO: parallel execution
	for key, val := range m.Resources {
		m.Log.Info("updating", "name", key.Namespace+"/"+key.Name)
		err := m.Client.Update(ctx, val)
		if err != nil {
			m.Log.Error(err, "error while updating item")
			return err
		}
	}

	return nil
}
