package common

import (
	"strings"
	"testing"

	"parsdevkit.net/pkg/utilities/file"
)

func GenerateEnvironment(t *testing.T, path string) string {
	return strings.Join(file.PathToArray(path), "-")
}
