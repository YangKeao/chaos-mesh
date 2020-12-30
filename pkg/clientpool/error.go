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

package clientpool

import (
	"k8s.io/client-go/rest"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// LRUError means the error returned by the `lru` pkg
// You cannot tell what the error in lru could be. However, if you look at
// the source code, you know that "Must provide a positive size" is the only
// possible error.
// `thaterror` is an attempt to solve the problem internally.
// +thaterror:error=fail to create lru with size: {{.Size}}. Err: {{.Err}}
type LRUError struct {
	Err  error
	Size int
}

// +thaterror:transparent
// +thaterror:wrap=*KubernetesCreateClientError
// +thaterror:wrap=*EmptyTokenError
type CreateClientError struct {
	Err error
}

// KubernetesCreateClientError means the error while creating a kubernetes client
// +thaterror:error=fail to create kubernetes client with config {{.Config}} and options {{.Options}}
type KubernetesCreateClientError struct {
	Err     error
	Config  *rest.Config
	Options pkgclient.Options
}

// EmptyTokenError means the token is empty
// +thaterror:error=token is empty
type EmptyTokenError struct{}
