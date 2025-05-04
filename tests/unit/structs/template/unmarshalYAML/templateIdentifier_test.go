package unmarshalYAML

import (
	"testing"

	codetemplate "parsdevkit.net/structs/template/code-template"

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

	var data codetemplate.TemplateIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateIdentifier(0, "CMD", "")

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

	var data codetemplate.TemplateIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
