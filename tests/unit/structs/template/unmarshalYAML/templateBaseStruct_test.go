package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/template"
	templateStruct "parsdevkit.net/structs/template"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_TemplateBaseStruct_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Template
Kind: Code
Name: Entity
Metadata:
  Tags:
  - tag1
  - tag2
Specifications:
  Name: Entity
  Set: Set
  Output: "{{ .Name }}.cs"
  Package: pars/cmd
  Template:
    File: path
`

	// Act

	var data codetemplate.TemplateBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := codetemplate.NewTemplateBaseStruct(
		templateStruct.NewHeader(structs.StructTypes.Template, templateStruct.StructKinds.Code, "Entity", structs.NewMetadata([]string{"tag1", "tag2"})),
		codetemplate.NewTemplateSpecification(0,
			"Entity",
			"",
			"Set",
			"",
			codetemplate.NewOutput("{{ .Name }}.cs"),
			[]string{"pars", "cmd"},
			[]label.Label(nil),
			[]codetemplate.Layer(nil), codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.File, "path"),
			workspace.WorkspaceSpecification{},
		),
		codetemplate.NewTemplateConfiguration(codetemplate.ChangeTrackers.OnChange, template.Selectors{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
