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

package chaosdaemon

// InvalidRequest means the content of the request is not valid
// +thaterror:error=grpc Request is invalid: Error: : {{.Err.Error()}}
// +thaterror:wrap=DNSServerShouldNotBeEmpty
// +thaterror:wrap=UnknownChainDirection
type InvalidRequest struct {
	Err error
}

// DNSServerShouldNotBeEmpty means dnschaos server gets a request with `DnsServer` field empty
// +thaterror:error="DnsServer field shouldn't be empty"
type DNSServerShouldNotBeEmpty struct{}

// BadRequest represents the grpc reqeust is invalid
// +thaterror:error=grpc request invalid: {{.Err.Error()}}
// +thaterror:wrap=EmptyDNSServerError
type BadRequest struct {
	Err error
}

// EmptyDNSServerError represents the DnsServer field is empty for a SetDNSServer grpc request
// +thaterror:error=DnsServer in the request is empty
type EmptyDNSServerError struct{}

// IptablesError represents the error in iptables service
// +thaterror:transparent
// +thaterror:wrap=CommandExecuteError
type IptablesError struct {
	Err error
}

// UnknownChainDirection represents the iptables server gets a direction
// other than Chain_INPUT and Chain_OUTPUT
// +thaterror:error=unknown chain direction {{.Direction}}
type UnknownChainDirection struct {
	Direction int32
}

// IPSetError represents the error in ipset service
// +thaterror:transparent
// +thaterror:wrap=CommandExecuteError
type IPSetError struct {
	Err error
}

// CommandExecuteError represents an error occured during the execution
// +thaterror:error:start
// Command {{.Command}} failed with return value {{.Ret}} and output:
// {{.Output}}
// +thaterror:error:stop
type CommandExecuteError struct {
	Ret     int32
	Output  string
	Command string
}

// ContainerClientError represents an error returned by container client
// +thaterror:transparent
// +thaterror:wrap=BadContainerState
// +thaterror:wrap=InvalidPidError
// +thaterror:wrap="github.com/YangKeao/thaterror/error".Anyhow
type ContainerClientError struct {
	Err error
}

// InvalidPidError means the container client gets an Pid 0
// +thaterror:error=invalid pid for container: {{.ContainerID}}
type InvalidPidError struct {
	ContainerID string
}

// BadContainerState means the container state is bad
// +thaterror:error=the status of container {{.ContainerID}} is {{.State}}
type BadContainerState struct {
	State       string
	ContainerID string
}
