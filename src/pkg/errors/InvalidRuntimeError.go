package errors

import "fmt"

type InvalidRuntimeError struct {
	Value string
}

func (e *InvalidRuntimeError) Error() string {
	return fmt.Sprintf("invalid runtime format: %s", e.Value)
}
