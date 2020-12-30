package flags

import (
	"bytes"
	"text/template"
)

var (
	BadFormatErrorTmpl = template.Must(template.New("BadFormatErrorTmpl").Parse("the format of {{.Value}} is incorrect! should be key=value[,key2=value2]"))
)

func (err *BadFormat) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := BadFormatErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
