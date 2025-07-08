package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Private,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter{
			object_resource_payload_structs.NewMethodParameter("ID", structs.New_Int(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
			object_resource_payload_structs.NewMethodParameter("Name", structs.New_String(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
		},
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType{
			structs.New_Int(),
			structs.New_Generic_Reference("List", structs.TypePackage{}, structs.New_String()),
		},
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")),
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")),
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation{
			object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument(nil)),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Public,
		[]object_resource_payload_structs.MethodParameter(nil),
		[]structs.DataType(nil),
		object_resource_payload_structs.Message{},
		object_resource_payload_structs.Message{},
		[]option.Option(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Annotation(nil),
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

	var data object_resource_payload_structs.Method
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewMethod("foo", structs.VisibilityTypeTypes.Private,
		[]object_resource_payload_structs.MethodParameter{
			object_resource_payload_structs.NewMethodParameter("ID", structs.New_Int(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
			object_resource_payload_structs.NewMethodParameter("Name", structs.New_String(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
		},
		[]structs.DataType{
			structs.New_Int(),
			structs.New_Generic_Reference("List", structs.TypePackage{}, structs.New_String()),
		},
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]object_resource_payload_structs.Annotation{
			object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument(nil)),
		},
		`print("hello world!")`,
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
