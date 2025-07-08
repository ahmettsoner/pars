package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Attribute_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(
			string(structs.ValueTypes.String),
			structs.TypePackage{},
			structs.DataTypeCategories.Value,
			structs.ModifierTypes.Object,
			[]structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Name_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_WithType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Type: Int
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_WithObjectType_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo Int
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_WithArrayTypeOnly_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
foo Int[]
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Array, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_WithVisibility(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Visibility: private
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Private,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_WithOrder(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Order: 3
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		3, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_WithGroup_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Group: CMD
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0,
		object_resource_payload_structs.NewAttributeGroup(object_resource_payload_structs.NewGroupIdentifier("CMD"), 0, []option.Option(nil)),
		object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_WithGroup(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Group: 
  RefGroup: CMD
  Order: 2
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0,
		object_resource_payload_structs.NewAttributeGroup(object_resource_payload_structs.NewGroupIdentifier("CMD"), 2, []option.Option(nil)),
		object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_WithEncapsulation(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Encapsulation:
  Getter: true
  Setter:
    Visibility: private
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{},
		object_resource_payload_structs.NewEncapsulation(
			object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{}, true),
			object_resource_payload_structs.NewEncapsulationSetter("", structs.VisibilityTypeTypes.Private, object_resource_payload_structs.MethodIdentifier{}, true),
		),
		object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_WithProperties(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Properties:
  Key: true
  Required: true
  ReadOnly: true
  Unique: true
  Default: 123
  Format: ddd
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0,
		object_resource_payload_structs.AttributeGroup{},
		object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.NewAttributeProperties(true, true, true, true, "123", "ddd"), object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_Hint_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint: message_text
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")), object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Hint_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Hint: 
  RefMessage: validationRules_patient_filter
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo",
		structs.VisibilityTypeTypes.Public, structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")), object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_Description_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description: message_text
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)
	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")), []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Description_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Description: 
  RefMessage: validationRules_patient_filter
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.NewMessage("", object_resource_payload_structs.NewDictionaryIdentifier("validationRules_patient_filter")), []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Options(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Options:
- row=1
- column=3
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		}, []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Labels(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Labels:
- foo=bar
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label{
			label.NewLabel("foo", "bar"),
		}, object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Validation(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Validation:
- Length: 10:150
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil),
		object_resource_payload_structs.NewValidation(
			object_resource_payload_structs.NewValidationLengthRule("", 10, 150, object_resource_payload_structs.Message{}),
		), []object_resource_payload_structs.Annotation(nil), true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Attribute_Annotations(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Annotations:
- Type: "auth.annotation"
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute("foo", structs.VisibilityTypeTypes.Public,
		structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation{
			object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument(nil)),
		}, true)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Attribute_Complete(t *testing.T) {

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
Visibility: private
Group: 
  RefGroup: CMD
  Order: 2
Encapsulation:
  Getter: true
  Setter:
    Visibility: private
Properties:
  Key: true
  Required: true
  ReadOnly: true
  Unique: true
  Default: 123
  Format: ddd
`

	// Act

	var data object_resource_payload_structs.Attribute
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewAttribute(
		"foo",
		structs.VisibilityTypeTypes.Private,
		structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
		3,
		object_resource_payload_structs.NewAttributeGroup(object_resource_payload_structs.NewGroupIdentifier("CMD"), 2, []option.Option(nil)),
		object_resource_payload_structs.NewEncapsulation(
			object_resource_payload_structs.NewEncapsulationGetter("", "", object_resource_payload_structs.MethodIdentifier{}, true),
			object_resource_payload_structs.NewEncapsulationSetter("", structs.VisibilityTypeTypes.Private, object_resource_payload_structs.MethodIdentifier{}, true),
		),
		object_resource_payload_structs.NewAttributeProperties(true, true, true, true, "123", "ddd"),
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		object_resource_payload_structs.NewMessage("message_text", object_resource_payload_structs.NewDictionaryIdentifier("")),
		[]option.Option{
			option.NewOption("row", "1"),
			option.NewOption("column", "3"),
		},
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		object_resource_payload_structs.NewValidation(
			object_resource_payload_structs.NewValidationLengthRule("", 10, 150, object_resource_payload_structs.Message{}),
		),
		[]object_resource_payload_structs.Annotation{
			object_resource_payload_structs.NewAnnotation("auth.annotation", []object_resource_payload_structs.MethodArgument(nil)),
		},
		true,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
