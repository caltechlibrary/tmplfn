package numbers

// this is the sketch
import (
	"testing"
)

type MathOpTest struct {
	Op       func(Number, Number) Number
	Expected Number
	Params   []Number
}

func TestMath(t *testing.T) {
	//zero := NewNumber(0)
	one := NewNumber(1)
	onePointZero := NewNumber(1.0)
	//onePointFive := NewNumber(1.5)
	two := NewNumber(2)
	twoPointZero := NewNumber(2.0)
	three := NewNumber(3)
	threePointZero := NewNumber(3.0)

	testSeq := map[string]MathOpTest{
		"addInt": MathOpTest{
			Op:       Add,
			Expected: three,
			Params:   []Number{one, two},
		},
		"addFloat": MathOpTest{
			Op:       Add,
			Expected: threePointZero,
			Params:   []Number{onePointZero, twoPointZero},
		},
	}

	for op, tObj := range testSeq {
		result := tObj.Op(tObj.Params[0], tObj.Params[1])
		if tObj.Expected.Value != result.Value {
			t.Errorf("%s, expected %s, got %s", op, tObj.Expected, result)
		}
	}
}
