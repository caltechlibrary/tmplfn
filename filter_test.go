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
	filter_src := `(and .one .two)`
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
		filter_src = src
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

func TestNumericComparisons(t *testing.T) {
src := `[
{"Username":"one@example.org","cnt2014":1,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/05/15","yr2015":"","yr2016":"","yr2017":""},
{"Username":"two@example.org","cnt2014":0,"cnt2015":1,"cnt2016":1,"cnt2017":1,"yr2014":"","yr2015":"09-10-15","yr2016":"2016/01/07","yr2017":"2017/05/01"},
{"Username":"three@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"05-06-15","yr2016":"","yr2017":""},
{"Username":"four@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"07-20-15","yr2016":"","yr2017":""},
{"Username":"five@example.org","cnt2014":3,"cnt2015":3,"cnt2016":5,"cnt2017":0,"yr2014":"2014/08/23 2014/08/24 2014/08/24","yr2015":"03-16-15 03-18-15 05-28-15","yr2016":"2016/07/20 2016/11/18 2016/11/19 2016/11/19 2016/11/20","yr2017":""},
{"Username":"fix@example.org","cnt2014":0,"cnt2015":2,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"08-29-15 09-23-15","yr2016":"","yr2017":""},
{"Username":"six@example.org","cnt2014":0,"cnt2015":1,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"06-09-15","yr2016":"2016/11/01","yr2017":""},
{"Username":"known@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"02-17-15","yr2016":"","yr2017":""},
{"Username":"habbit@example.org","cnt2014":1,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/02/05","yr2015":"","yr2016":"","yr2017":""},
{"Username":"zip@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"08-25-15","yr2016":"","yr2017":""},
{"Username":"bang@example.org","cnt2014":0,"cnt2015":0,"cnt2016":3,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/07/21 2016/07/21 2016/09/21","yr2017":""},
{"Username":"ping@example.org","cnt2014":0,"cnt2015":1,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"08-27-15","yr2016":"2016/12/07","yr2017":""},
{"Username":"slew@example.org","cnt2014":0,"cnt2015":0,"cnt2016":0,"cnt2017":1,"yr2014":"","yr2015":"","yr2016":"","yr2017":"2017/02/09"},
{"Username":"downbeat@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/05/13","yr2017":""},
{"Username":"eriff@example.org","cnt2014":0,"cnt2015":0,"cnt2016":0,"cnt2017":2,"yr2014":"","yr2015":"","yr2016":"","yr2017":"2017/02/13 2017/06/17"},
{"Username":"uriff@example.org","cnt2014":2,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"2014/04/21 2014/04/22","yr2015":"02-05-15","yr2016":"","yr2017":""},
{"Username":"triif@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"11-16-15","yr2016":"","yr2017":""},
{"Username":"fiferroon@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":1,"yr2014":"","yr2015":"02-22-15","yr2016":"","yr2017":"2017/02/05"},
{"Username":"plingdo@example.org","cnt2014":0,"cnt2015":0,"cnt2016":0,"cnt2017":1,"yr2014":"","yr2015":"","yr2016":"","yr2017":"2017/01/07"},
{"Username":"quizzip@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/09/24","yr2017":""},
{"Username":"oscar@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":3,"yr2014":"","yr2015":"","yr2016":"2016/07/08","yr2017":"2017/01/10 2017/01/10 2017/03/18"},
{"Username":"vellum@example.org","cnt2014":2,"cnt2015":0,"cnt2016":0,"cnt2017":2,"yr2014":"2014/07/29 2014/08/05","yr2015":"","yr2016":"","yr2017":"2017/05/28 2017/05/29"},
{"Username":"trainsspot@example.org","cnt2014":2,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"2014/11/24 2014/11/24","yr2015":"","yr2016":"2016/02/14","yr2017":""},
{"Username":"ipswitch@example.org","cnt2014":2,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/06/04 2014/05/02","yr2015":"","yr2016":"","yr2017":""},
{"Username":"ngapping@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"10-14-15","yr2016":"","yr2017":""},
{"Username":"rivieria@example.org","cnt2014":1,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/02/27","yr2015":"","yr2016":"","yr2017":""},
{"Username":"minimoo@example.org","cnt2014":0,"cnt2015":2,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"12-01-15 12-02-15","yr2016":"","yr2017":""},
{"Username":"moominis@example.org","cnt2014":2,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/05/07 2014/07/09","yr2015":"","yr2016":"","yr2017":""},
{"Username":"lalalala@example.org","cnt2014":0,"cnt2015":0,"cnt2016":0,"cnt2017":1,"yr2014":"","yr2015":"","yr2016":"","yr2017":"2017/03/31"},
{"Username":"blertbert@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":2,"yr2014":"","yr2015":"","yr2016":"2016/11/30","yr2017":"2017/02/18 2017/02/18"},
{"Username":"alskaemin@example.org","cnt2014":0,"cnt2015":1,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"06-28-15","yr2016":"2016/09/08","yr2017":""},
{"Username":"poikljn@example.org","cnt2014":1,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/05/06","yr2015":"","yr2016":"","yr2017":""},
{"Username":"tyrcrerghc@example.org","cnt2014":1,"cnt2015":0,"cnt2016":0,"cnt2017":0,"yr2014":"2014/05/21","yr2015":"","yr2016":"","yr2017":""},
{"Username":"poipoi@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/02/15","yr2017":""},
{"Username":"dernatue@example.org","cnt2014":0,"cnt2015":0,"cnt2016":2,"cnt2017":3,"yr2014":"","yr2015":"","yr2016":"2016/01/04 2016/04/03","yr2017":"2017/03/12 2017/04/05 2017/02/22"},
{"Username":"swineytodd@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/08/31","yr2017":""},
{"Username":"mymy@example.org","cnt2014":0,"cnt2015":0,"cnt2016":2,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/09/11 2016/09/11","yr2017":""},
{"Username":"youyou@example.org","cnt2014":0,"cnt2015":0,"cnt2016":1,"cnt2017":0,"yr2014":"","yr2015":"","yr2016":"2016/04/04","yr2017":""},
{"Username":"usus@example.org","cnt2014":3,"cnt2015":0,"cnt2016":5,"cnt2017":0,"yr2014":"2014/09/23 2014/10/01 2014/10/01","yr2015":"","yr2016":"2016/02/01 2016/03/03 2016/03/03 2016/04/14 2016/04/14","yr2017":""},
{"Username":"splat@example.org","cnt2014":0,"cnt2015":1,"cnt2016":0,"cnt2017":0,"yr2014":"","yr2015":"12-05-15","yr2016":"","yr2017":""}
]`

	tData := []map[string]interface{}{}
	if err := json.Unmarshal([]byte(src), &tData); err != nil {
		t.Error(err)
		t.FailNow()
	}
	//FIXME: It is ugly that I have to explicitly convert a JSON.Number to Go type
	f, err := ParseFilter(`(gt (int .cnt2014) 2)`)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// should find two records with .cnt2014 greater than 2
	expectedCnt := 2
	cnt := 0
	for _, rec := range tData {
		if ok, err := f.Apply(rec); err != nil {
			t.Errorf("Can't apply filter to %+v, %s", rec, err)
		} else if ok == true {
			cnt++
		}
	}
	if cnt != expectedCnt {
		t.Errorf("expected %d, got %d", expectedCnt, cnt)
	}
}
