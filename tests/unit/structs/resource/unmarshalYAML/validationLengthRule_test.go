package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_ValidationLengthRule_MinOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Min: 10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 10, 0, objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_MinOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Length: 10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 10, 0, objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

// // // Burda çözüm üretilebilir, yaml syntax benzerliğinden kullanılamıyor
// func Test_UnMarshall_ValidationLengthRule_MinOnlyInline_WithSeparator(t *testing.T) {

// 	// Arrange
// 	a := assert.New(t)
// 	yamlData := `
// Length: 10:
// `

// 	// Act

// 	var data objectresource.ValidationLengthRule
// 	err := yaml.Unmarshal([]byte(yamlData), &data)

// 	expected := objectresource.NewValidationLengthRule("", 10, 0, objectresource.Message{})

// 	// Assert
// 	a.NoError(err)
// 	a.Equal(expected, data)
// }

func Test_UnMarshall_ValidationLengthRule_MaxOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Max: 10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 0, 10, objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_MaxOnlyInline_WithSeparator(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Length: :10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 0, 10, objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_WithName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Name: test
Min: 10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("test", 10, 0, objectresource.Message{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_Message_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Min: 10
Message: message_text
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 10, 0, objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_Message_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Min: 10
Message:
  RefMessage: validationLengthRules_patient_filter
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("", 10, 0, objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationLengthRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Length
Name: test
Min: 10
Max: 50
Message:
  RefMessage: validationLengthRules_patient_filter
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewValidationLengthRule("test", 10, 50, objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationLengthRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ValidationLengthRule_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Min: 10
`

	// Act

	var data objectresource.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("ValidationRule.Type", fieldRequiredErr.FieldName)
}
