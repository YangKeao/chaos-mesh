package ptrace

import (
	"bytes"
	"text/template"
)

func (err *TraceError) Error() string {
	return err.Error()
}

type TraceErrorWrapUnion interface {
	PkgptraceTraceError()
	error
}

func (err *PtraceError) PkgptraceTraceError()  {}
func (err *WaitPidError) PkgptraceTraceError() {}
func TraceErrorWrap(err TraceErrorWrapUnion) *TraceError {
	return &TraceError{Err: err}
}
func (err *TraceError) Unwrap() error {
	return err.Err
}

var (
	PtraceErrorErrorTmpl = template.Must(template.New("PtraceErrorErrorTmpl").Parse("fail to ptrace({{.Operation}}) on process {{.Tid}}"))
)

func (err *PtraceError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := PtraceErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	EntryWithPaddingSizeErrorTmpl = template.Must(template.New("EntryWithPaddingSizeErrorTmpl").Parse("entry with padding size is not supported"))
)

func (err *EntryWithPaddingSize) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EntryWithPaddingSizeErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *MmapSliceError) Error() string {
	return err.Error()
}

type MmapSliceErrorWrapUnion interface {
	PkgptraceMmapSliceError()
	error
}

func (err *SyscallError) PkgptraceMmapSliceError()   {}
func (err *ProcessVMError) PkgptraceMmapSliceError() {}
func MmapSliceErrorWrap(err MmapSliceErrorWrapUnion) *MmapSliceError {
	return &MmapSliceError{Err: err}
}
func (err *MmapSliceError) Unwrap() error {
	return err.Err
}
func (err *ELFError) Error() string {
	return err.Error()
}
func (err *ELFOperationError) Error() string {
	return err.Error()
}

type ELFOperationErrorWrapUnion interface {
	PkgptraceELFOperationError()
	error
}

func (err *ProcessVMError) PkgptraceELFOperationError()       {}
func (err *EntryWithPaddingSize) PkgptraceELFOperationError() {}
func (err ELFError) PkgptraceELFOperationError()              {}
func (err *FailToFindSymbol) PkgptraceELFOperationError()     {}
func ELFOperationErrorWrap(err ELFOperationErrorWrapUnion) *ELFOperationError {
	return &ELFOperationError{Err: err}
}
func (err *ELFOperationError) Unwrap() error {
	return err.Err
}

var (
	FailToFindSymbolErrorTmpl = template.Must(template.New("FailToFindSymbolErrorTmpl").Parse("cannot find symbol"))
)

func (err *FailToFindSymbol) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := FailToFindSymbolErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *StepError) Error() string {
	return err.Error()
}

type StepErrorWrapUnion interface {
	PkgptraceStepError()
	error
}

func (err *WaitPidError) PkgptraceStepError() {}
func (err *PtraceError) PkgptraceStepError()  {}
func StepErrorWrap(err StepErrorWrapUnion) *StepError {
	return &StepError{Err: err}
}
func (err *StepError) Unwrap() error {
	return err.Err
}
func (err *SyscallError) Error() string {
	return err.Error()
}

type SyscallErrorWrapUnion interface {
	PkgptraceSyscallError()
	error
}

func (err *PtraceError) PkgptraceSyscallError()      {}
func (err *StepError) PkgptraceSyscallError()        {}
func (err *TooManyArguments) PkgptraceSyscallError() {}
func SyscallErrorWrap(err SyscallErrorWrapUnion) *SyscallError {
	return &SyscallError{Err: err}
}
func (err *SyscallError) Unwrap() error {
	return err.Err
}

var (
	TooManyArgumentsErrorTmpl = template.Must(template.New("TooManyArgumentsErrorTmpl").Parse("too many arguments for a syscall: {{.Arguments}}"))
)

func (err *TooManyArguments) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := TooManyArgumentsErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	ProcessVMErrorErrorTmpl = template.Must(template.New("ProcessVMErrorErrorTmpl").Parse("process_vm_readv returned an error: {{.Errno}}"))
)

func (err *ProcessVMError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ProcessVMErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	WaitPidErrorErrorTmpl = template.Must(template.New("WaitPidErrorErrorTmpl").Parse("waitpid failed"))
)

func (err *WaitPidError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := WaitPidErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
