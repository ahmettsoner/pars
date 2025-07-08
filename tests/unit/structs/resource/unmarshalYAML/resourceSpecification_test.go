package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"/foo",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer{object_resource_payload_structs.NewLayer(0, "layer1", []object_resource_payload_structs.Section(nil)), object_resource_payload_structs.NewLayer(0, "layer2", []object_resource_payload_structs.Section(nil))},
		[]object_resource_payload_structs.Attribute(nil),
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
		[]object_resource_payload_structs.Attribute{
			object_resource_payload_structs.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
				structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
			object_resource_payload_structs.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
				structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
				0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
		},
		[]object_resource_payload_structs.Method(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]object_resource_payload_structs.Layer(nil),
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

	var data object_resource_payload_structs.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := object_resource_payload_structs.NewResourceSpecification(0,
		"foo",
		"",
		"/foo",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]object_resource_payload_structs.Layer{object_resource_payload_structs.NewLayer(0, "layer1", []object_resource_payload_structs.Section(nil)), object_resource_payload_structs.NewLayer(0, "layer2", []object_resource_payload_structs.Section(nil))},
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
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
