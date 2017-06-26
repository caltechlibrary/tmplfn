package tmplfn

import (
	"encoding/json"
	"testing"
)

func TestFilter(t *testing.T) {
	json_src := []byte(`{ "one": true, "two": true }`)
	rec := map[string]interface{}{}
	if err := json.Unmarshal(json_src, &rec); err != nil {
		t.Error(err)
		t.FailNow()
	}

	expected := true
	filter_src := []byte(`(and .one .two)`)
	filter, err := ParseFilter(filter_src)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	result, err := filter.Apply(rec)
	if err != nil {
		t.Errorf("Applying filter failed, %s", err)
		t.FailNow()
	}
	if result != expected {
		t.Errorf("expected %t, got %t for filter %s", expected, result, filter_src)
	}

	rec["three"] = false
	for src, expected := range map[string]bool{
		`add 1 1`:                           true,
		`(add 1 1)`:                         true,
		`false`:                             false,
		`(false)`:                           false,
		`(and .one .three)`:                 false,
		`(or .one .three)`:                  true,
		`(and .one .two (eq .three false))`: true,
	} {
		filter_src = []byte(src)
		filter, err := ParseFilter(filter_src)
		if err != nil {
			t.Errorf("Can't parse filter source: %q, %s", filter_src, err)
			t.FailNow()
		}
		result, err := filter.Apply(rec)
		if err != nil {
			t.Errorf("Applying filter failed, %s", err)
			t.FailNow()
		}
		if result != expected {
			t.Errorf("expected %t, got %t for filter %s", expected, result, filter_src)
		}
	}

}
