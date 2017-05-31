package numbers

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// tmplfn supports calculations with two types of numbers, 
// int64 and float64 if you store a 32 bit counter part 
// they will be converted to 64 bit versions before
// executing a computation (e.g. add, substract)

const (
	NaNType = iota
	IntType
	Int64Type
	Float32Type
	Float64Type
	JSONNumberType
	StringType
)

type Number struct {
	Value interface{}
	NaN   bool
}

func  NumberType(value interface{}) int {
	switch value.(type) {
	case string:
		return StringType
	case json.Number:
		return JSONNumberType
	case float64:
		return Float64Type
	case int64:
		return Int64Type
	case float32:
		return Float32Type
	case int:
		return IntType
	default:
		return NaNType
	}
}

func toType(v interface{}, returnType int) {
	var a interface{}

	// covert to either a int64 or float64
	switch v.(type) {
	case string:
		if i, err := strconv.ParserInt64(v.(string)); err == nil {
			a = i
		} else if f, err := strconv.ParseFloat64(v.(string); err == nil {
			a = f
		}
	case json.Number:
		if i, err := json.Number.ParserInt64(v.(string)); err == nil {
			a = i
		} else if f, err := json.Number.ParseFloat64(v.(string); err == nil {
			a = f
		}
	case float32:
		a = float64(v.(float32))
	case int:
		a = int64(v.(int))
	default:
		a = v


	}



	return a
}

func asNumbers(v1, v2 interface{}, returnType int) (interface{}, interface{}) {
	var a, b interface{}
	a = toType(v1, returnType)
	b = toType(v2, returnType)
	return a, b
}

// Converts a string to a signed float64 or int64
func (n Number) stringToNumber(s string) {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		n.Value = f
		return
	}
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		n.Value = i
		return
	}
	n.NaN = true
}

func (n Number) float64ToNumber(f float64) {
	n.Value = float64(f)
}

func (n Number) int64ToNumber(i int64) {
	n.Value = int64(i)
}

// NewNumber takes a value and returns a new Number struct
func NewNumber(v interface{}) Number {
	n := Number{
		Value: 0,
		NaN:   false,
	}
	switch v.(type) {
	case json.Number:
		n.stringToNumber(string(v.(json.Number)))
	case string:
		n.stringToNumber(v.(string))
	case float64:
		n.Value = float64(v.(float64))
	case float32:
		n.Value = float32(v.(float32))
	case int64:
		n.Value = int64(v.(int64))
	case int:
		n.Value = int(v.(int))
	default:
		n.NaN = true
	}
	return n
}

// IsValidNumber returns true if NaN is false, true otherwise
func (n Number) IsValidNumber() bool {
	return (n.NaN == false)
}

// IsZero checks to see if a Number is zero
func (n Number) IsZero() bool {
	switch n.Value.(type) {
	case int:
		if n.Value == int(0) {
			return true
		}
	case int64:
		if n.Value == int64(0) {
			return true
		}
	case float32:
		if n.Value == float32(0) {
			return true
		}
	case float64:
		if n.Value == float64(0) {
			return true
		}
	}
	return false
}

func (n Number) IsInt64() bool {
	switch n.Value.(type) {
	case int64:
		return true
	default:
		return false
	}
}

func (n Number) IsInt() bool {
	switch n.Value.(type) {
	case int:
		return true
	default:
		return false
	}
}

func (n Number) IsFloat32() bool {
	switch n.Value.(type) {
	case float32:
		return true
	default:
		return false
	}
}

func (n Number) IsFloat64() bool {
	switch n.Value.(type) {
	case float64:
		return true
	default:
		return false
	}
}

// Float converts the internal representation to a Go float64
func (n Number) Float64() float64 {
	if n.NaN == true {
		return float64(0.0)
	}
	switch n.Value.(type) {
	case float64:
		return n.Value.(float64)
	case int64:
		return float64(n.Value.(int64))
	case int:
		return float64(n.Value.(int))
	case float32:
		return float64(n.Value.(float64))
	default:
		return float64(0.0)
	}
}

