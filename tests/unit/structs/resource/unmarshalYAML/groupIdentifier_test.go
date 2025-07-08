package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_GroupIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data object_resource_payload_structs.GroupIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewGroupIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data object_resource_payload_structs.GroupIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewGroupIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
