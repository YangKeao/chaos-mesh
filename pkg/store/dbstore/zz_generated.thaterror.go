package dbstore

import (
	"bytes"
	"text/template"
)

var (
	OpenDatabaseErrorErrorTmpl = template.Must(template.New("OpenDatabaseErrorErrorTmpl").Parse("failed to open DB({{.Driver}}:{{.Datasource}}): Err: {{.Err}}"))
)

func (err *OpenDatabaseError) Error() string {
	buf := new(bytes.Buffer)
	tmplErr := OpenDatabaseErrorErrorTmpl.Execute(buf, err)
	if tmplErr != nil {
		panic("fail to render error template")
	}
	return buf.String()
}
