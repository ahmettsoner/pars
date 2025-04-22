package errors

import "fmt"

type InvalidProjectFullnameFormatError struct {
	Value string
}

func (e *InvalidProjectFullnameFormatError) Error() string {
	return fmt.Sprintf("invalid project fullname format, name should be 'group/name' or 'name' format: %s", e.Value)
}
