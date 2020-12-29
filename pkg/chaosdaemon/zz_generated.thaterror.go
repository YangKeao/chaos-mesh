package chaosdaemon

import (
	"bytes"
	error1 "github.com/YangKeao/thaterror/error"
	"text/template"
)

func (err *IPSetError) Error() string {
	return err.Error()
}
func iPSetErrorWrap(err error) *IPSetError {
	switch err.(type) {
	case *CommandExecuteError:
		return &IPSetError{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *IPSetError) Unwrap() error {
	switch err.Err.(type) {
	case *CommandExecuteError:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	CommandExecuteErrorErrorTmpl = template.Must(template.New("CommandExecuteErrorErrorTmpl").Parse("Command {{.Command}} failed with return value {{.Ret}} and output:\n{{.Output}}\n+thaterror:error:end\n"))
)

func (err *CommandExecuteError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := CommandExecuteErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *CommandExecuteError) Unwrap() error {
	return nil
}
func (err *ContainerClientError) Error() string {
	return err.Error()
}
func containerClientErrorWrap(err error) *ContainerClientError {
	switch err.(type) {
	case *BadContainerState, *InvalidPidError, error1.Anyhow:
		return &ContainerClientError{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *ContainerClientError) Unwrap() error {
	switch err.Err.(type) {
	case *BadContainerState, *InvalidPidError, error1.Anyhow:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	InvalidPidErrorErrorTmpl = template.Must(template.New("InvalidPidErrorErrorTmpl").Parse("invalid pid for container: {{.ContainerID}}"))
)

func (err *InvalidPidError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := InvalidPidErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *InvalidPidError) Unwrap() error {
	return nil
}
func (err *IptablesError) Error() string {
	return err.Error()
}
func iptablesErrorWrap(err error) *IptablesError {
	switch err.(type) {
	case *CommandExecuteError:
		return &IptablesError{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *IptablesError) Unwrap() error {
	switch err.Err.(type) {
	case *CommandExecuteError:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	UnknownChainDirectionErrorTmpl = template.Must(template.New("UnknownChainDirectionErrorTmpl").Parse("unknown chain direction {{.Direction}}"))
)

func (err *UnknownChainDirection) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := UnknownChainDirectionErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *UnknownChainDirection) Unwrap() error {
	return nil
}

var (
	DNSServerShouldNotBeEmptyErrorTmpl = template.Must(template.New("DNSServerShouldNotBeEmptyErrorTmpl").Parse("\"DnsServer field shouldn't be empty\""))
)

func (err *DNSServerShouldNotBeEmpty) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := DNSServerShouldNotBeEmptyErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *DNSServerShouldNotBeEmpty) Unwrap() error {
	return nil
}

var (
	BadRequestErrorTmpl = template.Must(template.New("BadRequestErrorTmpl").Parse("grpc request invalid: {{.Err.Error()}}"))
)

func (err *BadRequest) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := BadRequestErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func badRequestWrap(err error) *BadRequest {
	switch err.(type) {
	case *EmptyDNSServerError:
		return &BadRequest{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *BadRequest) Unwrap() error {
	switch err.Err.(type) {
	case *EmptyDNSServerError:
		return err.Err
	default:
		panic("unexpected error type")
	}
}

var (
	EmptyDNSServerErrorErrorTmpl = template.Must(template.New("EmptyDNSServerErrorErrorTmpl").Parse("DnsServer in the request is empty"))
)

func (err *EmptyDNSServerError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := EmptyDNSServerErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *EmptyDNSServerError) Unwrap() error {
	return nil
}

var (
	BadContainerStateErrorTmpl = template.Must(template.New("BadContainerStateErrorTmpl").Parse("the status of container {{.ContainerID}} is {{.State}}"))
)

func (err *BadContainerState) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := BadContainerStateErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func (err *BadContainerState) Unwrap() error {
	return nil
}

var (
	InvalidRequestErrorTmpl = template.Must(template.New("InvalidRequestErrorTmpl").Parse("grpc Request is invalid: Error: {{.Err.Error()}}"))
)

func (err *InvalidRequest) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := InvalidRequestErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
func invalidRequestWrap(err error) *InvalidRequest {
	switch err.(type) {
	case *DNSServerShouldNotBeEmpty, *UnknownChainDirection:
		return &InvalidRequest{Err: err}
	default:
		panic("unexpected error type")
	}
}
func (err *InvalidRequest) Unwrap() error {
	switch err.Err.(type) {
	case *DNSServerShouldNotBeEmpty, *UnknownChainDirection:
		return err.Err
	default:
		panic("unexpected error type")
	}
}
