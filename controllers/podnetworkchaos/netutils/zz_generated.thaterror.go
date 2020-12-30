package netutils

import (
	"bytes"
	"text/template"
)

var (
	DNSLookupFailedErrorTmpl = template.Must(template.New("DNSLookupFailedErrorTmpl").Parse("fail to lookup domain: Name: {{.Name}} Error: {{.Err.Error()}}"))
)

func (err *DNSLookupFailed) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := DNSLookupFailedErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
