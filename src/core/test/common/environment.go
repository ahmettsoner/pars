package common

import (
	"strings"
	"testing"

	"parsdevkit.net/core/utils"
)

func GenerateEnvironment(t *testing.T, path string) string {
	return strings.Join(utils.PathToArray(path), "-")
}
