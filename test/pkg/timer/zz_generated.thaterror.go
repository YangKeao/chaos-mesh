package timer

import (
	"bytes"
	"text/template"
)

func (err *ProcessStartError) Error() string {
	return err.Error()
}
func (err *ProcessStartError) Unwrap() error {
	return nil
}

var (
	ErrorErrorTmpl = template.Must(template.New("ErrorErrorTmpl").Parse("failed to get time from timer output: Err: {{.Err}}"))
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
	case *ParseIntFailed, *IOError, *ProcessStartError:
		return &Error{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *Error) Unwrap() error {
	switch err.Err.(type) {
	case *ParseIntFailed, *IOError, *ProcessStartError:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	ParseIntFailedErrorTmpl = template.Must(template.New("ParseIntFailedErrorTmpl").Parse("failed to parse int: S: {{.S}}, Base: {{.Base}}, Bitsize: {{.Bitsize}}"))
)

func (err *ParseIntFailed) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ParseIntFailedErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *ParseIntFailed) Unwrap() error {
	return nil
}
func (err *IOError) Error() string {
	return err.Error()
}
func (err *IOError) Unwrap() error {
	return nil
}
