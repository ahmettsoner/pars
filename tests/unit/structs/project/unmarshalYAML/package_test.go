package unmarshalYAML

import (
	"testing"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Package_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act
	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPackage_Basic("foo")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Package_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act
	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPackage_Basic("foo")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Package_Inline_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo@bar
`

	// Act
	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPackage("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Package_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Version: bar
`

	// Act

	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPackage("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Package_FullName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Version: bar
`

	// Act

	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPackage("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
	a.Equal(expected.GetFullName(), "foo@bar")
}

func Test_UnMarshall_Package_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Version: bar
`

	// Act

	var data applicationproject.Package
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Name", fieldRequiredErr.FieldName)
}
