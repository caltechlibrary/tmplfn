package numbers

import (
	"encoding/json"
	"log"
)

// tmplfn supports calculations with four types of numbers,
// int64, int, float64 and float32. In Add, Substract, Mutliply, Divide,
// and Modulo the input values are normalized to the highest bit width of the two.
// If the input value is a string or json.Number it is normalized to either float64 or int64 or
// zero if parse fails.
const (
	naNType = iota
	intType
	int64Type
	float32Type
	float64Type
	jsonNumberType
)

// numberType returns the best guess of numeric types this package supports
func numberType(value interface{}) int {
	switch value.(type) {
	case json.Number:
		return jsonNumberType
	case float64:
		return float64Type
	case int64:
		return int64Type
	case float32:
		return float32Type
	case int:
		return intType
	default:
		return naNType
	}
}

// toType converts an interface to the targetType (must be a float64Type, float32Type, int64Type or intType)
func toType(v interface{}, targetType int) interface{} {
	var (
		a     interface{}
		nType int
	)

	// normalize to either a float64, float32, int64 or int
	switch v.(type) {
	case json.Number:
		a = 0
		nType = int64Type
		if i, err := v.(json.Number).Int64(); err == nil {
			a = i
			nType = int64Type
		} else if f, err := v.(json.Number).Float64(); err == nil {
			a = f
			nType = float64Type
		}
	case float64:
		a = v.(float64)
		nType = float64Type
	case float32:
		a = v.(float32)
		nType = float32Type
	case int64:
		a = v.(int64)
		nType = int64Type
	case int:
		a = v.(int)
		nType = intType
	default:
		// NOTE: If it is not a supported type then treat value as zero
		a = int(0)
		nType = intType
	}

	// now convert to target type
	switch targetType {
	case float64Type:
		if nType == int64Type {
			return float64(a.(int64))
		} else if nType == intType {
			return float64(a.(int))
		} else if nType == float32Type {
			return float64(a.(float32))
		}
		return a.(float64)
	case int64Type:
		if nType == int64Type {
			return a.(int64)
		} else if nType == intType {
			return int64(a.(int))
		} else if nType == float32Type {
			return int64(a.(float32))
		}
		return int64(a.(float64))
	case float32Type:
		if nType == int64Type {
			return float32(a.(int64))
		} else if nType == intType {
			return float32(a.(int))
		} else if nType == float32Type {
			return a.(float32)
		}
		return float32(a.(float64))
	default:
		// intType is the default type
		if nType == int64Type {
			return int(a.(int64))
		} else if nType == intType {
			return a.(int)
		} else if nType == float32Type {
			return int(a.(float32))
		}
		return int(a.(float64))
	}
}

// normalizeJSONNumber converts to either a float64 or int64 or zero
func normalizeJSONNumberType(n json.Number) int {
	if _, err := n.Float64(); err == nil {
		return float64Type
	}
	return int64Type
}

// normalizeNumbers takes to interface values and promotes both to float64 if any floats present
// otherwise promotes to int64, returns normalized values with type chosen
func normalizeNumbers(v1, v2 interface{}) (interface{}, interface{}, int) {
	var (
		a, b                interface{}
		aType, bType, nType int
	)
	aType = numberType(v1)
	bType = numberType(v2)
	// NOTE: If type is jsonNumberType then convert to float64 or int64 before normalizing
	if aType == jsonNumberType {
		aType = normalizeJSONNumberType(v1.(json.Number))
	}
	if bType == jsonNumberType {
		bType = normalizeJSONNumberType(v2.(json.Number))
	}
	switch {
	case aType == float64Type || bType == float64Type:
		a = toType(v1, float64Type)
		b = toType(v2, float64Type)
		nType = float64Type
	case aType == float32Type || bType == float32Type:
		a = toType(v1, float32Type)
		b = toType(v2, float32Type)
		nType = float32Type
	case aType == int64Type || bType == int64Type:
		a = toType(v1, int64Type)
		b = toType(v2, int64Type)
		nType = int64Type
	default:
		a = toType(v1, intType)
		b = toType(v2, intType)
		nType = intType
	}
	return a, b, nType
}

