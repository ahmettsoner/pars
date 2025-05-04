package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/errors"

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
		workspace.WorkspaceSpecification{},
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
		workspace.WorkspaceSpecification{},
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
		workspace.WorkspaceSpecification{},
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
		workspace.WorkspaceSpecification{},
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
		workspace.WorkspaceSpecification{},
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
		workspace.WorkspaceSpecification{},
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
			objectresource.NewAttribute("yea", objectresource.VisibilityTypeTypes.Private,
				objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
				0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			objectresource.NewAttribute("hoo", objectresource.VisibilityTypeTypes.Public,
				objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
				0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
		},
		[]objectresource.Method(nil),
		workspace.WorkspaceSpecification{},
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
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_ResourceSpecification_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Set: bar
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Name", fieldRequiredErr.FieldName)
}
func Test_UnMarshall_ResourceSpecification_WithoutSet(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
`

	// Act

	var data objectresource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Resource.Set", fieldRequiredErr.FieldName)
}
