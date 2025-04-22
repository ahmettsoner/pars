package convert

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func ToBool(in interface{}) bool {
	if b, ok := in.(bool); ok {
		return b
	}

	if str, ok := in.(string); ok {
		str = strings.ToLower(str)
		switch str {
		case "1", "true":
			return true
		default:
			f, _ := strToFloat64(str)
			return f == 1
		}
	}

	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int() == 1
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return val.Uint() == 1
	case reflect.Float32, reflect.Float64:
		return val.Float() == 1
	default:
		return false
	}
}

func ToString(in interface{}) string {
	if in == nil {
		return "nil"
	}
	if s, ok := in.(string); ok {
		return s
	}
	if s, ok := in.(fmt.Stringer); ok {
		return s.String()
	}
	if s, ok := in.([]byte); ok {
		return string(s)
	}

	v, ok := printableValue(reflect.ValueOf(in))
	if ok {
		in = v
	}

	return fmt.Sprint(in)
}

func ToInt64(v interface{}) (int64, error) {
	if str, ok := v.(string); ok {
		return strToInt64(str)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int(), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return int64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		tv := val.Uint()

		// this can overflow and give -1, but IMO this is better than
		// returning maxint64
		return int64(tv), nil
	case reflect.Float32, reflect.Float64:
		return int64(val.Float()), nil
	case reflect.Bool:
		if val.Bool() {
			return 1, nil
		}

		return 0, nil
	default:
		return 0, fmt.Errorf("could not convert %v to int64", v)
	}
}

func ToInt(in interface{}) (int, error) {
	i, err := ToInt64(in)
	if err != nil {
		return 0, err
	}

	// Bounds-checking to protect against CWE-190 and CWE-681
	// https://cwe.mitre.org/data/definitions/190.html
	// https://cwe.mitre.org/data/definitions/681.html
	if i >= math.MinInt && i <= math.MaxInt {
		return int(i), nil
	}

	// maybe we're on a 32-bit system, so we can't represent this number
	return 0, fmt.Errorf("could not convert %v to int", in)
}

func ToInt64s(in ...interface{}) ([]int64, error) {
	out := make([]int64, len(in))
	for i, v := range in {
		n, err := ToInt64(v)
		if err != nil {
			return nil, err
		}

		out[i] = n
	}

	return out, nil
}

func ToInts(in ...interface{}) ([]int, error) {
	out := make([]int, len(in))
	for i, v := range in {
		n, err := ToInt(v)
		if err != nil {
			return nil, err
		}

		out[i] = n
	}

	return out, nil
}

// ToFloat64 - convert input to a float64, if convertible. Otherwise, errors.
func ToFloat64(v interface{}) (float64, error) {
	if str, ok := v.(string); ok {
		return strToFloat64(str)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return float64(val.Int()), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return float64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Bool:
		if val.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("could not convert %v to float64", v)
	}
}

func strToInt64(str string) (int64, error) {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}

	iv, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		// maybe it's a float?
		var fv float64
		fv, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("could not convert %q to int64: %w", str, err)
		}

		return ToInt64(fv)
	}

	return iv, nil
}

func strToFloat64(str string) (float64, error) {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}

	// this is inefficient, but it's the only way I can think of to
	// properly convert octal integers to floats
	iv, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		// ok maybe it's a float?
		var fv float64
		fv, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("could not convert %q to float64: %w", str, err)
		}

		return fv, nil
	}

	return float64(iv), nil
}

// ToFloat64s -
func ToFloat64s(in ...interface{}) ([]float64, error) {
	out := make([]float64, len(in))
	for i, v := range in {
		f, err := ToFloat64(v)
		if err != nil {
			return nil, err
		}
		out[i] = f
	}

	return out, nil
}

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

// printableValue returns the, possibly indirected, interface value inside v that
// is best for a call to formatted printer.
func printableValue(v reflect.Value) (interface{}, bool) {
	if v.Kind() == reflect.Ptr {
		v, _ = indirect(v) // fmt.Fprint handles nil.
	}
	if !v.IsValid() {
		return "<no value>", true
	}

	if !v.Type().Implements(errorType) && !v.Type().Implements(fmtStringerType) {
		if v.CanAddr() && (reflect.PtrTo(v.Type()).Implements(errorType) || reflect.PtrTo(v.Type()).Implements(fmtStringerType)) {
			v = v.Addr()
		} else {
			switch v.Kind() {
			case reflect.Chan, reflect.Func:
				return nil, false
			}
		}
	}
	return v.Interface(), true
}

// indirect returns the item at the end of indirection, and a bool to indicate if it's nil.
func indirect(v reflect.Value) (rv reflect.Value, isNil bool) {
	for ; v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface; v = v.Elem() {
		if v.IsNil() {
			return v, true
		}
	}
	return v, false
}
