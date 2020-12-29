package fusedev

import (
	"bytes"
	error1 "github.com/YangKeao/thaterror/error"
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
func errorWrap(err error) *Error {
	switch err.(type) {
	case *IOError:
		return &Error{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *Error) Unwrap() error {
	switch err.Err.(type) {
	case *IOError:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	IOErrorErrorTmpl = template.Must(template.New("IOErrorErrorTmpl").Parse("io failed: Error: {{.Err.Error()}}"))
)

func (err *IOError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := IOErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func iOErrorWrap(err error) *IOError {
	switch err.(type) {
	case error1.Anyhow:
		return &IOError{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *IOError) Unwrap() error {
	switch err.Err.(type) {
	case error1.Anyhow:
		return err.Err
	default:
		panic("unexpected error type")
	}
}
