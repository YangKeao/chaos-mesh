package clientpool

import (
	"bytes"
	"text/template"
)

var (
	EmptyTokenErrorErrorTmpl = template.Must(template.New("EmptyTokenErrorErrorTmpl").Parse("token is empty"))
)

func (err *EmptyTokenError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EmptyTokenErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}

var (
	LRUErrorErrorTmpl = template.Must(template.New("LRUErrorErrorTmpl").Parse("fail to create lru with size: {{.Size}}. Err: {{.Err}}"))
)

func (err *LRUError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := LRUErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *CreateClientError) Error() string {
	return err.Error()
}

type CreateClientErrorWrapUnion interface {
	PkgclientpoolCreateClientError()
	error
}

func (err *KubernetesCreateClientError) PkgclientpoolCreateClientError() {}
func (err *EmptyTokenError) PkgclientpoolCreateClientError()             {}
func CreateClientErrorWrap(err CreateClientErrorWrapUnion) *CreateClientError {
	return &CreateClientError{Err: err}
}
func (err *CreateClientError) Unwrap() CreateClientErrorWrapUnion {
	return err.Err.(CreateClientErrorWrapUnion)
}

var (
	KubernetesCreateClientErrorErrorTmpl = template.Must(template.New("KubernetesCreateClientErrorErrorTmpl").Parse("fail to create kubernetes client with config {{.Config}} and options {{.Options}}"))
)

func (err *KubernetesCreateClientError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := KubernetesCreateClientErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
