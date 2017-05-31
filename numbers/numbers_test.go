package numbers

// this is the sketch
import (
	"encoding/json"
	"testing"
)

func TestTypeHandling(t *testing.T) {
	jsN1 := json.Number("1")
	e1 := int64(1)
	r1 := toType(jsN1, int64Type)
	if r1 != e1 {
		t.Errorf("expected %d, got %T %v", e1, r1, r1)
	}

	jsN2 := json.Number("1.5")
	e2 := float64(1.5)
	r2 := toType(jsN2, float64Type)
	if r2 != e2 {
		t.Errorf("expected %g, got %T %v", e2, r2, r2)
	}

	e3 := float64(2.5)
	r3 := Add(jsN1, jsN2)
	rType := numberType(r3)
	if rType != float64Type {
		t.Errorf("expected float64Type, got %T %v", r3, r3)
	}
	r3f := toType(r3, float64Type)
	if r3f != e3 {
		t.Errorf("expected %g, got %T %v", e3, r3, r3)
	}

	e4 := float64(0.5)
	r4 := Subtract(jsN2, jsN1)
	rType = numberType(r4)
	if rType != float64Type {
		t.Errorf("expected float64Type, got %T %v", r4, r4)
	}
	r4f := toType(r4, float64Type)
	if r4f != e4 {
		t.Errorf("expected %g, got %T %v", e4, r4, r4)
	}

	e5 := float64(-0.5)
	r5 := Subtract(jsN1, jsN2)
	rType = numberType(r5)
	if rType != float64Type {
		t.Errorf("expected float64Type, got %T %v", r5, r5)
	}
	r5f := toType(r5, float64Type)
	if r5f != e5 {
		t.Errorf("expected %g, got %T %v", e5, r5, r5)
	}

	e6 := int(6)
	r6 := Addi(int64(2), float64(4.0))
	rType = numberType(r6)
	if rType != intType {
		t.Errorf("expected intType, got %T %v", r6, r6)
	}
	r6i := toType(r6, intType)
	if r6i != e6 {
		t.Errorf("expected %d, got %T %v", e6, r6, r6)
	}
}
