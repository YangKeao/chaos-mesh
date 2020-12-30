package commonerror

import (
	"bytes"
	"text/template"
)

func (err *ParseIntError) TestpkgtimerError()                        {}
func (err *IOError) TestpkgtimerError()                              {}
func (err *IOError) PkgmapreaderError()                              {}
func (err *ParseIntError) PkgmapreaderError()                        {}
func (err *ParseDurationError) PkgconfigParsePersistTTLConfigError() {}
func (err *IOError) PkgfusedevError()                                {}
func (err *IOError) PkgwebhookconfigError()                          {}
func (err *UnmarshalError) PkgwebhookconfigError()                   {}
func (err *KubernetesError) PkggrpcCreateConnectionError()           {}
func (err *IOError) PkgptraceTraceError()                            {}
func (err *ParseIntError) PkgptraceTraceError()                      {}

var (
	ParseDurationErrorErrorTmpl = template.Must(template.New("ParseDurationErrorErrorTmpl").Parse("failed to parse duration: {{.S}}"))
)

func (err *ParseDurationError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ParseDurationErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *IOError) Error() string {
	return err.Error()
}
func (err *UnmarshalError) Error() string {
	return err.Error()
}
func (err *KubernetesError) Error() string {
	return err.Error()
}

var (
	ContainerStatusEmptyErrorTmpl = template.Must(template.New("ContainerStatusEmptyErrorTmpl").Parse("the container statuses of {{.Name}}/{{.Namespace}} is empty"))
)

func (err *ContainerStatusEmpty) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ContainerStatusEmptyErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	ParseIntErrorErrorTmpl = template.Must(template.New("ParseIntErrorErrorTmpl").Parse("failed to parse int: S: {{.S}}, Base: {{.Base}}, Bitsize: {{.Bitsize}}"))
)

func (err *ParseIntError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := ParseIntErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