// Float converts the internal representation to a Go float64
func (n Number) Float32() float32 {
	if n.NaN == true {
		return float32(0.0)
	}
	switch n.Value.(type) {
	case float64:
		return float32(n.Value.(float64))
	case int64:
		return float32(n.Value.(int64))
	case int:
		return float32(n.Value.(int))
	case float32:
		return n.Value.(float32)
	default:
		return float32(0.0)
	}
}

// Int64 converts the internal representation to a Go int64
func (n Number) Int64() int64 {
	if n.NaN == true {
		return int64(0)
	}
	switch n.Value.(type) {
	case int64:
		return n.Value.(int64)
	case int:
		return int64(n.Value.(int64))
	case float32:
		return int64(n.Value.(float32))
	case float64:
		return int64(n.Value.(float64))
	default:
		return int64(0)
	}
}

// Int convert the internal representation to a Go int
func (n Number) Int() int {
	if n.NaN == true {
		return 0
	}
	switch n.Value.(type) {
	case int64:
		return int(n.Value.(int64))
	case int:
		return n.Value.(int)
	case float64:
		return int(n.Value.(float64))
	case float32:
		return int(n.Value.(float32))
	default:
		return int(0)
	}
}

// String converts n to its string presentation
func (n Number) String() string {
	if n.NaN == true {
		return "NaN"
	}
	switch n.Value.(type) {
	case int64:
		return fmt.Sprintf("%d", n.Value)
	case int:
		return fmt.Sprintf("%d", n.Value)
	case float64:
		return fmt.Sprintf("%g", n.Value)
	case float32:
		return fmt.Sprintf("%g", n.Value)
	default:
		return "NaN"
	}
}

func IsGreater(a, b Number) bool {
	if a.NaN == true || b.NaN == true {
		return false
	}
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return (a.Value.(float64) > b.Value.(float64))
	case aType == Float32Type && bType == Float32Type:
		return (a.Value.(float32) > b.Value.(float32))
	case aType == Int64Type && bType == Int64Type:
		return (a.Value.(int64) > b.Value.(int64))
	case aType == IntType && bType == IntType:
		return (a.Value.(int) > b.Value.(int))
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return (a.Value.(float64) > float64(b.Value.(float32)))
	case aType == Float64Type && bType == Int64Type:
		return (a.Value.(float64) > float64(b.Value.(int64)))
	case aType == Float64Type && bType == IntType:
		return (a.Value.(float64) > float64(b.Value.(int)))
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return (a.Value.(float32) > float32(b.Value.(float64)))
	case aType == Float32Type && bType == Int64Type:
		return (a.Value.(float32) > float32(b.Value.(int64)))
	case aType == Float32Type && bType == IntType:
		return (a.Value.(float32) > float32(b.Value.(int)))
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return (a.Value.(int64) > int64(b.Value.(float64)))
	case aType == Int64Type && bType == Float32Type:
		return (a.Value.(int64) > int64(b.Value.(float32)))
	case aType == Int64Type && bType == IntType:
		return (a.Value.(int64) > int64(b.Value.(int)))
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return (a.Value.(int) > int(b.Value.(float64)))
	case aType == IntType && bType == Float32Type:
		return (a.Value.(int) > int(b.Value.(float32)))
	case aType == IntType && bType == Int64Type:
		return (a.Value.(int) > int(b.Value.(int64)))
	default:
		// NOTE: if not a number then return false
		return false
	}
}

