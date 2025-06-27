package functions

import (
	"fmt"

	"parsdevkit.net/core/utils/base64"
)

type Base64Funcs struct{}

func (b Base64Funcs) Encode(in []byte) string {
	result, err := base64.Encode(in)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (b Base64Funcs) Decode(in string) []byte {
	result, err := base64.Decode(in)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}
