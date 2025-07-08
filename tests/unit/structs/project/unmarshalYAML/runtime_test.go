package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	dotnetModels "parsdevkit.net/platforms/dotnet/models"

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
	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, "")

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
	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, "")

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
	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, "")

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
	var data application_project_payload_structs.Runtime
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
	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, "")
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
	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, dotnetModels.DotnetRuntimeVersions.Net8.String())

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

	var data application_project_payload_structs.Runtime
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewRuntime(models.RuntimeTypes.Dotnet, dotnetModels.DotnetRuntimeVersions.Net8.String())

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
