package watcher

import (
	"bytes"
	"text/template"
)

var (
	ErrorErrorTmpl = template.Must(template.New("ErrorErrorTmpl").Parse("fail to validate the parameter configuration"))
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
	PkgwebhookconfigwatcherError()
	error
}

func (err *EmptyTemplateLabels) PkgwebhookconfigwatcherError()  {}
func (err *EmptyConfigMapLabels) PkgwebhookconfigwatcherError() {}
func (err *EmptyTargetNamespace) PkgwebhookconfigwatcherError() {}
func ErrorWrap(err ErrorWrapUnion) *Error {
	return &Error{Err: err}
}
func (err *Error) Unwrap() error {
	return err.Err
}

var (
	EmptyTemplateLabelsErrorTmpl = template.Must(template.New("EmptyTemplateLabelsErrorTmpl").Parse("envconfig:\"TEMPLATE_LABELS\" template labels must be set"))
)

func (err *EmptyTemplateLabels) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EmptyTemplateLabelsErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	EmptyConfigMapLabelsErrorTmpl = template.Must(template.New("EmptyConfigMapLabelsErrorTmpl").Parse("envconfig:\"CONFIGMAP_LABELS\" template labels must be set"))
)

func (err *EmptyConfigMapLabels) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EmptyConfigMapLabelsErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	EmptyTargetNamespaceErrorTmpl = template.Must(template.New("EmptyTargetNamespaceErrorTmpl").Parse("envconfig:\"TARGET_NAMESPACE\" conf labels must be set while CLUSTER_SCOPED is false"))
)

func (err *EmptyTargetNamespace) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EmptyTargetNamespaceErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	FailToRenderTemplateErrorTmpl = template.Must(template.New("FailToRenderTemplateErrorTmpl").Parse(""))
)

func (err *FailToRenderTemplate) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := FailToRenderTemplateErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
