package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/group"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/schemas"
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

	var data group.GroupBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group.GroupBaseStruct{
		Header: schemas.SchemaHeader{
			Type: schemas.StructTypes.Group,
			Kind: "",
			Name: "Pars.CMD",
			Metadata: schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		},
		Specifications: group.NewGroupSpecification(0, "CMD", "cmd", []string{"pars", "cmd"}),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
