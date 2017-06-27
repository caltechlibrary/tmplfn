package tmplfn

import (
	"bytes"
	"fmt"
	"text/template"
)

// Filter holds the parsed filter source that can be applied to a JSON record to determine if the record is included or exluded
// NOTE: We're riding on the back of a text template to render true or false
type Filter struct {
	tmpl *template.Template
}

// ParseFilter parses a byte slice and returns a Filter struct and error
func ParseFilter(src string) (*Filter, error) {
	tmpl, err := template.New("filter").Funcs(AllFuncs()).Parse(fmt.Sprintf("{{- if %s -}}true{{- else -}}false{{- end -}}", src))
	if err != nil {
		return nil, err
	}
	f := new(Filter)
	f.tmpl = tmpl
	return f, nil
}

// Apply uses Filter Struct and takes an interface{} object returns true if filter matched or false
func (f *Filter) Apply(data interface{}) (bool, error) {
	buf := bytes.NewBuffer([]byte{})
	err := f.tmpl.Execute(buf, data)
	if err != nil {
		return false, err
	}
	if buf.String() == "true" {
		return true, nil
	}
	return false, nil
}
