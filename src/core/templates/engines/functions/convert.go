package functions

import (
	"fmt"

	"parsdevkit.net/core/utils/convert"
)

type ConvertFuncs struct{}

func (c ConvertFuncs) ToBool(in interface{}) bool {
	return convert.ToBool(in)
}

func (c ConvertFuncs) ToString(in interface{}) string {
	return convert.ToString(in)
}

func (c ConvertFuncs) ToInt64(v interface{}) int64 {
	result, err := convert.ToInt64(v)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ConvertFuncs) ToInt(in interface{}) int {
	result, err := convert.ToInt(in)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ConvertFuncs) ToInt64s(in ...interface{}) []int64 {
	result, err := convert.ToInt64s(in...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ConvertFuncs) ToInts(in ...interface{}) []int {
	result, err := convert.ToInts(in...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ConvertFuncs) ToFloat64(v interface{}) float64 {
	result, err := convert.ToFloat64(v)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ConvertFuncs) ToFloat64s(in ...interface{}) []float64 {
	result, err := convert.ToFloat64s(in...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}
