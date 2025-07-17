package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	applicationTemplate "parsdevkit.net/application/structs/template"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_TemplateSpecification_File_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(
		0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_File_Object(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Template:
  Source: file
  Content: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Code_Object(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Template:
  Source: code
  Content: code_sample
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(
		0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.Code, "code_sample"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Output_File_Object(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: 
  File: filename.ext
Package: pars/cmd
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Package_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Package: pars/cmd
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Package_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Package: 
- pars
- cmd
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Layer_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Layers:
- layer1
- layer2
Template:
  file: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer{
			layerPkg.NewLayer(0, "layer1", []sectionPkg.Section(nil)),
			layerPkg.NewLayer(0, "layer2", []sectionPkg.Section(nil)),
		},
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_Label_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Labels:
- label1
- label2=value2
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label{
			label.NewLabel_KeyOnly("label1"),
			label.NewLabel("label2", "value2"),
		},
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TemplateSpecification_LabelObject_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Set: Set
Output: filename.ext
Labels:
- label1
- Key: label2
  Value: value2
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label{
			label.NewLabel_KeyOnly("label1"),
			label.NewLabel("label2", "value2"),
		},
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_TemplateSpecification_ID_ShouldBeZero(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Id: 100
Name: CMD
Set: Set
Output: filename.ext
Template:
  File: path
`

	// Act

	var data applicationTemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		applicationTemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]layerPkg.Layer(nil),
		applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
		applicationWorkspace.WorkspaceIdentifier{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
