package workspace

import (
	"testing"

	"parsdevkit.net/application"

	"github.com/stretchr/testify/require"
)

func Test_GetCodeBaseLocation(t *testing.T) {

	// Arrange
	codebaseLocation := application.GetCodeBaseLocation()

	// Act

	// Assert
	require.NotEmpty(t, codebaseLocation, "not found codebase location")
}
