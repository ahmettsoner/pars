package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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
		resource.NewHeader(
			structs.StructTypes.Resource,
			resource.StructKinds.Object,
			"Pars.CMD",
			structs.Metadata{
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
				objectresource.NewAttribute("yea", objectresource.VisibilityTypeTypes.Private,
					objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
				objectresource.NewAttribute("hoo", objectresource.VisibilityTypeTypes.Public,
					objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			},
			[]objectresource.Method{
				objectresource.NewMethod("soe", objectresource.VisibilityTypeTypes.Public,
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
				),
			},
			workspace.WorkspaceSpecification{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
