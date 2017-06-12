package tmplfn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

// assembleString like Tmpl.Assemble but using a string as a source for the template
// this is used for testing Tmpl functions
func assembleString(tmplFuncs template.FuncMap, src string) (*template.Template, error) {
	return template.New("master").Funcs(tmplFuncs).Parse(src)
}

func TestCodeBlock(t *testing.T) {
	data := map[string]interface{}{
		"data": `
echo "Hello World!"
`,
	}
	tSrc := `
This is a codeblock below

{{codeblock .data 0 0 "shell"}}
`

	expected := fmt.Sprintf(`
This is a codeblock below

%sshell
    echo "Hello World!"
%s
`, "```", "```")

	fMap := Join(Time, Page)
	tmpl, err := assembleString(fMap, tSrc)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result := fmt.Sprintf("%s", buf)
	if strings.Compare(expected, result) != 0 {
		t.Errorf("codeblock expected:\n\n%q\n\ngot:\n\n%q\n", expected, result)
		t.FailNow()
	}

	data["data"] = `
# This is a comment.
if [[ i > 1 ]]; then
    echo "i is $i"
fi

# done!
`
	expected = fmt.Sprintf(`
This is a codeblock below

%sshell
    # This is a comment.
    if [[ i > 1 ]]; then
        echo "i is $i"
    fi

    # done!
%s
`, "```", "```")

	buf = bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	result = fmt.Sprintf("%s", buf)
	if len(result) != len(expected) {
		t.Errorf("codeblock expected len: %d, got %d\n", len(expected), len(result))
	}
	if strings.Compare(expected, result) != 0 {
		t.Errorf("codeblock expected:\n\n%q\n\ngot:\n\n%q\n", expected, result)
		t.FailNow()
	}

}

func TestJoin(t *testing.T) {
	m1 := template.FuncMap{
		"helloworld": func() string {
			return "Hello World!"
		},
	}

	m2 := Join(m1, Time, Page)
	for _, key := range []string{"year", "helloworld", "nl2p"} {
		if _, OK := m2[key]; OK != true {
			t.Errorf("Can't find %s in m2", key)
		}
	}
}