func IsLess(v1, v2 interface{}) bool {
	a, b := asNumbers(v1, v2)
	if a.NaN == true || b.NaN == true {
		return false
	}
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return (a.Value.(float64) < b.Value.(float64))
	case aType == Float32Type && bType == Float32Type:
		return (a.Value.(float32) < b.Value.(float32))
	case aType == Int64Type && bType == Int64Type:
		return (a.Value.(int64) < b.Value.(int64))
	case aType == IntType && bType == IntType:
		return (a.Value.(int) < b.Value.(int))
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return (a.Value.(float64) < float64(b.Value.(float32)))
	case aType == Float64Type && bType == Int64Type:
		return (a.Value.(float64) < float64(b.Value.(int64)))
	case aType == Float64Type && bType == IntType:
		return (a.Value.(float64) < float64(b.Value.(int)))
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return (a.Value.(float32) < float32(b.Value.(float64)))
	case aType == Float32Type && bType == Int64Type:
		return (a.Value.(float32) < float32(b.Value.(int64)))
	case aType == Float32Type && bType == IntType:
		return (a.Value.(float32) < float32(b.Value.(int)))
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return (a.Value.(int64) < int64(b.Value.(float64)))
	case aType == Int64Type && bType == Float32Type:
		return (a.Value.(int64) < int64(b.Value.(float32)))
	case aType == Int64Type && bType == IntType:
		return (a.Value.(int64) < int64(b.Value.(int)))
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return (a.Value.(int) < int(b.Value.(float64)))
	case aType == IntType && bType == Float32Type:
		return (a.Value.(int) < int(b.Value.(float32)))
	case aType == IntType && bType == Int64Type:
		return (a.Value.(int) < int(b.Value.(int64)))
	default:
		// NOTE: if not a number then return false
		return false
	}
}

func IsEqual(v1, v2 interface{}) bool {
	a, b := asNumbers(v1, v2)
	if a.NaN == true || b.NaN == true {
		return false
	}
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return (a.Value.(float64) == b.Value.(float64))
	case aType == Float32Type && bType == Float32Type:
		return (a.Value.(float32) == b.Value.(float32))
	case aType == Int64Type && bType == Int64Type:
		return (a.Value.(int64) == b.Value.(int64))
	case aType == IntType && bType == IntType:
		return (a.Value.(int) == b.Value.(int))
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return (a.Value.(float64) == float64(b.Value.(float32)))
	case aType == Float64Type && bType == Int64Type:
		return (a.Value.(float64) == float64(b.Value.(int64)))
	case aType == Float64Type && bType == IntType:
		return (a.Value.(float64) == float64(b.Value.(int)))
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return (a.Value.(float32) == float32(b.Value.(float64)))
	case aType == Float32Type && bType == Int64Type:
		return (a.Value.(float32) == float32(b.Value.(int64)))
	case aType == Float32Type && bType == IntType:
		return (a.Value.(float32) == float32(b.Value.(int)))
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return (a.Value.(int64) == int64(b.Value.(float64)))
	case aType == Int64Type && bType == Float32Type:
		return (a.Value.(int64) == int64(b.Value.(float32)))
	case aType == Int64Type && bType == IntType:
		return (a.Value.(int64) == int64(b.Value.(int)))
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return (a.Value.(int) == int(b.Value.(float64)))
	case aType == IntType && bType == Float32Type:
		return (a.Value.(int) == int(b.Value.(float32)))
	case aType == IntType && bType == Int64Type:
		return (a.Value.(int) == int(b.Value.(int64)))
	default:
		// NOTE: if not a number then return false
		return false
	}
}

