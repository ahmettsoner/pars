package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_MethodParameter_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Name_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_WithType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Type: Int
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_WithObjectType_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo Int
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_WithArrayTypeOnly_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo Int[]
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Array, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_WithOrder(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Order: 3
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 3, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_MethodParameter_Hint_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint: message_text
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")), objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Hint_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint: 
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")), objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_MethodParameter_Description_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description: message_text
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")), []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Description_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description: 
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")), []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Options(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Options:
- row=1
- column=3
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option{
		option.NewOption("row", "1"),
		option.NewOption("column", "3"),
	}, []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Labels(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Labels:
- foo=bar
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label{
		label.NewLabel("foo", "bar"),
	}, objectresource.Validation{}, []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Validation(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Validation:
- Length: 10:150
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil),
		objectresource.NewValidation(
			objectresource.NewValidationLengthRule("", 10, 150, objectresource.Message{}),
		), []objectresource.Annotation(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_MethodParameter_Annotations(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Annotations:
- Type: "auth.annotation"
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter("foo", objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation{
		objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument(nil)),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Type: Int
Order: 3
Hint: message_text
Description: message_text
Options:
- row=1
- column=3
Labels:
- foo=bar
Validation:
- Length: 10:150
Annotations:
- Type: "auth.annotation"
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethodParameter(
		"foo",
		objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
		3,
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		objectresource.NewValidation(
			objectresource.NewValidationLengthRule("", 10, 150, objectresource.Message{}),
		),
		[]objectresource.Annotation{
			objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument(nil)),
		},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_MethodParameter_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Int
`

	// Act

	var data objectresource.MethodParameter
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Variable.Name", fieldRequiredErr.FieldName)
}
