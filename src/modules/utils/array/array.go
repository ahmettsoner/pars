package array

import (
	"reflect"
	"sort"
	"strings"

	"parsdevkit.net/core/utils/common"
)

func Slice(args ...interface{}) []interface{} {
	return args
}

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

// Append v to the end of list. No matter what type of input slice or array list is, a new []interface{} is always returned.
func Append(v interface{}, list interface{}) ([]interface{}, error) {
	l, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}

	return append(l, v), nil
}

// Prepend v to the beginning of list. No matter what type of input slice or array list is, a new []interface{} is always returned.
func Prepend(v interface{}, list interface{}) ([]interface{}, error) {
	l, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}

	return append([]interface{}{v}, l...), nil
}

// Uniq finds the unique values within list. No matter what type of input slice or array list is, a new []interface{} is always returned.
func Uniq(list interface{}) ([]interface{}, error) {
	l, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}

	out := []interface{}{}
	for _, v := range l {
		if !Has(out, v) {
			out = append(out, v)
		}
	}
	return out, nil
}

// Reverse the list. No matter what type of input slice or array list is, a new []interface{} is always returned.
func Reverse(list interface{}) ([]interface{}, error) {
	l, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}

	// nifty trick from https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(l)-1; left < right; left, right = left+1, right-1 {
		l[left], l[right] = l[right], l[left]
	}
	return l, nil
}

// Sort a given array or slice. Uses natural sort order if possible. If a
// non-empty key is given and the list elements are maps, this will attempt to
// sort by the values of those entries.
//
// Does not modify the input list.
func Sort(key string, list interface{}) (out []interface{}, err error) {
	if list == nil {
		return nil, nil
	}

	ia, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}
	// if the types are all the same, we can sort the slice
	if sameTypes(ia) {
		s := make([]interface{}, len(ia))
		// make a copy so the original is unmodified
		copy(s, ia)
		sort.SliceStable(s, func(i, j int) bool {
			return lessThan(key)(s[i], s[j])
		})
		return s, nil
	}
	return ia, nil
}

// lessThan - compare two values of the same type
func lessThan(key string) func(left, right interface{}) bool {
	return func(left, right interface{}) bool {
		val := reflect.Indirect(reflect.ValueOf(left))
		rval := reflect.Indirect(reflect.ValueOf(right))
		switch val.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			return val.Int() < rval.Int()
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			return val.Uint() < rval.Uint()
		case reflect.Float32, reflect.Float64:
			return val.Float() < rval.Float()
		case reflect.String:
			return val.String() < rval.String()
		case reflect.MapOf(
			reflect.TypeOf(reflect.String),
			reflect.TypeOf(reflect.Interface),
		).Kind():
			kval := reflect.ValueOf(key)
			if !val.MapIndex(kval).IsValid() {
				return false
			}
			newleft := val.MapIndex(kval).Interface()
			newright := rval.MapIndex(kval).Interface()
			return lessThan("")(newleft, newright)
		case reflect.Struct:
			if !val.FieldByName(key).IsValid() {
				return false
			}
			newleft := val.FieldByName(key).Interface()
			newright := rval.FieldByName(key).Interface()
			return lessThan("")(newleft, newright)
		default:
			// it's not really comparable, so...
			return false
		}
	}
}

func sameTypes(a []interface{}) bool {
	var t reflect.Type
	for _, v := range a {
		if t == nil {
			t = reflect.TypeOf(v)
		}
		if reflect.ValueOf(v).Kind() != t.Kind() {
			return false
		}
	}
	return true
}

// Flatten a nested array or slice to at most 'depth' levels. Use depth of -1
// to completely flatten the input.
// Returns a new slice without modifying the input.
func Flatten(list interface{}, depth int) ([]interface{}, error) {
	l, err := common.InterfaceSlice(list)
	if err != nil {
		return nil, err
	}
	if depth == 0 {
		return l, nil
	}
	out := make([]interface{}, 0, len(l)*2)
	for _, v := range l {
		s := reflect.ValueOf(v)
		kind := s.Kind()
		switch kind {
		case reflect.Slice, reflect.Array:
			vl, err := Flatten(v, depth-1)
			if err != nil {
				return nil, err
			}
			out = append(out, vl...)
		default:
			out = append(out, v)
		}
	}
	return out, nil
}

func IsFirst(index int, slice interface{}) bool {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return false
	}

	return index == 0
}

func IsLast(index int, slice interface{}) bool {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return false
	}

	return index == val.Len()-1
}

func HasElements(slice interface{}) bool {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return false
	}

	return val.Len() > 0
}

func Count(slice interface{}) int {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return 0
	}

	return val.Len()
}
func First(slice interface{}) interface{} {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	return val.Index(0).Interface()
}

func Last(slice interface{}) interface{} {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	return val.Index(val.Len() - 1).Interface()
}

func Contains(slice interface{}, key string, value interface{}) bool {
	val := reflect.ValueOf(slice)
	if val.Type().Kind() != reflect.Slice {
		return false
	}

	valueType := reflect.TypeOf(value)
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)
		if item.Type().Kind() == reflect.Struct && key != "" {
			fieldValue := item.FieldByName(key)
			if !fieldValue.IsValid() {
				continue
			}
			if fieldValue.Type().Comparable() && value != nil && valueType.Comparable() {
				if fieldValue.Interface() == value {
					return true
				}
			}
		} else {
			if item.CanInterface() && item.Interface() == value {
				return true
			}
		}
	}
	return false
}

func Find(slice interface{}, key string, value interface{}) interface{} {
	val := reflect.ValueOf(slice)
	if val.Type().Kind() != reflect.Slice {
		return nil
	}

	valueType := reflect.TypeOf(value)
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)
		if item.Type().Kind() == reflect.Struct && key != "" {
			// Handling for struct types with a specified key
			fieldValue := item.FieldByName(key)
			if !fieldValue.IsValid() || !fieldValue.Type().Comparable() || !valueType.Comparable() {
				continue // Skip if the field is not valid or not comparable
			}
			if fieldValue.Interface() == value {
				return item.Interface()
			}
		} else {
			// Handling for non-struct types or struct without specified key
			if item.CanInterface() {
				itemInterface := item.Interface()
				if itemInterface == value {
					return itemInterface
				}
			}
		}
	}
	return nil
}

func Filter(data interface{}, path string, value string) []interface{} {
	var results []interface{}

	keys := strings.Split(path, ".")
	keys = removeEmpty(keys)

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(data)
		for i := 0; i < slice.Len(); i++ {
			item := slice.Index(i).Interface()
			if hasKeyValue(item, keys, value) {
				results = append(results, item)
			}
		}
	case reflect.Struct:
		if hasKeyValue(data, keys, value) {
			results = append(results, data)
		}
	}

	return results
}

func hasKeyValue(data interface{}, keys []string, value string) bool {
	current := reflect.ValueOf(data)
	for _, key := range keys {
		if key != "" {
			if current.Kind() == reflect.Ptr {
				current = current.Elem()
			}
			current = current.FieldByName(key)
			if !current.IsValid() {
				return false
			}
		}
	}
	return current.String() == value
}
func removeEmpty(s []string) []string {
	var result []string
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
