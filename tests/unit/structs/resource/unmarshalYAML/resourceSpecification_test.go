package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	objectresource "parsdevkit.net/modules/resource/object_resource_payload"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_ResourceSpecification_NameAndSet(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceSpecification_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Path: /foo
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"/foo",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceSpecification_Package_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Package: pars/cmd
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceSpecification_Package_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Package:
- pars
- cmd
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceSpecification_Labels(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Labels:
- foo=bar
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ResourceSpecification_Layers(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Layers:
- layer1
- layer2
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]objectresource.Layer{objectresource.NewLayer(0, "layer1", []objectresource.Section(nil)), objectresource.NewLayer(0, "layer2", []objectresource.Section(nil))},
		[]objectresource.Attribute(nil),
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ResourceSpecification_Attributes(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Attributes:
- Name: yea
  Visibility: private
- Name: hoo
  Type: Int
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute{
			objectresource.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
				structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			objectresource.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
				structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
		},
		[]objectresource.Method(nil),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_ResourceSpecification_Methods(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
Methods:
- Name: soe
  Parameters:
  - ID Int
  - Name String
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]objectresource.Layer(nil),
		[]objectresource.Attribute(nil),
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
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ResourceSpecification_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
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

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewResourceSpecification(0,
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
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
