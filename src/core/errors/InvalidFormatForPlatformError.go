package errors

import "fmt"

type InvalidFormatForPlatformError struct {
	Value string
}

func (e *InvalidFormatForPlatformError) Error() string {
	return fmt.Sprintf("Invalid format for platform: %s", e.Value)
}
