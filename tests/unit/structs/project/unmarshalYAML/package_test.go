package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	applicationProject "parsdevkit.net/application/structs/project"
)

func Test_UnMarshall_Dependency_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act
	var data applicationProject.Dependency
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewDependency_Basic("foo")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Dependency_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act
	var data applicationProject.Dependency
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewDependency_Basic("foo")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Dependency_Inline_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo@bar
`

	// Act
	var data applicationProject.Dependency
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewDependency("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Dependency_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Version: bar
`

	// Act

	var data applicationProject.Dependency
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewDependency("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Dependency_FullName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Version: bar
`

	// Act

	var data applicationProject.Dependency
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewDependency("foo", "bar")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
	a.Equal(expected.GetFullName(), "foo@bar")
}
