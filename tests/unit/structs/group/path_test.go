package group

import (
	"path/filepath"
	"testing"

	applicationGroup "parsdevkit.net/application/structs/group"

	"parsdevkit.net/modules/group/group_payload"

	"github.com/stretchr/testify/assert"
)

func Test_Group_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := group_payload.GroupSpecification{
		GroupIdentifier: applicationGroup.GroupIdentifier{
			Path: "path",
		},
	}
	// Act
	expected := filepath.Join("path")

	// Assert
	a.Equal(expected, data.GetRelativeGroupPath())
}
