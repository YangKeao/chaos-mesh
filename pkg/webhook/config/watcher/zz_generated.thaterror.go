package watcher

import (
	"bytes"
	"text/template"
)

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
func (err *EmptyTargetNamespace) Unwrap() error {
	return nil
}

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
func errorWrap(err error) *Error {
	switch err.(type) {
	case *EmptyTemplateLabels, *EmptyConfigMapLabels, *EmptyTargetNamespace:
		return &Error{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *Error) Unwrap() error {
	switch err.Err.(type) {
	case *EmptyTemplateLabels, *EmptyConfigMapLabels, *EmptyTargetNamespace:
		return err.Err
	default:
		panic("unexpected error type")
	}
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
func (err *EmptyTemplateLabels) Unwrap() error {
	return nil
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
func (err *EmptyConfigMapLabels) Unwrap() error {
	return nil
}
