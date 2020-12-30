package config

import (
	"bytes"
	"text/template"
)

var (
	MissingTemplateNameErrorTmpl = template.Must(template.New("MissingTemplateNameErrorTmpl").Parse("template field is required for template args config"))
)

func (err *MissingTemplateName) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := MissingTemplateNameErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	MissingNameErrorTmpl = template.Must(template.New("MissingNameErrorTmpl").Parse("name field is required for template args config"))
)

func (err *MissingName) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := MissingNameErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	ErrorErrorTmpl = template.Must(template.New("ErrorErrorTmpl").Parse(""))
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
	PkgwebhookconfigError()
	error
}

func (err *NotFound) PkgwebhookconfigError()            {}
func (err *MissingTemplateName) PkgwebhookconfigError() {}
func (err *MissingName) PkgwebhookconfigError()         {}
func ErrorWrap(err ErrorWrapUnion) *Error {
	return &Error{Err: err}
}
func (err *Error) Unwrap() error {
	return err.Err
}

var (
	NotFoundErrorTmpl = template.Must(template.New("NotFoundErrorTmpl").Parse("injection config not found  {{if not .Key}}for key {{.Key}} {{end}} at namespace {{.Namespace}}"))
)

func (err *NotFound) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := NotFoundErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
