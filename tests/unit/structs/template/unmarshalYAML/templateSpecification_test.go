package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/label"
	templateStruct "parsdevkit.net/structs/template"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"

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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(
		0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(
		0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.Code, "code_sample"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string{"pars", "cmd"},
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]codetemplate.Layer{
			codetemplate.NewLayer(0, "layer1", []templateStruct.Section(nil)),
			codetemplate.NewLayer(0, "layer2", []templateStruct.Section(nil)),
		},
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label{
			label.NewLabel_KeyOnly("label1"),
			label.NewLabel("label2", "value2"),
		},
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label{
			label.NewLabel_KeyOnly("label1"),
			label.NewLabel("label2", "value2"),
		},
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
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

	var data codetemplate.TemplateSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateSpecification(0,
		"CMD",
		"",
		"Set",
		"",
		codetemplate.NewOutput("filename.ext"),
		[]string(nil),
		[]label.Label(nil),
		[]codetemplate.Layer(nil),
		codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
		workspace.WorkspaceSpecification{},
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
