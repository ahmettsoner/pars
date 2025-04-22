package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_WorkspaceBaseStruct_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Workspace
Name:  Pars.CMD
Metadata:
  Tags: tag1, tag2
Specifications:
  Name: CMD
  Package: pars/cmd
  Path: cmd
`

	// Act

	var data workspace.WorkspaceBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := workspace.WorkspaceBaseStruct{
		Header: structs.Header{
			Type: structs.StructTypes.Workspace,
			Name: "Pars.CMD",
			Metadata: structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		},
		Specifications: workspace.NewWorkspaceSpecification(0, "CMD", "cmd"),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
