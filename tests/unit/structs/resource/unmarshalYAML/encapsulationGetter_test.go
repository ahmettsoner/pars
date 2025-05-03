package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

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
	var data objectresource.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{}, true)

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
	var data objectresource.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{Name: "GetName"}, true)

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

	var data objectresource.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationGetter("", objectresource.VisibilityTypeTypes.Protected, objectresource.MethodIdentifier{}, true)

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

	var data objectresource.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{Name: "getterMethod"}, true)

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

	var data objectresource.EncapsulationGetter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewEncapsulationGetter("", "", objectresource.MethodIdentifier{}, false)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
