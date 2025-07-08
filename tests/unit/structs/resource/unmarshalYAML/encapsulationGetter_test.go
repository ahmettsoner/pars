package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_EncapsulationGetter_Inline_BooleanDefinition(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
true
`

	// Act
	var data object_resource_payload_structs.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationGetter_Inline_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
GetName
`

	// Act
	var data object_resource_payload_structs.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "GetName"}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationGetter_Visibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Visibility: protected
`

	// Act

	var data object_resource_payload_structs.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationGetter("", structs.VisibilityTypeTypes.Protected, object_resource_payload_structs.MethodIdentifier{}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationGetter_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefMethod: getterMethod
`

	// Act

	var data object_resource_payload_structs.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "getterMethod"}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_EncapsulationGetter_WithoutAnyValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
`

	// Act

	var data object_resource_payload_structs.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{}, false)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
