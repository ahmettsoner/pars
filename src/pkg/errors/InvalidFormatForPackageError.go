package errors

import "fmt"

type InvalidFormatForPackageError struct {
	Value string
}

func (e *InvalidFormatForPackageError) Error() string {
	return fmt.Sprintf("Invalid format for package: %s", e.Value)
}
