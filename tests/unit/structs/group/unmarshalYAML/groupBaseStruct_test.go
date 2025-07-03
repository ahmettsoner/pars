package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/modules/group/group_payload"

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

	var data group_payload.GroupBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.GroupBaseStruct{
		Header: schemas.SchemaHeader{
			Type: schemas.StructTypes.Group,
			Kind: "",
			Name: "Pars.CMD",
			Metadata: schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		},
		Specifications: group_payload.NewGroupSpecification(0, "CMD", "cmd", []string{"pars", "cmd"}),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
