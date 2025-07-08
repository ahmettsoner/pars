package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_AttributeProperties_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: true
Required: true
ReadOnly: true
Unique: true
Default: 123
Format: ddd
`

	// Act

	var data object_resource_payload_structs.AttributeProperties
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttributeProperties(true, true, true, true, "123", "ddd")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
