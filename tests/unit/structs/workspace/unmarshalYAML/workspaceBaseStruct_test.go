package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
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

	var data basic_workspace_payload_structs.WorkspaceBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := basic_workspace_payload_structs.WorkspaceBaseStruct{
		Header: schemas.SchemaHeader{
			Type: schemas.StructTypes.Workspace,
			Name: "Pars.CMD",
			Metadata: schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		},
		Specifications: basic_workspace_payload_structs.NewWorkspaceSpecification(0, "CMD", "cmd"),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
