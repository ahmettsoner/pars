package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

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

	var data objectresource.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument(nil))

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

	var data objectresource.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument{
		objectresource.NewMethodArgument("", "foo"),
		objectresource.NewMethodArgument("", "bar"),
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

	var data objectresource.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument{
		objectresource.NewMethodArgument("param1", "foo"),
		objectresource.NewMethodArgument("param2", "bar"),
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

	var data objectresource.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument{
		objectresource.NewMethodArgument("", "foo"),
		objectresource.NewMethodArgument("", "bar"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Annotation_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Arguments:
  - Name: param1
    Value: foo
  - Name: param2
    Value: bar
`

	// Act

	var data objectresource.Annotation
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Annotation.Type", fieldRequiredErr.FieldName)
}
