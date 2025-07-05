package errors

import "fmt"

type InvalidFormatForDependencyError struct {
	Value string
}

func (e *InvalidFormatForDependencyError) Error() string {
	return fmt.Sprintf("Invalid format for package: %s", e.Value)
}
