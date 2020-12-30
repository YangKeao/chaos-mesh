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

package grpc

import v1 "k8s.io/api/core/v1"

// FailToFindEndpointIP means failed to find an ip for node in endpoints
// +thaterror:error=failed to find node {{.NodeName}} in {{.Endpoints}}
type FailToFindEndpointIP struct {
	NodeName  string
	Endpoints v1.Endpoints
}

// GrpcError represents an error returned by the grpc library
// +thaterror:transparent
type GrpcError struct {
	Err error
}

// CreateConnectionError represents all possible error while creating connection
// to chaos daemon
// +thaterror:transparent
// +thaterror:wrap=*FailToFindEndpointIP
// +thaterror:wrap=*GrpcError
// +thaterror:wrap="pkg/commonerror".*KubernetesError
type CreateConnectionError struct {
	Err error
}
