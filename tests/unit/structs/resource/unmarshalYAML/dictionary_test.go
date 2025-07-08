package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Dictionary_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: username_summary
`

	// Act

	var data object_resource_payload_structs.Dictionary
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewDictionary("username_summary", map[string]string(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Dictionary_Translates(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: username_summary
Translates:
  tr: Turkish
  en: English
  de: Deutsche
`

	// Act

	var data object_resource_payload_structs.Dictionary
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewDictionary("username_summary", map[string]string{
		"tr": "Turkish",
		"en": "English",
		"de": "Deutsche",
	})
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
