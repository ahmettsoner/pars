package errors

import "fmt"

type InvalidPlatformError struct {
	Value string
}

func (e *InvalidPlatformError) Error() string {
	return fmt.Sprintf("invalid platform format: %s", e.Value)
}
