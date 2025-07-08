package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.Message{})

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

// 	var data object_resource_payload_structs.ValidationLengthRule
// 	err := yaml.Unmarshal([]byte(yamlData), &data)

// 	expected := object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 0, 10, object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 0, 10, object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("test", 10, 0, object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")))

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationLengthRules_patient_filter")))

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

	var data object_resource_payload_structs.ValidationLengthRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationLengthRule("test", 10, 50, object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationLengthRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
