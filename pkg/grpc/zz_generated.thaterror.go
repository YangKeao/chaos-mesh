package grpc

import (
	"bytes"
	"text/template"
)

func (err *GrpcError) Error() string {
	return err.Error()
}
func (err *CreateConnectionError) Error() string {
	return err.Error()
}

type CreateConnectionErrorWrapUnion interface {
	PkggrpcCreateConnectionError()
	error
}

func (err *FailToFindEndpointIP) PkggrpcCreateConnectionError() {}
func (err *GrpcError) PkggrpcCreateConnectionError()            {}
func CreateConnectionErrorWrap(err CreateConnectionErrorWrapUnion) *CreateConnectionError {
	return &CreateConnectionError{Err: err}
}
func (err *CreateConnectionError) Unwrap() CreateConnectionErrorWrapUnion {
	return err.Err.(CreateConnectionErrorWrapUnion)
}

var (
	FailToFindEndpointIPErrorTmpl = template.Must(template.New("FailToFindEndpointIPErrorTmpl").Parse("failed to find node {{.NodeName}} in {{.Endpoints}}"))
)

func (err *FailToFindEndpointIP) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := FailToFindEndpointIPErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
