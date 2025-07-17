package unmarshalYAML

import (
	"testing"

	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	applicationResource "parsdevkit.net/application/structs/resource"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
)

func Test_UnMarshall_ResourceBaseStruct_ObjectKind_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Resource
Kind: Object
Name:  Pars.CMD
Metadata:
  Tags: tag1, tag2
Specifications:
  Name: foo
  Set: bar
  Path: /foo
  Package: pars/cmd
  Labels:
  - foo=bar
  Layers:
  - layer1
  - layer2
Object:
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

	var data object_resource_payload_structs.ResourceBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			object_resource_payload_structs.RESOURCE_KIND,
			"Pars.CMD",
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		applicationResource.NewResourceSpecification(0,
			"foo",
			"",
			"/foo",
			"bar",
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]layerPkg.Layer{layerPkg.NewLayer(0, "layer1", []sectionPkg.Section(nil)), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section(nil))},
			applicationWorkspace.WorkspaceIdentifier{},
		),
		object_resource_payload_structs.NewResourceObject(
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
		),
		object_resource_payload_structs.NewResourceConfiguration(object_resource_payload_structs.ChangeTrackers.OnChange),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