func Add(v1, v2 interface{}) Number {
	a, b := asNumbers(v1, v2)
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float64) + b.Value.(float64)),
			NaN:   false,
		}
	case aType == Float32Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float32) + b.Value.(float32)),
			NaN:   false,
		}

	case aType == Int64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(int64) + b.Value.(int64)),
			NaN:   false,
		}
	case aType == IntType && bType == IntType:
		return Number{
			Value: (a.Value.(int) + b.Value.(int)),
			NaN:   false,
		}
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float64) + float64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Float64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float64) + float64(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float64Type && bType == IntType:
		return Number{
			Value: (a.Value.(float64) + float64(b.Value.(int))),
			NaN:   false,
		}
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float32) + float32(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float32) + float32(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == IntType:
		return Number{
			Value: (a.Value.(float32) + float32(b.Value.(int))),
			NaN:   false,
		}
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(int64) + int64(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Int64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(int64) + int64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Int64Type && bType == IntType:
		return Number{
			Value: (a.Value.(int64) + int64(b.Value.(int))),
			NaN:   false,
		}
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return Number{
			Value: (a.Value.(int) + int(b.Value.(float64))),
			NaN:   false,
		}
	case aType == IntType && bType == Float32Type:
		return Number{
			Value: (a.Value.(int) + int(b.Value.(float32))),
			NaN:   false,
		}
	case aType == IntType && bType == Int64Type:
		return Number{
			Value: (a.Value.(int) + int(b.Value.(int64))),
			NaN:   false,
		}
	default:
		return Number{
			Value: 0,
			NaN:   true,
		}
	}
}

func Subtract(v1, v2 interface{}) Number {
	a, b := asNumbers(v1, v2)
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float64) - b.Value.(float64)),
			NaN:   false,
		}
	case aType == Float32Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float32) - b.Value.(float32)),
			NaN:   false,
		}

	case aType == Int64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(int64) - b.Value.(int64)),
			NaN:   false,
		}
	case aType == IntType && bType == IntType:
		return Number{
			Value: (a.Value.(int) - b.Value.(int)),
			NaN:   false,
		}
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float64) - float64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Float64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float64) - float64(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float64Type && bType == IntType:
		return Number{
			Value: (a.Value.(float64) - float64(b.Value.(int))),
			NaN:   false,
		}
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float32) - float32(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float32) - float32(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == IntType:
		return Number{
			Value: (a.Value.(float32) - float32(b.Value.(int))),
			NaN:   false,
		}
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(int64) - int64(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Int64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(int64) - int64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Int64Type && bType == IntType:
		return Number{
			Value: (a.Value.(int64) - int64(b.Value.(int))),
			NaN:   false,
		}
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return Number{
			Value: (a.Value.(int) - int(b.Value.(float64))),
			NaN:   false,
		}
	case aType == IntType && bType == Float32Type:
		return Number{
			Value: (a.Value.(int) - int(b.Value.(float32))),
			NaN:   false,
		}
	case aType == IntType && bType == Int64Type:
		return Number{
			Value: (a.Value.(int) - int(b.Value.(int64))),
			NaN:   false,
		}
	default:
		return Number{
			Value: 0,
			NaN:   true,
		}
	}
}

func Multiply(v1, v2 interface{}) Number {
	a, b := asNumbers(v1, v2)
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float64) * b.Value.(float64)),
			NaN:   false,
		}
	case aType == Float32Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float32) * b.Value.(float32)),
			NaN:   false,
		}

	case aType == Int64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(int64) * b.Value.(int64)),
			NaN:   false,
		}
	case aType == IntType && bType == IntType:
		return Number{
			Value: (a.Value.(int) * b.Value.(int)),
			NaN:   false,
		}
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float64) * float64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Float64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float64) * float64(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float64Type && bType == IntType:
		return Number{
			Value: (a.Value.(float64) * float64(b.Value.(int))),
			NaN:   false,
		}
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float32) * float32(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float32) * float32(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == IntType:
		return Number{
			Value: (a.Value.(float32) * float32(b.Value.(int))),
			NaN:   false,
		}
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(int64) * int64(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Int64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(int64) * int64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Int64Type && bType == IntType:
		return Number{
			Value: (a.Value.(int64) * int64(b.Value.(int))),
			NaN:   false,
		}
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return Number{
			Value: (a.Value.(int) * int(b.Value.(float64))),
			NaN:   false,
		}
	case aType == IntType && bType == Float32Type:
		return Number{
			Value: (a.Value.(int) * int(b.Value.(float32))),
			NaN:   false,
		}
	case aType == IntType && bType == Int64Type:
		return Number{
			Value: (a.Value.(int) * int(b.Value.(int64))),
			NaN:   false,
		}
	default:
		return Number{
			Value: 0,
			NaN:   true,
		}
	}
}

