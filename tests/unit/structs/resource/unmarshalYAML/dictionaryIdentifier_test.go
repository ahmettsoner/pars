package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_DictionaryIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: CMD
`

	// Act

	var data objectresource.DictionaryIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDictionaryIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_DictionaryIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data objectresource.DictionaryIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDictionaryIdentifier("CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
