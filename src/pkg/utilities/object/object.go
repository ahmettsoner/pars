package object

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

func PrintFields2(obj interface{}) {
	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Println("YAML convert error:", err)
		return
	}
	fmt.Println(string(yamlData))
}
func PrintFields(obj interface{}, depth int) {
	val := reflect.ValueOf(obj)

	switch val.Kind() {
	case reflect.Struct:
		typ := val.Type()
		fmt.Printf("%v\n", typ)
		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldValue := val.Field(i).Interface()

			fmt.Printf("%s%s: ", strings.Repeat("   ", depth), fieldName)

			if reflect.ValueOf(fieldValue).Kind() == reflect.Struct || reflect.ValueOf(fieldValue).Kind() == reflect.Array || reflect.ValueOf(fieldValue).Kind() == reflect.Slice {
				PrintFields(fieldValue, depth+1)
			} else {
				fmt.Printf("%v\n", fieldValue)
			}
		}
	case reflect.Array, reflect.Slice:
		fmt.Printf("%s- ", strings.Repeat("   ", depth))
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()
			if reflect.ValueOf(elem).Kind() == reflect.Struct || reflect.ValueOf(elem).Kind() == reflect.Array || reflect.ValueOf(elem).Kind() == reflect.Slice {
				PrintFields(elem, depth+1)
			} else {
				fmt.Printf("%v", elem)
			}
			if i < val.Len()-1 {
				fmt.Printf("\n%s- ", strings.Repeat("   ", depth))
			}
		}
		fmt.Println()
	default:
		fmt.Println("Error: this explains only structs and arrays/slices.")
	}
}
func IsZeroValue(v any) bool {
	return reflect.ValueOf(v).IsZero()
}

func EmptyValue(v any) string {
	rv := reflect.ValueOf(v)

	if !rv.IsValid() {
		return "(none)"
	}

	switch rv.Kind() {
	case reflect.String:
		return "N/A"
	case reflect.Slice, reflect.Array:
		return "(empty)"
	case reflect.Map:
		return "(empty)"
	case reflect.Ptr:
		if rv.IsNil() {
			return "(none)"
		}
		return EmptyValue(rv.Elem().Interface())
	case reflect.Interface:
		if rv.IsNil() {
			return "(none)"
		}
		return EmptyValue(rv.Elem().Interface())
	default:
		return "(none)"
	}
}

func IsEmpty(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)

	// Handle nil interface or pointer
	if !rv.IsValid() || (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && rv.IsNil() {
		return true
	}

	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len() == 0
	case reflect.Struct:
		// Optional: could also check for all zero fields
		return reflect.DeepEqual(v, reflect.Zero(rv.Type()).Interface())
	default:
		// Fallback for other types
		return reflect.DeepEqual(v, reflect.Zero(rv.Type()).Interface())
	}
}
