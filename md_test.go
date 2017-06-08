package tmplfn

import (
	"testing"
)

func TestMarkdownFn(t *testing.T) {
	var md func(string) string

	if fn, ok := Markdown["markdown"]; ok == true {
		md = fn.(func(string) string)
	} else {
		t.Errorf("Can't retrieve markdown function from Markdown map")
		t.FailNow()
	}

	src := `
markdown
: a simple text formatting language easily rendered as HTML and LaTeX
`

	expected := `<dl>
<dt>markdown</dt>
<dd>a simple text formatting language easily rendered as HTML and LaTeX</dd>
</dl>
`

	result := md(src)
	if result != expected {
		t.Errorf("expected\n---------------\n%s\n---------------\ngot\n---------------\n%s\n---------------\n", expected, result)
	}
}
