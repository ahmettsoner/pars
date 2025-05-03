package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// // Object yapısı çözümleme alt yapısı hazırlanması
// func Test_UnMarshall_ValidationRegexRule_PatternOnly_AsObject(t *testing.T) {

// 	// Arrange
// 	a := assert.New(t)
// 	yamlData := `
// Regex:
//   Pattern: CMD
// `

// 	// Act

// 	var data objectresource.ValidationRegexRule
// 	err := yaml.Unmarshal([]byte(yamlData), &data)

// 	expected := objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{})

// 	// Assert
// 	a.NoError(err)
// 	a.Equal(expected, data)
// }

func Test_UnMarshall_ValidationRegexRule_PatternOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Pattern: CMD
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_PatternOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Regex: CMD
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("", "CMD", objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_WithName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Name: "test"
Pattern: CMD
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("test", "CMD", objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_Message_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Pattern: CMD
Message: message_text
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("", "CMD", objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_Message_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Pattern: CMD
Message: 
  RefMessage: validationRegexRules_patient_filter
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("", "CMD", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRegexRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Name: test
Pattern: CMD
Message:
  RefMessage: validationRegexRules_patient_filter
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationRegexRule("test", "CMD", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRegexRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationRegexRule_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Pattern: CMD
Message:
  RefMessage: validationRegexRules_patient_filter
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("ValidationRule.Type", fieldRequiredErr.FieldName)
}

func Test_UnMarshall_ValidationRegexRule_WithoutPattern(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Regex
Message:
  RefMessage: validationRegexRules_patient_filter
`

	// Act

	var data objectresource.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("ValidationRegexRule.Pattern", fieldRequiredErr.FieldName)
}
