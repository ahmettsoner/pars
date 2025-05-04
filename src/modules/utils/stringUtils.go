package utils

import "strings"

func IsEmpty(text string) bool {
	trimmedText := strings.TrimSpace(text)
	return trimmedText == ""
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
