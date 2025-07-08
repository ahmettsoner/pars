package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

	var data object_resource_payload_structs.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidation(
		object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{}),
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

	var data object_resource_payload_structs.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidation(
		object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{}),
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

	var data object_resource_payload_structs.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidation(
		object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{}),
		object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.Message{}),
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

	var data object_resource_payload_structs.Validation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewValidation(
		object_resource_payload_structs.NewValidationRegexRule("", "CMD", object_resource_payload_structs.Message{}),
		object_resource_payload_structs.NewValidationLengthRule("", 10, 0, object_resource_payload_structs.Message{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
