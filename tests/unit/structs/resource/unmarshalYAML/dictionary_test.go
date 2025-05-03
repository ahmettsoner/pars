package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

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

	var data objectresource.Dictionary
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDictionary("username_summary", map[string]string(nil))

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

	var data objectresource.Dictionary
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDictionary("username_summary", map[string]string{
		"tr": "Turkish",
		"en": "English",
		"de": "Deutsche",
	})
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Dictionary_Key(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Translates:
  tr: Turkish
  en: English
  de: Deutsche
`

	// Act

	var data objectresource.Dictionary
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Dictionary.Key", fieldRequiredErr.FieldName)
}