// IsNumber checks to see if the input type can render to a float64, float32, int64 or int
func IsNumber(v interface{}) bool {
	switch v.(type) {
	case json.Number:
		if _, err := v.(json.Number).Float64(); err == nil {
			return true
		}
		if _, err := v.(json.Number).Int64(); err == nil {
			return true
		}
	case float64:
		return true
	case float32:
		return true
	case int64:
		return true
	case int:
		return true
	}
	return false
}

// IsZero checks to see if a Number is zero
func IsZero(v interface{}) bool {
	switch v.(type) {
	case int:
		if v == int(0) {
			return true
		}
	case int64:
		if v == int64(0) {
			return true
		}
	case float32:
		if v == float32(0) {
			return true
		}
	case float64:
		if v == float64(0) {
			return true
		}
	}
	return false
}

// Add v1 and v2 or return zero if type issue
func Add(v1, v2 interface{}) interface{} {
	a, b, nType := normalizeNumbers(v1, v2)
	switch nType {
	case intType:
		return a.(int) + b.(int)
	case float32Type:
		return a.(float32) + b.(float32)
	case int64Type:
		return a.(int64) + b.(int64)
	case float64Type:
		return a.(float64) + b.(float64)
	default:
		return 0
	}
}

// Substract v2 from v1 or return zero if a type issue
func Subtract(v1, v2 interface{}) interface{} {
	a, b, nType := normalizeNumbers(v1, v2)
	switch nType {
	case intType:
		return a.(int) - b.(int)
	case float32Type:
		return a.(float32) - b.(float32)
	case int64Type:
		return a.(int64) - b.(int64)
	case float64Type:
		return a.(float64) - b.(float64)
	default:
		return 0
	}
}

// Multiply v1 by v2 returning the result or zero if type issue
func Multiply(v1, v2 interface{}) interface{} {
	a, b, nType := normalizeNumbers(v1, v2)
	switch nType {
	case intType:
		return a.(int) * b.(int)
	case float32Type:
		return a.(float32) * b.(float32)
	case int64Type:
		return a.(int64) * b.(int64)
	case float64Type:
		return a.(float64) * b.(float64)
	default:
		return 0
	}
}

// Divide v1 by v2 for non-zero v2 or return zero
func Divide(v1, v2 interface{}) interface{} {
	a, b, nType := normalizeNumbers(v1, v2)
	//FIXME: this is ugly, divide by zero error will be logged
	if IsZero(b) {
		log.Printf("Divide by Zero error for values %T %v, %T %v\n", a, a, b, b)
		return 0
	}
	switch nType {
	case intType:
		return a.(int) / b.(int)
	case float32Type:
		return a.(float32) / b.(float32)
	case int64Type:
		return a.(int64) / b.(int64)
	case float64Type:
		return a.(float64) / b.(float64)
	default:
		return 0
	}
}

// Modulo returns the modulo of int or int64 or zero
func Modulo(v1, v2 interface{}) interface{} {
	a, b, nType := normalizeNumbers(v1, v2)
	switch nType {
	case intType:
		return a.(int) % b.(int)
	case int64Type:
		return a.(int64) % b.(int64)
	default:
		return 0
	}
}

// Addi adds two values return an int or zero
func Addi(v1, v2 interface{}) int {
	v := Add(v1, v2)
	return toType(v, intType).(int)
}

// Subi adds two values return an int or zero
func Subi(v1, v2 interface{}) int {
	v := Subtract(v1, v2)
	return toType(v, intType).(int)
}

// Int64 returns an int64 for value provided
func Int64(v interface{}) int64 {
	return toType(v, int64Type).(int64)
}

// Int returns a int for value provided
func Int(v interface{}) int {
	return toType(v, intType).(int)
}

// Float32 returns a float32 for value provided
func Float32(v interface{}) float32 {
	return toType(v, float32Type).(float32)
}

// Float64 returns a float64 for value provided
func Float64(v interface{}) float64 {
	return toType(v, float64Type).(float64)
}
