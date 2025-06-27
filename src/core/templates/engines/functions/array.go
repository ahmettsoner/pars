package functions

import (
	"fmt"

	"parsdevkit.net/core/utils/array"
)

type ArrayFuncs struct{}

func (c ArrayFuncs) Slice(args ...interface{}) []interface{} {
	return array.Slice(args...)
}

func (c ArrayFuncs) Has(in interface{}, key interface{}) bool {
	return array.Has(in, key)
}

func (c ArrayFuncs) Append(v interface{}, list interface{}) []interface{} {
	result, err := array.Append(v, list)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) Prepend(v interface{}, list interface{}) []interface{} {
	result, err := array.Prepend(v, list)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) Uniq(list interface{}) []interface{} {
	result, err := array.Uniq(list)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) Reverse(list interface{}) []interface{} {
	result, err := array.Reverse(list)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) Sort(key string, list interface{}) (out []interface{}) {
	result, err := array.Sort(key, list)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) Flatten(list interface{}, depth int) []interface{} {
	result, err := array.Flatten(list, depth)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c ArrayFuncs) IsFirst(index int, slice interface{}) bool {
	return array.IsFirst(index, slice)
}

func (c ArrayFuncs) IsLast(index int, slice interface{}) bool {
	return array.IsLast(index, slice)
}

func (c ArrayFuncs) HasElements(slice interface{}) bool {
	return array.HasElements(slice)
}

func (c ArrayFuncs) Count(slice interface{}) int {
	return array.Count(slice)
}
func (c ArrayFuncs) First(slice interface{}) interface{} {
	return array.First(slice)
}

func (c ArrayFuncs) Last(slice interface{}) interface{} {
	return array.Last(slice)
}

func (c ArrayFuncs) Contains(slice interface{}, key string, value interface{}) bool {
	return array.Contains(slice, key, value)
}

func (c ArrayFuncs) Find(slice interface{}, key string, value interface{}) interface{} {
	return array.Find(slice, key, value)
}

func (c ArrayFuncs) Filter(data interface{}, path string, value string) []interface{} {
	return array.Filter(data, path, value)
}
