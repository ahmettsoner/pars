package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

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
	var data objectresource.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{}, true)

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
	var data objectresource.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{Name: "GetName"}, true)

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

	var data objectresource.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationSetter("", objectresource.VisibilityTypeTypes.Protected, objectresource.MethodIdentifier{}, true)

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

	var data objectresource.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{Name: "setterMethod"}, true)

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

	var data objectresource.EncapsulationSetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationSetter("", "", objectresource.MethodIdentifier{}, false)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
