package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"

	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Runtime_TypeOnly_Lowercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: dotnet
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Runtime_TypeOnly_Uppercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: DOTNET
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Runtime_TypeOnly_Camelcase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Dotnet
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Runtime_InvalidType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: unknown_type
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
}
func Test_UnMarshall_Runtime_TypeOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
dotnet
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, "")
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Runtime_Inline_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
dotnet@Net8
`

	// Act
	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, dotnetModels.DotnetRuntimeVersions.Net8.String())

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Runtime_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Dotnet
Version: Net8
`

	// Act

	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewRuntime(models.RuntimeTypes.Dotnet, dotnetModels.DotnetRuntimeVersions.Net8.String())

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Runtime_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Version: Net8
`

	// Act

	var data applicationproject.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Type", fieldRequiredErr.FieldName)
}
