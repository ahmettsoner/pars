package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/errors"

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

func Test_UnMarshall_WorkspaceSpecification_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Path: CMD
`

	// Act

	var data workspace.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Name", fieldRequiredErr.FieldName)
}

func Test_UnMarshall_WorkspaceSpecification_WithoutPath(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data workspace.WorkspaceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Path", fieldRequiredErr.FieldName)
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
