package utilities

import "strings"

func IsEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return trimmedText == ""
}
