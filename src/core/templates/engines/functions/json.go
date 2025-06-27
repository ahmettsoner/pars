package functions

import (
	"fmt"

	"parsdevkit.net/core/utils/json"
)

type JsonFuncs struct{}

func (j JsonFuncs) ToJson(data interface{}) string {
	result, err := json.ToJson(data)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (j JsonFuncs) PrettifyJSON(jsonString string) string {
	result, err := json.PrettifyJSON(jsonString)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}
