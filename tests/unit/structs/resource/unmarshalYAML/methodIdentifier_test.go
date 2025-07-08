package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_MethodIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data object_resource_payload_structs.MethodIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_MethodIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data object_resource_payload_structs.MethodIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
