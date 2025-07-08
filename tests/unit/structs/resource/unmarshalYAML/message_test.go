package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Message_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Text: CMD
`

	// Act

	var data object_resource_payload_structs.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMessage("CMD", object_resource_payload_structs.DictionaryIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Message_ByTextInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data object_resource_payload_structs.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMessage("CMD", object_resource_payload_structs.DictionaryIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Message_ByReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefMessage: dictionary
`

	// Act

	var data object_resource_payload_structs.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("dictionary"))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
