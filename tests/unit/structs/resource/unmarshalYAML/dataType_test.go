package unmarshalYAML

import (
	"testing"

	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType(objectresource.ValueTypes.Int.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("language.Language", objectresource.TypePackage{}, objectresource.DataTypeCategories.Resource, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType(objectresource.ValueTypes.Int.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("language.Language", objectresource.TypePackage{}, objectresource.DataTypeCategories.Resource, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("language.Language", objectresource.TypePackage{}, objectresource.DataTypeCategories.Reference, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("language.Language", objectresource.NewTypePackageOnly("type_pack"), objectresource.DataTypeCategories.Reference, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType(objectresource.ValueTypes.Int.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Array, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType(objectresource.ValueTypes.Int.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Resource, objectresource.ModifierTypes.Object, []objectresource.DataType(nil))

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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("Dictionary", objectresource.TypePackage{}, objectresource.DataTypeCategories.Reference, objectresource.ModifierTypes.Object, []objectresource.DataType{
		objectresource.NewDataType(objectresource.ValueTypes.String.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
		objectresource.NewDataType(objectresource.ValueTypes.Int.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
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

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewDataType("Dictionary", objectresource.TypePackage{}, objectresource.DataTypeCategories.Reference, objectresource.ModifierTypes.Object, []objectresource.DataType{
		objectresource.NewDataType(objectresource.ValueTypes.String.String(), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
		objectresource.NewDataType("language.Language", objectresource.TypePackage{}, objectresource.DataTypeCategories.Reference, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_DataType_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Category: reference
`

	// Act

	var data objectresource.DataType
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("DataType.Name", fieldRequiredErr.FieldName)
}
