package errors

import "fmt"

type InvalidFormatForLanguageError struct {
	Value string
}

func (e *InvalidFormatForLanguageError) Error() string {
	return fmt.Sprintf("Invalid format for language: %s", e.Value)
}
