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

func Test_UnMarshall_ResourceObject_Attributes(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Attributes:
- Name: yea
  Visibility: private
- Name: hoo
  Type: Int
`

	// Act

	var data object_resource_payload_structs.ResourceObject
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceObject(
		[]object_resource_payload_structs.Attribute{
			object_resource_payload_structs.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
				structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
			object_resource_payload_structs.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
				structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
		},
		[]object_resource_payload_structs.Method(nil),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceObject_Methods(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Methods:
- Name: soe
  Parameters:
  - ID Int
  - Name String
`

	// Act

	var data object_resource_payload_structs.ResourceObject
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceObject(
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method{
			object_resource_payload_structs.NewMethod("soe", structs.VisibilityTypeTypes.Public,
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
			),
		},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ResourceObject_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Attributes:
- Name: yea
  Visibility: private
- Name: hoo
  Type: Int
Methods:
- Name: soe
  Parameters:
  - ID Int
  - Name String
`

	// Act

	var data object_resource_payload_structs.ResourceObject
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceObject(
		[]object_resource_payload_structs.Attribute{
			object_resource_payload_structs.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
				structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
			object_resource_payload_structs.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
				structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
		},
		[]object_resource_payload_structs.Method{
			object_resource_payload_structs.NewMethod("soe", structs.VisibilityTypeTypes.Public,
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
			),
		},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
