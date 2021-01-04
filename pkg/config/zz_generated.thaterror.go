package config

import (
	"bytes"
	"text/template"
)

func (err *EnvConfigError) Error() string {
	return err.Error()
}

var (
	ParsePersistTTLConfigErrorErrorTmpl = template.Must(template.New("ParsePersistTTLConfigErrorErrorTmpl").Parse(""))
)

func (err *ParsePersistTTLConfigError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ParsePersistTTLConfigErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

type ParsePersistTTLConfigErrorWrapUnion interface {
	PkgconfigParsePersistTTLConfigError()
	error
}

func (err EnvConfigError) PkgconfigParsePersistTTLConfigError() {}
func ParsePersistTTLConfigErrorWrap(err ParsePersistTTLConfigErrorWrapUnion) *ParsePersistTTLConfigError {
	return &ParsePersistTTLConfigError{Err: err}
}
func (err *ParsePersistTTLConfigError) Unwrap() ParsePersistTTLConfigErrorWrapUnion {
	return err.Err.(ParsePersistTTLConfigErrorWrapUnion)
}
