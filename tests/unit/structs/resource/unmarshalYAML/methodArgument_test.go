package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

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
	var data objectresource.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodArgument("", "bar")

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
	var data objectresource.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodArgument("", "bar")

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
	var data objectresource.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodArgument("foo", "bar")

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

	var data objectresource.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodArgument("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_MethodArgument_WithoutValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act

	var data objectresource.MethodArgument
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("MethodArgument.Value", fieldRequiredErr.FieldName)
}
