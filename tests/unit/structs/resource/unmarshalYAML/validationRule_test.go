package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

	var data object_resource_payload_structs.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRule("regex", "", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRule("regex", "test", object_resource_payload_structs.Message{})

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

	var data object_resource_payload_structs.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRule("regex", "", object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")))

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

	var data object_resource_payload_structs.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRule("regex", "", object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")))

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

	var data object_resource_payload_structs.ValidationRule
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidationRule("regex", "test", object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
