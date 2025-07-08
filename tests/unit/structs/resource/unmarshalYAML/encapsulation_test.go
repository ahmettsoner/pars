package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Encapsulation_Getter_BooleanDefinition_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Getter: true
`

	// Act
	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{}, true), object_resource_payload_structs.EncapsulationSetter{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Getter_MethodReference_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Getter: GetName
`

	// Act
	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "GetName"}, true), object_resource_payload_structs.EncapsulationSetter{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Getter_Visibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Getter:
  Visibility: protected
`

	// Act

	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.NewEncapsulationGetter("", structs.VisibilityTypeTypes.Protected, object_resource_payload_structs.MethodIdentifier{}, true), object_resource_payload_structs.EncapsulationSetter{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Getter_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Getter:
  RefMethod: GetName
`

	// Act

	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "GetName"}, true), object_resource_payload_structs.EncapsulationSetter{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

//************************************************************************************************

func Test_UnMarshall_Encapsulation_Setter_BooleanDefinition_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Setter: true
`

	// Act
	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.EncapsulationGetter{}, object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{}, true))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Setter_MethodReference_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Setter: SetName
`

	// Act
	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.EncapsulationGetter{}, object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "SetName"}, true))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Setter_Visibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Setter:
  Visibility: protected
`

	// Act

	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.EncapsulationGetter{}, object_resource_payload_structs.NewEncapsulationSetter("", structs.VisibilityTypeTypes.Protected, object_resource_payload_structs.MethodIdentifier{}, true))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Encapsulation_Setter_MethodReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Setter:
  RefMethod: SetName
`

	// Act

	var data object_resource_payload_structs.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewEncapsulation(object_resource_payload_structs.EncapsulationGetter{}, object_resource_payload_structs.NewEncapsulationSetter("", "", object_resource_payload_structs.MethodIdentifier{Name: "SetName"}, true))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
