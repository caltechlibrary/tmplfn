package tmplfn

import (
	"testing"
	"text/template"
)

func TestJoin(t *testing.T) {
	m1 := template.FuncMap{
		"helloworld": func() string {
			return "Hello World!"
		},
	}

	m2 := Join(m1, Time)
	for _, key := range []string{"year", "helloworld"} {
		if _, OK := m2[key]; OK != true {
			t.Errorf("Can't find %s in m2", key)
		}
	}
}
