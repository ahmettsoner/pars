package workspace

import (
	"testing"

	"github.com/stretchr/testify/require"
	"parsdevkit.net/core/utils"
)

func Test_GetCodeBaseLocation(t *testing.T) {

	// Arrange
	codebaseLocation := utils.GetCodeBaseLocation()

	// Act

	// Assert
	require.NotEmpty(t, codebaseLocation, "not found codebase location")
}
