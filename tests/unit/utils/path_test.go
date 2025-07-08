package group

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	applicationGroup "parsdevkit.net/application/structs/group"

	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

func Test_Group_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := basic_group_payload_structs.GroupSpecification{
		GroupIdentifier: applicationGroup.GroupIdentifier{
			Path: "path",
		},
	}
	// Act
	expected := filepath.Join("path")

	// Assert
	a.Equal(expected, data.GetRelativeGroupPath())
}
