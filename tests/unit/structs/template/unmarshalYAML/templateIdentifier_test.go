package unmarshalYAML

import (
	"testing"

	applicationTemplate "parsdevkit.net/application/structs/template"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_TemplateIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data applicationTemplate.TemplateIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data applicationTemplate.TemplateIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
