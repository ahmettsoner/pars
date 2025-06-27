package functions

import (
	"fmt"

	"parsdevkit.net/core/utils/_map"
)

type MapFuncs struct{}

func (c MapFuncs) Has(in interface{}, key interface{}) bool {
	return _map.Has(in, key)
}

func (c MapFuncs) Dictionary(v ...interface{}) map[string]interface{} {
	return _map.Dictionary(v...)
}
func (c MapFuncs) HasKey(m map[string]interface{}, key string) bool {
	return _map.HasKey(m, key)
}

func (c MapFuncs) Keys(in ...map[string]interface{}) []string {
	result, err := _map.Keys(in...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c MapFuncs) Values(in ...map[string]interface{}) []interface{} {
	result, err := _map.Values(in...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c MapFuncs) Merge(dst map[string]interface{}, srcs ...map[string]interface{}) map[string]interface{} {
	result, err := _map.Merge(dst, srcs...)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (c MapFuncs) Omit(in map[string]interface{}, keys ...string) map[string]interface{} {
	return _map.Omit(in, keys...)
}

func (c MapFuncs) Pick(in map[string]interface{}, keys ...string) map[string]interface{} {
	return _map.Pick(in, keys...)
}
