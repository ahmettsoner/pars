package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_DataType_Inline_Value_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Int
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType(structs.ValueTypes.Int.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Inline_Resource_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
language.Language
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("language.Language", structs.TypePackage{}, structs.DataTypeCategories.Resource, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Value_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Int
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType(structs.ValueTypes.Int.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Resource_TypeOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: language.Language
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("language.Language", structs.TypePackage{}, structs.DataTypeCategories.Resource, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Reference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: language.Language
Category: reference
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("language.Language", structs.TypePackage{}, structs.DataTypeCategories.Reference, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Reference_WithPackage(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: language.Language
Category: reference
Package: type_pack
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("language.Language", structs.NewTypePackageOnly("type_pack"), structs.DataTypeCategories.Reference, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Value_WithModifier(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Int
Modifier: array
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType(structs.ValueTypes.Int.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Array, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Value_WithCategory(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Int
Category: Resource
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType(structs.ValueTypes.Int.String(), structs.TypePackage{}, structs.DataTypeCategories.Resource, structs.ModifierTypes.Object, []structs.DataType(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Generic_Reference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Dictionary
Category: reference
Generics:
- String
- Int
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("Dictionary", structs.TypePackage{}, structs.DataTypeCategories.Reference, structs.ModifierTypes.Object, []structs.DataType{
		structs.NewDataType(structs.ValueTypes.String.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		structs.NewDataType(structs.ValueTypes.Int.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_Generic_Reference_WithReferenceArgument(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Dictionary
Category: reference
Generics:
- String
- Name: language.Language
  Category: reference
`

	// Act

	var data structs.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.NewDataType("Dictionary", structs.TypePackage{}, structs.DataTypeCategories.Reference, structs.ModifierTypes.Object, []structs.DataType{
		structs.NewDataType(structs.ValueTypes.String.String(), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		structs.NewDataType("language.Language", structs.TypePackage{}, structs.DataTypeCategories.Reference, structs.ModifierTypes.Object, []structs.DataType(nil)),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
