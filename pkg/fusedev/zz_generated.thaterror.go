package fusedev

import (
	"bytes"
	"text/template"
)

var (
	ErrorErrorTmpl = template.Must(template.New("ErrorErrorTmpl").Parse("fail to operate the fuse device: Error: {{.Err.Error()}}"))
)

func (err *Error) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

type ErrorWrapUnion interface {
	PkgfusedevError()
	error
}

func (err *InvalidCgroupEntry) PkgfusedevError()   {}
func (err *DeviceCgroupNotFound) PkgfusedevError() {}
func ErrorWrap(err ErrorWrapUnion) *Error {
	return &Error{Err: err}
}
func (err *Error) Unwrap() ErrorWrapUnion {
	return err.Err.(ErrorWrapUnion)
}

var (
	InvalidCgroupEntryErrorTmpl = template.Must(template.New("InvalidCgroupEntryErrorTmpl").Parse("invalid cgroup entry: {{.Text}}"))
)

func (err *InvalidCgroupEntry) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := InvalidCgroupEntryErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	DeviceCgroupNotFoundErrorTmpl = template.Must(template.New("DeviceCgroupNotFoundErrorTmpl").Parse("fail to find device cgroup"))
)

func (err *DeviceCgroupNotFound) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := DeviceCgroupNotFoundErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
