package unmarshalYAML

import (
	"testing"

	applicationGroup "parsdevkit.net/application/structs/group"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
)

func Test_UnMarshall_GroupBaseStruct_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Group
Name:  Pars.CMD
Metadata:
  Tags: tag1, tag2
Specifications:
  Name: CMD
  Package: pars/cmd
  Path: cmd
`

	// Act

	var data basic_group_payload_structs.GroupBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := basic_group_payload_structs.GroupBaseStruct{
		Header: schemas.SchemaHeader{
			Type: schemas.StructTypes.Group,
			Kind: "",
			Name: "Pars.CMD",
			Metadata: schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		},
		Specifications: applicationGroup.NewGroupSpecification(0, "CMD", "cmd", []string{"pars", "cmd"}),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
