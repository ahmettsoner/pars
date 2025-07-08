package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_EncapsulationSetter_Inline_BooleanDefinition(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
true
`

	// Act
	var data object_resource_payload_structs.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationSetter_Inline_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
GetName
`

	// Act
	var data object_resource_payload_structs.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "GetName"}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationSetter_Visibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Visibility: protected
`

	// Act

	var data object_resource_payload_structs.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationSetter("", structs.VisibilityTypeTypes.Protected, object_resource_payload_structs.MethodIdentifier{}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationSetter_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefMethod: setterMethod
`

	// Act

	var data object_resource_payload_structs.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "setterMethod"}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationSetter_WithoutAnyValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
`

	// Act

	var data object_resource_payload_structs.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{}, false)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
