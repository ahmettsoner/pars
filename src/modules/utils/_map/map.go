package _map

import (
	"fmt"
	"reflect"
	"sort"

	"parsdevkit.net/core/utils/convert"
)

func Has(in interface{}, key interface{}) bool {
	av := reflect.ValueOf(in)

	switch av.Kind() {
	case reflect.Map:
		kv := reflect.ValueOf(key)
		return av.MapIndex(kv).IsValid()
	case reflect.Slice, reflect.Array:
		l := av.Len()
		for i := 0; i < l; i++ {
			v := av.Index(i).Interface()
			if reflect.DeepEqual(v, key) {
				return true
			}
		}
	}

	return false
}

func Dictionary(v ...interface{}) map[string]interface{} {
	if len(v)%2 != 0 {
		panic("MakeMap: odd number of arguments")
	}
	m := make(map[string]interface{})
	for i := 0; i < len(v); i += 2 {
		key := convert.ToString(v[i])
		value := v[i+1]
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			value = reflect.ValueOf(value).Elem().Interface()
		}
		m[key] = value
	}
	return m
}
func HasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

// Keys returns the list of keys in one or more maps. The returned list of keys
// is ordered by map, each in sorted key order.
func Keys(in ...map[string]interface{}) ([]string, error) {
	if len(in) == 0 {
		return nil, fmt.Errorf("need at least one argument")
	}
	keys := []string{}
	for _, m := range in {
		k, _ := splitMap(m)
		keys = append(keys, k...)
	}
	return keys, nil
}

func splitMap(m map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, len(m))
	values := make([]interface{}, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i, k := range keys {
		values[i] = m[k]
	}
	return keys, values
}

// Values returns the list of values in one or more maps. The returned list of values
// is ordered by map, each in sorted key order. If the Keys function is called with
// the same arguments, the key/value mappings will be maintained.
func Values(in ...map[string]interface{}) ([]interface{}, error) {
	if len(in) == 0 {
		return nil, fmt.Errorf("need at least one argument")
	}
	values := []interface{}{}
	for _, m := range in {
		_, v := splitMap(m)
		values = append(values, v...)
	}
	return values, nil
}

// Merge source maps (srcs) into dst. Precedence is in left-to-right order, with
// the left-most values taking precedence over the right-most.
func Merge(dst map[string]interface{}, srcs ...map[string]interface{}) (map[string]interface{}, error) {
	for _, src := range srcs {
		dst = mergeValues(src, dst)
	}
	return dst, nil
}

// returns whether or not a contains v
func contains(v string, a []string) bool {
	for _, n := range a {
		if n == v {
			return true
		}
	}
	return false
}

// Omit returns a new map without any entries that have the
// given keys (inverse of Pick).
func Omit(in map[string]interface{}, keys ...string) map[string]interface{} {
	out := map[string]interface{}{}
	for k, v := range in {
		if !contains(k, keys) {
			out[k] = v
		}
	}
	return out
}

// Pick returns a new map with any entries that have the
// given keys (inverse of Omit).
func Pick(in map[string]interface{}, keys ...string) map[string]interface{} {
	out := map[string]interface{}{}
	for k, v := range in {
		if contains(k, keys) {
			out[k] = v
		}
	}
	return out
}

func copyMap(m map[string]interface{}) map[string]interface{} {
	n := map[string]interface{}{}
	for k, v := range m {
		n[k] = v
	}
	return n
}

// Merges a default and override map
func mergeValues(d map[string]interface{}, o map[string]interface{}) map[string]interface{} {
	def := copyMap(d)
	over := copyMap(o)
	for k, v := range over {
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := def[k]; !exists {
			def[k] = v
			continue
		}
		nextMap, ok := v.(map[string]interface{})
		// If it isn't another map, overwrite the value
		if !ok {
			def[k] = v
			continue
		}
		// Edge case: If the key exists in the default, but isn't a map
		defMap, isMap := def[k].(map[string]interface{})
		// If the override map has a map for this key, prefer it
		if !isMap {
			def[k] = v
			continue
		}
		// If we got to this point, it is a map in both, so merge them
		def[k] = mergeValues(defMap, nextMap)
	}
	return def
}
