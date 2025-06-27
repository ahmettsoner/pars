package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_WorkspaceSpecification_Path_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Path: sample-path
`

	// Act

	var data workspace.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := workspace.NewWorkspaceSpecification(0, "CMD", "sample-path")

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

	var data workspace.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := workspace.NewWorkspaceSpecification(0, "CMD", "sample-path")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
