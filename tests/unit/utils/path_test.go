package group

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	applicationGroup "parsdevkit.net/application/structs/group"
	group "parsdevkit.net/modules/group/group"
)

func Test_Group_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := group.GroupSpecification{
		GroupIdentifier: applicationGroup.GroupIdentifier{
			Path: "path",
		},
	}
	// Act
	expected := filepath.Join("path")

	// Assert
	a.Equal(expected, data.GetRelativeGroupPath())
}
