package errors

import "fmt"

type ErrFieldRequired struct {
	FieldName string
}

func (e *ErrFieldRequired) Error() string {
	return fmt.Sprintf("field '%s' is required", e.FieldName)
}