func Divide(v1, v2 Number) Number {
	a, b := asNumbers(v1, v2)
	if b.IsZero() == true {
		return Number{
			Value: int64(0),
			NaN:   true,
		}
	}
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Float64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float64) / b.Value.(float64)),
			NaN:   false,
		}
	case aType == Float32Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float32) / b.Value.(float32)),
			NaN:   false,
		}

	case aType == Int64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(int64) / b.Value.(int64)),
			NaN:   false,
		}
	case aType == IntType && bType == IntType:
		return Number{
			Value: (a.Value.(int) / b.Value.(int)),
			NaN:   false,
		}
	// Permutations float64 plus other types
	case aType == Float64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(float64) / float64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Float64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float64) / float64(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float64Type && bType == IntType:
		return Number{
			Value: (a.Value.(float64) / float64(b.Value.(int))),
			NaN:   false,
		}
	// Permutations float32 plus other types
	case aType == Float32Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(float32) / float32(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(float32) / float32(b.Value.(int64))),
			NaN:   false,
		}
	case aType == Float32Type && bType == IntType:
		return Number{
			Value: (a.Value.(float32) / float32(b.Value.(int))),
			NaN:   false,
		}
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(int64) / int64(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Int64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(int64) / int64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Int64Type && bType == IntType:
		return Number{
			Value: (a.Value.(int64) / int64(b.Value.(int))),
			NaN:   false,
		}
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return Number{
			Value: (a.Value.(int) / int(b.Value.(float64))),
			NaN:   false,
		}
	case aType == IntType && bType == Float32Type:
		return Number{
			Value: (a.Value.(int) / int(b.Value.(float32))),
			NaN:   false,
		}
	case aType == IntType && bType == Int64Type:
		return Number{
			Value: (a.Value.(int) / int(b.Value.(int64))),
			NaN:   false,
		}
	default:
		return Number{
			Value: 0,
			NaN:   true,
		}
	}
}

func Modulo(v1, v2 interface{}) Number {
	a, b := asNumbers(v1, v2)
	aType := NumberType(a)
	bType := NumberType(b)

	switch {
	// Types match
	case aType == Int64Type && bType == Int64Type:
		return Number{
			Value: (a.Value.(int64) % b.Value.(int64)),
			NaN:   false,
		}
	case aType == IntType && bType == IntType:
		return Number{
			Value: (a.Value.(int) % b.Value.(int)),
			NaN:   false,
		}
	// Permutations Int64 plus other types
	case aType == Int64Type && bType == Float64Type:
		return Number{
			Value: (a.Value.(int64) % int64(b.Value.(float64))),
			NaN:   false,
		}
	case aType == Int64Type && bType == Float32Type:
		return Number{
			Value: (a.Value.(int64) % int64(b.Value.(float32))),
			NaN:   false,
		}
	case aType == Int64Type && bType == IntType:
		return Number{
			Value: (a.Value.(int64) % int64(b.Value.(int))),
			NaN:   false,
		}
		// Permutatins Int plus other types
	case aType == IntType && bType == Float64Type:
		return Number{
			Value: (a.Value.(int) % int(b.Value.(float64))),
			NaN:   false,
		}
	case aType == IntType && bType == Float32Type:
		return Number{
			Value: (a.Value.(int) % int(b.Value.(float32))),
			NaN:   false,
		}
	case aType == IntType && bType == Int64Type:
		return Number{
			Value: (a.Value.(int) % int(b.Value.(int64))),
			NaN:   false,
		}
	default:
		return Number{
			Value: 0,
			NaN:   true,
		}
	}
}
