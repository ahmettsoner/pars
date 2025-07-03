package errors

import "fmt"

type InvalidPackageError struct {
	Value string
}

func (e *InvalidPackageError) Error() string {
	return fmt.Sprintf("invalid package format: %s", e.Value)
}
