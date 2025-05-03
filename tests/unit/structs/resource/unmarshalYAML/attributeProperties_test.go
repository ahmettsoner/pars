package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

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

	var data objectresource.AttributeProperties
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAttributeProperties(true, true, true, true, "123", "ddd")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
