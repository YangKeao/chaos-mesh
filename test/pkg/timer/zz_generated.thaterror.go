package timer

import (
	"bytes"
	"text/template"
)

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

type ErrorWrapUnion interface {
	TestpkgtimerError()
	error
}

func (err *ProcessStartError) TestpkgtimerError() {}
func ErrorWrap(err ErrorWrapUnion) *Error {
	return &Error{Err: err}
}
func (err *Error) Unwrap() error {
	return err.Err
}
func (err *ProcessStartError) Error() string {
	return err.Error()
}
