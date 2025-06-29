package unmarshalYAML

import (
	"testing"

	applicationResource "parsdevkit.net/application/structs/resource"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_ResourceIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data applicationResource.ResourceIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data applicationResource.ResourceIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
