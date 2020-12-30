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

package ptrace

import "syscall"

// TraceError is the error in tracing a process
// +thaterror:transparent
// +thaterror:wrap=*PtraceError
// +thaterror:wrap=*WaitPidError
// +thaterror:wrap="pkg/commonerror".*IOError
// +thaterror:wrap="pkg/commonerror".*ParseIntError
// +thaterror:wrap="pkg/mapreader".*Error
type TraceError struct {
	Err error
}

// PtraceError represents an error returned by ptrace syscall
// +thaterror:error=fail to ptrace({{.Operation}}) on process {{.Tid}}
type PtraceError struct {
	Err       error
	Tid       int
	Operation int
}

// StepError represents an error in stepping forward
// +thaterror:transparent
// +thaterror:wrap=*WaitPidError
// +thaterror:wrap=*PtraceError
type StepError struct {
	Err error
}

// SyscallError represents an error in asking another process to run syscall
// +thaterror:transparent
// +thaterror:wrap=*PtraceError
// +thaterror:wrap=*StepError
// +thaterror:wrap=*TooManyArguments
type SyscallError struct {
	Err error
}

// TooManyArguments means the number of syscall arguments shouldn't be more than 6
// +thaterror:error=too many arguments for a syscall: {{.Arguments}}
type TooManyArguments struct {
	Arguments []uint64
}

// ProcessVMError represents an error returned by process_vm_readv or process_vm_writev
// +thaterror:error=process_vm_readv returned an error: {{.Errno}}
type ProcessVMError struct {
	Errno syscall.Errno
	Pid   int
}

// WaitPidError means the waitpid syscall returns an error
// +thaterror:error=waitpid failed
// TODO: get errno for this error
type WaitPidError struct{}

// ELFOperationError represents all error in an ELF operation
// +thaterror:transparent
// +thaterror:wrap=*ProcessVMError
// +thaterror:wrap=*EntryWithPaddingSize
// +thaterror:wrap=ELFError
// +thaterror:wrap=*FailToFindSymbol
type ELFOperationError struct {
	Err error
}

// EntryWithPaddingSize is not supported
// +thaterror:error=entry with padding size is not supported
type EntryWithPaddingSize struct {
}

// MmapSlice represents all possible error in MmapSlice function
// +thaterror:transparent
// +thaterror:wrap=*SyscallError
// +thaterror:wrap=*ProcessVMError
type MmapSliceError struct {
	Err error
}

// ELFError represents the error returned by the "elf" library
// +thaterror:transparent
type ELFError struct {
	Err error
}

// FailToFindSymbol means failed to find the symbol
// +thaterror:error=cannot find symbol
type FailToFindSymbol struct{}
