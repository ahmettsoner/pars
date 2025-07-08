package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Arguments_Arguments_ValueOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - Value: foo
  - Value: bar
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("", "foo"),
		object_resource_payload_structs.NewMethodArgument("", "bar"),
	}, object_resource_payload_structs.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Arguments_Arguments_ValueOnly_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - foo
  - bar
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("", "foo"),
		object_resource_payload_structs.NewMethodArgument("", "bar"),
	}, object_resource_payload_structs.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Arguments_Arguments_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - foo bar
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("foo", "bar"),
	}, object_resource_payload_structs.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Arguments_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo bar, hoo, faust poe
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("foo", "bar"),
		object_resource_payload_structs.NewMethodArgument("", "hoo"),
		object_resource_payload_structs.NewMethodArgument("faust", "poe"),
	}, object_resource_payload_structs.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Arguments_Arguments_WithName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - Name: foo
    Value: bar
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("foo", "bar"),
	}, object_resource_payload_structs.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Arguments_Reference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Reference:
  Name: CMD
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument(nil), object_resource_payload_structs.NewMethodIdentifier("CMD"))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Arguments_Reference_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Reference: CMD
`

	// Act

	var data object_resource_payload_structs.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewArgument([]object_resource_payload_structs.MethodArgument(nil), object_resource_payload_structs.NewMethodIdentifier("CMD"))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
