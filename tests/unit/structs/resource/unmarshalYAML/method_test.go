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

func Test_UnMarshall_Method_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Name_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_WithVisibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Visibility: private
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Private,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_WithParameters(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Parameters:
- ID Int
- Name String
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter{
			objectresource.NewMethodParameter("ID", objectresource.New_Int(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
			objectresource.NewMethodParameter("Name", objectresource.New_String(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
		},
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Method_WithReturnTypes(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Returns:
- Int
- Name: List
  Category: reference
  Generics:
    - String
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType{
			objectresource.New_Int(),
			objectresource.New_Generic_Reference("List", objectresource.TypePackage{}, objectresource.New_String()),
		},
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Hint_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint: message_text
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Hint_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint:
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")),
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Method_Description_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description: message_text
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Description_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description:
  RefMessage: validationRules_patient_filter
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("validationRules_patient_filter")),
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)
	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Options(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Options:
- row=1
- column=3
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Labels(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Labels:
- foo=bar
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]objectresource.Annotation(nil),
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Annotations(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Annotations:
- Type: "auth.annotation"
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation{
			objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument(nil)),
		},
		"",
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Code(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Code: |
  print("hello world!")`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Public,
		[]objectresource.MethodParameter(nil),
		[]objectresource.DataType(nil),
		objectresource.Message{},
		objectresource.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]objectresource.Annotation(nil),
		`print("hello world!")`,
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Visibility: private
Parameters:
- ID Int
- Name String
Returns:
- Int
- Name: List
  Category: reference
  Generics:
    - String
Hint: message_text
Description: message_text
Options:
- row=1
- column=3
Labels:
- foo=bar
Annotations:
- Type: "auth.annotation"
Code: |
  print("hello world!")`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewMethod("foo", objectresource.VisibilityTypeTypes.Private,
		[]objectresource.MethodParameter{
			objectresource.NewMethodParameter("ID", objectresource.New_Int(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
			objectresource.NewMethodParameter("Name", objectresource.New_String(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
		},
		[]objectresource.DataType{
			objectresource.New_Int(),
			objectresource.New_Generic_Reference("List", objectresource.TypePackage{}, objectresource.New_String()),
		},
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		objectresource.NewMessage("message_text", objectresource.NewDictionaryIdentifier("")),
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]objectresource.Annotation{
			objectresource.NewAnnotation("auth.annotation", []objectresource.MethodArgument(nil)),
		},
		`print("hello world!")`,
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Method_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Int
`

	// Act

	var data objectresource.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Method.Name", fieldRequiredErr.FieldName)
}
