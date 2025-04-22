package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

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

	var data objectresource.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMessage("CMD", objectresource.DictionaryIdentifier{})

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

	var data objectresource.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMessage("CMD", objectresource.DictionaryIdentifier{})

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

	var data objectresource.Message
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("dictionary"))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
