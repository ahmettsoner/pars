package errors

import "fmt"

type InvalidLanguageError struct {
	Value string
}

func (e *InvalidLanguageError) Error() string {
	return fmt.Sprintf("invalid language format: %s", e.Value)
}
