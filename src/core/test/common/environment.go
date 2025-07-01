package common

import (
	"strings"
	"testing"

	"parsdevkit.net/core/utilities"
)

func GenerateEnvironment(t *testing.T, path string) string {
	return strings.Join(utilities.PathToArray(path), "-")
}