func TestRender(t *testing.T) {
	src := []byte(`{"id":7877,"uri":"/repositories/2/accessions/7877","external_ids":[{"external_id":"3381","source":"Excel File :: ACCESSION","lock_version":0,"jsonmodel_type":"external_id","created_by":"admin","last_modified_by":"admin","user_mtime":"2015-10-19T23:02:01Z","system_mtime":"2015-10-19T23:02:01Z","create_time":"2015-10-19T23:02:01Z"}],"title":"Voyage médical en Italie, fait en l'année 1820","display_string":"Voyage médical en Italie, fait en l'année 1820","id_0":"1992","id_1":"00134","content_description":"Full title:  Voyage médical en Italie, fait en l'année 1820, précédé d'une excursion au volcan du Mont-Vésuve, et aux ruines d'Herculanum et de Pompeia.  2 p.l., 166 pp. 8vo, mid-19th century purple sheep-backed marbled boards (some foxing), spine gilt. First edition.  The physician Valentin (1758-1820)","condition_description":"ORIGINAL CONDITION: Very Good; PHYSICAL CONDITION: Treated; DATE: 05-Aug-1992; ACTION: Inspected; BY: Shelley Erwin","disposition":"","inventory":"","provenance":"ACQUIRED HOW OR DONOR: Purchased; ACQUIRED WHERE: Jonathan Hill Bookseller; ACQUISITION COST OR VALUE: 450.0","related_accessions":[],"accession_date":"1992-07-17","publish":true,"classifications":[],"subjects":[{"ref":"/subjects/36"}],"linked_events":[],"extents":[{"portion":"whole","number":"1","extent_type":"Multimedia","container_summary":"","physical_details":"MEDIUM: Sheep-Backed Marble Boards; FORMAT: 8vo","dimensions":"","lock_version":0,"jsonmodel_type":"extent","created_by":"admin","last_modified_by":"admin","user_mtime":"2015-10-19T23:02:01Z","system_mtime":"2015-11-25T18:05:34Z","create_time":"2015-10-19T23:02:01Z"}],"dates":[],"external_documents":[],"rights_statements":[],"user_defined":{"text_2":"17-Jul-1992: Original accession\n","text_3":"Date Record Created: 31-Jul-1992","text_4":"Storage Location: R517 .V343 1822","lock_version":0,"jsonmodel_type":"user_defined","created_by":"admin","last_modified_by":"admin","user_mtime":"2015-10-19T23:02:01Z","system_mtime":"2015-10-19T23:02:01Z","create_time":"2015-10-19T23:02:01Z","repository":{"ref":"/repositories/2"}},"suppressed":false,"resource_type":"","restrictions_apply":false,"general_note":"WHERE CREATED: Nancy, France","access_restrictions":false,"access_restrictions_note":"","use_restrictions":false,"use_restrictions_note":"","linked_agents":[{"ref":"/agents/people/2711","role":"creator","terms":[]}],"instances":[],"lock_version":0,"jsonmodel_type":"accession","created_by":"admin","last_modified_by":"admin","user_mtime":"2015-10-19T23:02:01Z","system_mtime":"2016-02-25T23:57:09Z","create_time":"2015-10-19T23:02:01Z","repository":{"ref":"/repositories/2"}}`)

	data := map[string]interface{}{}
	err := json.Unmarshal(src, &data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	tSrc := `Title: {{ .title }}`

	fMap := AllFuncs()

	tmpl, err := assembleString(fMap, tSrc)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	s := fmt.Sprintf("%s", buf)
	expected := `Title: Voyage médical en Italie, fait en l'année 1820`
	if s != expected {
		t.Errorf("expected %q, got %q", expected, s)
		t.FailNow()
	}
}

func TestMathIntFunc(t *testing.T) {
	tMap := map[string]interface{}{
		"1": 1,
		"2": 2.3,
		"3": json.Number("3"),
		"4": json.Number("4.0"),
		"5": json.Number("5.3"),
	}
	for k, v := range tMap {
		expected, _ := strconv.Atoi(k)
		fn := Math["int"].(func(interface{}) int)
		result := fn(v)
		if expected != result {
			t.Errorf("expected %d, got %T %v", expected, result, result)
		}
	}
}

func TestTempleExec(t *testing.T) {
	var (
		tpl *template.Template
		err error
	)

	tName := "stdin"
	tSrc := []byte("Hello {{ .Name -}}!")
	tmpl := New(AllFuncs())
	if err := tmpl.Add(tName, tSrc); err != nil {
		t.Errorf("%s", err)
	}
	if tpl, err = tmpl.Assemble(); err != nil {
		t.Errorf("%s", err)
	}

	var data interface{}
	json.Unmarshal([]byte(`{"Name":"Robert"}`), &data)

	if err := tpl.Execute(os.Stdout, data); err != nil {
		t.Errorf("%s", err)
	}

	tName = "hello.tmpl"
	tSrc = []byte(`
	Hello {{ .Name -}},

	Counting...

	{{range $i := (ints 1 10 2)}}
	Cnt: {{$i}},
	{{end}}
`)

	tmpl = New(AllFuncs())
	if err := tmpl.Add(tName, tSrc); err != nil {
		t.Errorf("%s", err)
	}
	if tpl, err = tmpl.Assemble(); err != nil {
		t.Errorf("%s", err)
	}
	if err := tpl.Execute(os.Stdout, data); err != nil {
		t.Errorf("%s", err)
	}
}

func TestURLEncodeDecode(t *testing.T) {
	if fn, ok := Page["urlencode"]; ok == true {
		urlencode := fn.(func(string) string)
		input := `-name:"Jack" -name:"Flanders"`
		expected := "-name%3A%22Jack%22+-name%3A%22Flanders%22"
		output := urlencode(input)
		if expected != output {
			t.Errorf("expected %q, got %s", expected, output)
		}
	} else {
		t.Errorf("Can't get function urlencode from Page map")
	}

	if fn, ok := Page["urldecode"]; ok == true {
		urldecode := fn.(func(string) string)
		input := "-name%3A%22Jack%22+-name%3A%22Flanders%22"
		expected := `-name:"Jack" -name:"Flanders"`
		output := urldecode(input)
		if expected != output {
			t.Errorf("expected %q, got %s", expected, output)
		}
	} else {
		t.Errorf("Can't get function urldecode from Page map")
	}

}

func TestPath(t *testing.T) {
	if fn, ok := Path["basename"]; ok == true {
		basename := fn.(func(...string) string)
		input := "/one/two/three.bleve"
		expected := "three.bleve"
		output := basename(input)
		if expected != output {
			t.Errorf("expected %q, got %s", expected, output)
		}
		expected = "three"
		output = basename(input, ".bleve")
		if expected != output {
			t.Errorf("expected %q, got %s", expected, output)
		}

	} else {
		t.Errorf("Can't get function basename from Path map")
	}
}

func TestTypeOf(t *testing.T) {
	if fn, ok := Math["typeof"]; ok == true {
		typeof := fn.(func(interface{}) string)
		input1 := "Hello World"
		expected := "string"
		result := typeof(input1)
		if expected != result {
			t.Errorf("Expected %s, got %s", expected, result)
		}
		input2 := []string{
			"one",
			"two",
			"three",
		}
		expected = "[]string"
		result = typeof(input2)
		if expected != result {
			t.Errorf("Expected %s, got %s", expected, result)
		}
		var input3 interface{}
		input3 = input2
		expected = "[]string"
		result = typeof(input3)
		if expected != result {
			t.Errorf("Expected %s, got %s", expected, result)
		}

		input3 = nil
		expected = "<nil>"
		result = typeof(input3)
		if expected != result {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	} else {
		t.Errorf("Can't get function typeof from Math map")
	}
}

func TestCols2Rows(t *testing.T) {
	if fn, ok := Iterables["cols2rows"]; ok == true {
		names_family := []interface{}{
			"Doiel",
			"Morrel",
			"Keswick",
		}
		names_given := []interface{}{
			"Robert",
			"Tom",
			"Tommy",
		}
		cols2rows := fn.(func(...[]interface{}) [][]interface{})
		tbl := cols2rows([]interface{}(names_family), []interface{}(names_given))
		for i := 0; i < len(tbl); i++ {
			if len(tbl[i]) != 2 {
				t.Errorf("expected 2 rows, got %d", len(tbl))
			} else {
				if tbl[i][0] != names_family[i] {
					t.Errorf("expected %d, got %d", names_family[i], tbl[i][0])
				}
				if tbl[i][1] != names_given[i] {
					t.Errorf("expected %d, got %d", names_given[i], tbl[i][1])
				}
			}
		}
	} else {
		t.Errorf("Can't get function cols2rows from Iterable map")
	}

}
