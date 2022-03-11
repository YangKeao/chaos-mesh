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

package detachvolume

import (
	"context"
	"encoding/json"

	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha2"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type Impl struct {
	client.Client

	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha2.Record, obj v1alpha2.InnerObject) (v1alpha2.Phase, error) {
	awschaos := obj.(*v1alpha2.AWSChaos)

	opts := []func(*awscfg.LoadOptions) error{
		awscfg.WithRegion(awschaos.Spec.AWSRegion),
	}
	if awschaos.Spec.SecretName != nil {
		secret := &v1.Secret{}
		err := impl.Client.Get(ctx, types.NamespacedName{
			Name:      *awschaos.Spec.SecretName,
			Namespace: awschaos.Namespace,
		}, secret)
		if err != nil {
			impl.Log.Error(err, "fail to get cloud secret")
			return v1alpha2.NotInjected, err
		}
		opts = append(opts, awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			string(secret.Data["aws_access_key_id"]),
			string(secret.Data["aws_secret_access_key"]),
			"",
		)))
	}
	cfg, err := awscfg.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		impl.Log.Error(err, "unable to load aws SDK config")
		return v1alpha2.NotInjected, err
	}
	ec2client := ec2.NewFromConfig(cfg)

	var selected v1alpha2.AWSSelector
	json.Unmarshal([]byte(records[index].Id), &selected)
	_, err = ec2client.DetachVolume(context.TODO(), &ec2.DetachVolumeInput{
		VolumeId:   selected.EbsVolume,
		Device:     selected.DeviceName,
		Force:      true,
		InstanceId: &selected.Ec2Instance,
	})

	if err != nil {
		impl.Log.Error(err, "fail to detach the volume")
		return v1alpha2.NotInjected, err
	}

	return v1alpha2.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha2.Record, obj v1alpha2.InnerObject) (v1alpha2.Phase, error) {
	awschaos := obj.(*v1alpha2.AWSChaos)

	opts := []func(*awscfg.LoadOptions) error{
		awscfg.WithRegion(awschaos.Spec.AWSRegion),
	}
	if awschaos.Spec.SecretName != nil {
		secret := &v1.Secret{}
		err := impl.Client.Get(ctx, types.NamespacedName{
			Name:      *awschaos.Spec.SecretName,
			Namespace: awschaos.Namespace,
		}, secret)
		if err != nil {
			impl.Log.Error(err, "fail to get cloud secret")
			return v1alpha2.Injected, err
		}
		opts = append(opts, awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			string(secret.Data["aws_access_key_id"]),
			string(secret.Data["aws_secret_access_key"]),
			"",
		)))
	}
	cfg, err := awscfg.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		impl.Log.Error(err, "unable to load aws SDK config")
		return v1alpha2.Injected, err
	}
	ec2client := ec2.NewFromConfig(cfg)

	var selected v1alpha2.AWSSelector
	json.Unmarshal([]byte(records[index].Id), &selected)

	_, err = ec2client.AttachVolume(context.TODO(), &ec2.AttachVolumeInput{
		Device:     selected.DeviceName,
		InstanceId: &selected.Ec2Instance,
		VolumeId:   selected.EbsVolume,
	})

	if err != nil {
		impl.Log.Error(err, "fail to attach the volume")
		return v1alpha2.Injected, err
	}

	return v1alpha2.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *Impl {
	return &Impl{
		Client: c,
		Log:    log.WithName("detachvolume"),
	}
}
