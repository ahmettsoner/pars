package errors

import "fmt"

type InvalidDependencyError struct {
	Value string
}

func (e *InvalidDependencyError) Error() string {
	return fmt.Sprintf("invalid package format: %s", e.Value)
}
