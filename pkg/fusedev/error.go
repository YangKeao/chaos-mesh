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

package fusedev

// Error means fail to operate the fuse device
// +thaterror:error=fail to operate the fuse device: Error: {{.Err.Error()}}
// +thaterror:wrap="pkg/commonerror".*IOError
// +thaterror:wrap=*InvalidCgroupEntry
// +thaterror:wrap=*DeviceCgroupNotFound
type Error struct {
	Err error
}

// InvalidCgroupEntry represents the cgroup entry is not expected
// +thaterror:error=invalid cgroup entry: {{.Text}}
type InvalidCgroupEntry struct {
	Text string
}

// DeviceCgroupNotFound means failed to find device cgroup
// +thaterror:error=fail to find device cgroup
type DeviceCgroupNotFound struct{}
