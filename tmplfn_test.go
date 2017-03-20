package tmplfn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"text/template"
)

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
	src := []byte(`{
"id": "7877",
"uri": "/repositories/2/accessions/7877",
"title": "Voyage médical en Italie, fait en l'année 1820",
"identifier": "1992-00134",
"resource_type": "",
"content_description": "Full title: Voyage médical en Italie, fait en l'année 1820, précédé d'une excursion au volcan du Mont-Vésuve, et aux ruines d'Herculanum et de Pompeia. 2 p.l., 166 pp. 8vo, mid-19th century purple sheep-backed marbled boards (some foxing), spine gilt. First edition. The physician Valentin (1758-1820)",
"condition_description": "ORIGINAL CONDITION: Very Good; PHYSICAL CONDITION: Treated; DATE: 05-Aug-1992; ACTION: Inspected; BY: Shelley Erwin",
"access_restrictions": false,
"access_restrictions_notes": "",
"use_restrictions": false,
"use_restrictions_notes": "",
"dates": [ ],
"date_expression": "",
"subjects": [
"Rare Book Collection"
],
"subjects_function": [
"Rare Book Collection"
],
"extents": [
"MEDIUM: Sheep-Backed Marble Boards; FORMAT: 8vo"
],
"linked_agents_creators": [
"Valentin, Louis"
],
"linked_agents_subjects": null,
"accession_date": "1992-07-17",
"created_by": "admin",
"created": "2015-10-19T23:02:01Z",
"last_modified_by": "admin",
"last_modified": "2015-10-19T23:02:01Z"
}`)

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
