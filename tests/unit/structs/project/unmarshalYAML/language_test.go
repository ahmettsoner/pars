package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Language_TypeOnly_Lowercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: csharp
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage_LanguageOnly(models.LanguageTypes.CSharp)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Language_TypeOnly_Uppercase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: CSHARP
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage_LanguageOnly(models.LanguageTypes.CSharp)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Language_TypeOnly_Camelcase(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: CSharp
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage_LanguageOnly(models.LanguageTypes.CSharp)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Language_InvalidType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: unknown_type
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
}
func Test_UnMarshall_Language_TypeOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
csharp
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage_LanguageOnly(models.LanguageTypes.CSharp)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Language_Inline_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
csharp@V8
`

	// Act
	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage(models.LanguageTypes.CSharp, string(models.CSharpVersions.V8))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Language_WithVersion(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: CSharp
Version: V8
`

	// Act

	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLanguage(models.LanguageTypes.CSharp, string(models.CSharpVersions.V8))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Language_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Version: V8
`

	// Act

	var data applicationproject.Language
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Type", fieldRequiredErr.FieldName)
}
