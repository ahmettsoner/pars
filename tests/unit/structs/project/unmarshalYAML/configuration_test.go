package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Configuration_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Layers:
Options:
Dependencies:
- gopkg.in/yaml.v3@v3.0.1
References:
- Name: Logging
  Group: Core
  Workspace: pars-project
`

	// Act

	var data applicationproject.Configuration
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewConfiguration(
		[]applicationproject.Layer(nil),
		[]applicationproject.Package{
			applicationproject.NewPackage("gopkg.in/yaml.v3", "v3.0.1"),
		},
		[]applicationproject.ProjectBaseStruct{
			applicationproject.NewProjectBaseStruct(
				schemas.NewSchemaHeader(schemas.StructTypes.Project, string(project.ProjectKinds.Application), "Logging", schemas.Metadata{}),
				applicationproject.NewProjectSpecification(
					0,
					"",
					"Core",
					"pars-project",
					"",
					group.GroupSpecification{},
					"",
					[]string(nil),
					[]label.Label(nil),
					[]string(nil),
					workspace.WorkspaceSpecification{},
					applicationproject.Platform{},
					applicationproject.Runtime{},
					applicationproject.Schema{},
					applicationproject.Configuration{},
				),
			),
		},
		[]string(nil),
		[]string(nil),
		[]string(nil),
		[]string(nil),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
