package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

// 	var data object_resource_payload_structs.ValidationRegexRule
// 	err := yaml.Unmarshal([]byte(yamlData), &data)

// 	expected := object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("test", "CMD", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")))

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRegexRules_patient_filter")))

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

	var data object_resource_payload_structs.ValidationRegexRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRegexRule("test", "CMD", object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRegexRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
