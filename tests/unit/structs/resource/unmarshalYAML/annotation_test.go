package unmarshalYAML

import (
	"testing"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Annotation_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: auth.annotation
`

	// Act

	var data object_resource_payload_structs.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Annotation_Arguments_InlineValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: auth.annotation
Arguments:
  - foo
  - bar
`

	// Act

	var data object_resource_payload_structs.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("", "foo"),
		object_resource_payload_structs.NewMethodArgument("", "bar"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Annotation_Arguments(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: auth.annotation
Arguments:
  - Name: param1
    Value: foo
  - Name: param2
    Value: bar
`

	// Act

	var data object_resource_payload_structs.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("param1", "foo"),
		object_resource_payload_structs.NewMethodArgument("param2", "bar"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Annotation_Arguments_OnlyValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: auth.annotation
Arguments:
  - Value: foo
  - Value: bar
`

	// Act

	var data object_resource_payload_structs.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument{
		object_resource_payload_structs.NewMethodArgument("", "foo"),
		object_resource_payload_structs.NewMethodArgument("", "bar"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
