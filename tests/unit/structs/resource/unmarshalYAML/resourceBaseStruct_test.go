package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"

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

	var data objectresource.ResourceBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			string(resource.ResourceKinds.Object),
			"Pars.CMD",
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		objectresource.NewResourceSpecification(0,
			"foo",
			"",
			"/foo",
			"bar",
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]objectresource.Layer{objectresource.NewLayer(0, "layer1", []objectresource.Section(nil)), objectresource.NewLayer(0, "layer2", []objectresource.Section(nil))},
			[]objectresource.Attribute{
				objectresource.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
					structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
				objectresource.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
					structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			},
			[]objectresource.Method{
				objectresource.NewMethod("soe", structs.VisibilityTypeTypes.Public,
					[]objectresource.MethodParameter{
						objectresource.NewMethodParameter("ID", structs.New_Int(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
						objectresource.NewMethodParameter("Name", structs.New_String(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
					},
					[]structs.DataType(nil),
					objectresource.Message{},
					objectresource.Message{},
					[]option.Option(nil),
					[]label.Label(nil),
					[]objectresource.Annotation(nil),
					"",
					true,
				),
			},
			applicationWorkspace.WorkspaceIdentifier{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
