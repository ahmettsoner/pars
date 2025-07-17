package group

import (
	"path/filepath"
	"testing"

	applicationGroup "parsdevkit.net/application/structs/group"

	"github.com/stretchr/testify/assert"
)

func Test_Group_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := applicationGroup.GroupSpecification{
		GroupIdentifier: applicationGroup.GroupIdentifier{
			Path: "path",
		},
	}
	// Act
	expected := filepath.Join("path")

	// Assert
	a.Equal(expected, data.GetRelativeGroupPath())
}
