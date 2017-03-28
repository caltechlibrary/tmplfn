package tmplfn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"text/template"
)

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

	fMap := Join(TimeMap, PageMap)
	tmpl, err := AssembleString(fMap, tSrc)
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
	if result != expected {
		t.Errorf("codeblock expected:\n\n%s\n\ngot:\n\n%s\n", expected, buf)
		t.FailNow()
	}

	data["data"] = `
// This is a comment.
if i > 1 {
	fmt.Printf("i is %d\n", i)
}

// done!
`
	expected = fmt.Sprintf(`
This is a codeblock below

%sgolang
    // This is a comment.
    if i > 1 {
        fmt.Printf("i is %%d\n", i)
    }

    // done!
%s
`, "```", "```")

	buf = bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, data)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

	result = fmt.Sprintf("%s", buf)
	if result != expected {
		t.Errorf("codeblock expected:\n\n%s\n\ngot:\n\n%s\n", expected, buf)
		t.FailNow()
	}

}

func TestJoin(t *testing.T) {
	m1 := template.FuncMap{
		"helloworld": func() string {
			return "Hello World!"
		},
	}

	m2 := Join(m1, TimeMap, PageMap)
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

	fMap := Join(TimeMap, PageMap)
	tmpl, err := AssembleString(fMap, tSrc)
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
