package errors

import "fmt"

type InvalidFormatForRuntimeError struct {
	Value string
}

func (e *InvalidFormatForRuntimeError) Error() string {
	return fmt.Sprintf("Invalid format for runtime: %s", e.Value)
}
