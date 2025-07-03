package math

import (
	"fmt"
	"math"
	"strconv"

	"parsdevkit.net/pkg/utilities/convert"
)

func Abs(x interface{}) float64 {
	return math.Abs(toFloat64(x))
}

func Add(values ...interface{}) float64 {
	sum := 0.0
	for _, value := range values {
		sum += toFloat64(value)
	}
	return sum
}

func Sub(x, y interface{}) float64 {
	return toFloat64(x) - toFloat64(y)
}

func Pow(x, y interface{}) float64 {
	return math.Pow(toFloat64(x), toFloat64(y))
}

func Mul(values ...interface{}) float64 {
	product := 1.0
	for _, value := range values {
		product *= toFloat64(value)
	}
	return product
}

func Div(x, y interface{}) float64 {
	if y == 0 {
		panic("division by zero")
	}
	return toFloat64(x) / toFloat64(y)
}

func Max(x, y interface{}) float64 {
	return math.Max(toFloat64(x), toFloat64(y))
}

func Min(x, y interface{}) float64 {
	return math.Min(toFloat64(x), toFloat64(y))
}

func Mod(x, y interface{}) float64 {
	return math.Mod(toFloat64(x), toFloat64(y))
}

func Round(x, y interface{}) float64 {
	return math.Round(toFloat64(x))
}

func Floor(x, y interface{}) float64 {
	return math.Floor(toFloat64(x))
}

func IsInt(n interface{}) bool {
	switch i := n.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		_, err := strconv.ParseInt(i, 0, 64)
		return err == nil
	}
	return false
}

func IsFloat(n interface{}) bool {
	switch i := n.(type) {
	case float32, float64:
		return true
	case string:
		_, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return false
		}
		if IsInt(i) {
			return false
		}
		return true
	}
	return false
}
func IsNum(n interface{}) bool {
	return IsInt(n) || IsFloat(n)
}

func toFloat64(value interface{}) float64 {
	res, err := convert.ToFloat64(value)
	if err != nil {
		panic(fmt.Sprintf("unsupported type: %T", value))
	}

	return res

}
