package utils

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
