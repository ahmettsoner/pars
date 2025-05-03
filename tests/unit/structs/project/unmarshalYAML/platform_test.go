package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"

	parsModels "parsdevkit.net/platforms/pars/models"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Platform_TypeOnly_Lowercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: pars
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform_Basic(models.PlatformTypes.Pars)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Platform_TypeOnly_Uppercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: PARS
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform_Basic(models.PlatformTypes.Pars)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Platform_TypeOnly_Camelcase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Pars
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform_Basic(models.PlatformTypes.Pars)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Platform_InvalidType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: unknown_type
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
}
func Test_UnMarshall_Platform_TypeOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
pars
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform_Basic(models.PlatformTypes.Pars)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Platform_Inline_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
pars@BetaV1
`

	// Act
	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform(models.PlatformTypes.Pars, parsModels.ParsPlatformVersions.BetaV1.String())

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Platform_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Pars
Version: BetaV1
`

	// Act

	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewPlatform(models.PlatformTypes.Pars, parsModels.ParsPlatformVersions.BetaV1.String())

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Platform_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Version: BetaV1
`

	// Act

	var data applicationproject.Platform
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Type", fieldRequiredErr.FieldName)
}
