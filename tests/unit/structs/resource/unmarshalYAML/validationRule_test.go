package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_ValidationRule_TypeOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
regex
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRule("regex", "", objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRule_WithName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: regex
Name: "test"
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRule("regex", "test", objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRule_Message_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: regex
Message: message_text
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRule("regex", "", objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRule_Message_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: regex
Message: 
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRule("regex", "", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRule_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: regex
Name: test
Message:
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRule("regex", "test", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRule_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: test
`

	// Act

	var data objectresource.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("ValidationRule.Type", fieldRequiredErr.FieldName)
}
