package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Option_OnlyKey_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act

	var data option.Option
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := option.NewOption("foo", nil)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Option_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo=bar
`

	// Act

	var data option.Option
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := option.NewOption("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Option_OnlyKey(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: foo
`

	// Act

	var data option.Option
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := option.NewOption("foo", nil)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Option_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Key: foo
Value: bar
`

	// Act

	var data option.Option
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := option.NewOption("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Option_WithoutKey(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Value: bar
`

	// Act

	var data option.Option
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Option.Key", fieldRequiredErr.FieldName)
}
