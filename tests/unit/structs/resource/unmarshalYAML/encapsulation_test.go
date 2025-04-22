package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

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
	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{}, true), objectresource.EncapsulationSetter{})

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
	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{Name: "GetName"}, true), objectresource.EncapsulationSetter{})

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

	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.NewEncapsulationGetter("", objectresource.VisibilityTypeTypes.Protected, objectresource.MethodIdentifier{}, true), objectresource.EncapsulationSetter{})

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

	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{Name: "GetName"}, true), objectresource.EncapsulationSetter{})

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
	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.EncapsulationGetter{}, objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{}, true))

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
	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.EncapsulationGetter{}, objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{Name: "SetName"}, true))

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

	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.EncapsulationGetter{}, objectresource.NewEncapsulationSetter("", objectresource.VisibilityTypeTypes.Protected, objectresource.MethodIdentifier{}, true))

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

	var data objectresource.Encapsulation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulation(objectresource.EncapsulationGetter{}, objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{Name: "SetName"}, true))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
