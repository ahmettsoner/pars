package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_WorkspaceIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data workspace.WorkspaceIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := workspace.NewWorkspaceIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_WorkspaceIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data workspace.WorkspaceIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := workspace.NewWorkspaceIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
