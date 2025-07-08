package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

func Test_UnMarshall_WorkspaceSpecification_Path_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Path: sample-path
`

	// Act

	var data basic_workspace_payload_structs.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := basic_workspace_payload_structs.NewWorkspaceSpecification(0, "CMD", "sample-path")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_WorkspaceSpecification_ID_ShouldBeZero(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Id: 100
Name: CMD
Path: sample-path
`

	// Act

	var data basic_workspace_payload_structs.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := basic_workspace_payload_structs.NewWorkspaceSpecification(0, "CMD", "sample-path")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
