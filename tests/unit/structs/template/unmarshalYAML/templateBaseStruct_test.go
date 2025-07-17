package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	templateStruct "parsdevkit.net/application/structs/template"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	applicationTemplate "parsdevkit.net/application/structs/template"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
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

	var data code_template_payload_structs.TemplateBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := code_template_payload_structs.NewTemplateBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Template, code_template_payload_structs.TEMPLATE_KIND, "Entity", schemas.NewMetadata([]string{"tag1", "tag2"})),
		applicationTemplate.NewTemplateSpecification(0,
			"Entity",
			"",
			"Set",
			"",
			applicationTemplate.NewOutput("{{ .Name }}.cs"),
			[]string{"pars", "cmd"},
			[]label.Label(nil),
			[]layerPkg.Layer(nil), applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.File, "path"),
			applicationWorkspace.WorkspaceIdentifier{},
		),
		code_template_payload_structs.NewTemplateConfiguration(code_template_payload_structs.ChangeTrackers.OnChange, templateStruct.Selectors{}),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
