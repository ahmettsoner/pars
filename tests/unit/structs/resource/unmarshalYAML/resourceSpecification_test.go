package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	"parsdevkit.net/application/structs"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	applicationResource "parsdevkit.net/application/structs/resource"
)

func Test_UnMarshall_ResourceObject_NameAndSet(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: foo
Set: bar
`

	// Act

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"/foo",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label{
			label.NewLabel("foo", "bar"),
		},
		[]layerPkg.Layer(nil),
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
		"foo",
		"",
		"",
		"bar",
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer{layerPkg.NewLayer(0, "layer1", []sectionPkg.Section(nil)), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section(nil))},
		applicationWorkspace.WorkspaceIdentifier{},
		structs.ChangeTrackers.OnChange,
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
`

	// Act

	var data applicationResource.ResourceSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationResource.NewResourceSpecification(0,
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
		structs.ChangeTrackers.OnChange,
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
