package errors

import "fmt"

type BaseError struct {
	Code        string
	Message     string
	Description string
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("[%s] %s - %s", e.Code, e.Message, e.Description)
}
