package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_MethodArgument_ValueOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Value: bar
`

	// Act
	var data object_resource_payload_structs.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodArgument("", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodArgument_InlineValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
bar
`

	// Act
	var data object_resource_payload_structs.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodArgument("", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodArgument_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo bar
`

	// Act
	var data object_resource_payload_structs.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodArgument("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodArgument_Arguments_WithName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Value: bar
`

	// Act

	var data object_resource_payload_structs.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethodArgument("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
