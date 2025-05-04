package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument{
		objectresource.NewMethodArgument("", "foo"),
		objectresource.NewMethodArgument("", "bar"),
	}, objectresource.MethodIdentifier{})

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument{
		objectresource.NewMethodArgument("", "foo"),
		objectresource.NewMethodArgument("", "bar"),
	}, objectresource.MethodIdentifier{})

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument{
		objectresource.NewMethodArgument("foo", "bar"),
	}, objectresource.MethodIdentifier{})

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument{
		objectresource.NewMethodArgument("foo", "bar"),
		objectresource.NewMethodArgument("", "hoo"),
		objectresource.NewMethodArgument("faust", "poe"),
	}, objectresource.MethodIdentifier{})

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument{
		objectresource.NewMethodArgument("foo", "bar"),
	}, objectresource.MethodIdentifier{})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Arguments_Arguments_WithoutValue(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - Name: foo
`

	// Act

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("MethodArgument.Value", fieldRequiredErr.FieldName)
}

func Test_UnMarshall_Arguments_Reference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Reference:
  Name: CMD
`

	// Act

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument(nil), objectresource.NewMethodIdentifier("CMD"))

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

	var data objectresource.Arguments
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewArgument([]objectresource.MethodArgument(nil), objectresource.NewMethodIdentifier("CMD"))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
