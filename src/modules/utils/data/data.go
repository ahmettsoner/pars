package data

import "reflect"

func Default(value any, defaultValue any) any {
	if value == nil {
		return defaultValue
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if val.Len() == 0 {
			return defaultValue
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() == 0 {
			return defaultValue
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() == 0 {
			return defaultValue
		}
	case reflect.Bool:
		if !val.Bool() {
			return defaultValue
		}
	}

	return value
}

func IsDefault(value any) bool {
	if value == nil {
		return true
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if val.Len() == 0 {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() == 0 {
			return true
		}
	case reflect.Float32, reflect.Float64:
		if val.Float() == 0 {
			return true
		}
	case reflect.Bool:
		if !val.Bool() {
			return true
		}
	}

	return false
}
