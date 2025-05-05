package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Validation_RegexType_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
- Regex: CMD
`

	// Act

	var data objectresource.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidation(
		objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Validation_RegexType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
- Type: Regex
  Pattern: CMD
`

	// Act

	var data objectresource.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidation(
		objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Validation_MultipleType_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
- Regex: CMD
- Length: 10
`

	// Act

	var data objectresource.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidation(
		objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{}),
		objectresource.NewValidationLengthRule("", 10, 0, objectresource.Message{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Validation_MultipleType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
- Type: Regex
  Pattern: CMD
- Type: Length
  Min: 10
`

	// Act

	var data objectresource.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidation(
		objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{}),
		objectresource.NewValidationLengthRule("", 10, 0, objectresource.Message{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
