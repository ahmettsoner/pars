package errors

import "fmt"

type UnkownValueTypeError struct {
	Value string
}

func (e *UnkownValueTypeError) Error() string {
	return fmt.Sprintf("unknown value type format: %s", e.Value)
}
