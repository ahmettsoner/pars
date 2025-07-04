package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

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
		[]applicationProject.Layer(nil),
		[]applicationProject.Package{
			applicationProject.NewPackage("gopkg.in/yaml.v3", "v3.0.1"),
		},
		[]applicationproject.ProjectBaseStruct{
			applicationproject.NewProjectBaseStruct(
				schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, "Logging", schemas.Metadata{}),
				applicationproject.NewProjectSpecification(
					0,
					"",
					"Core",
					"pars-project",
					"",
					applicationGroup.GroupIdentifier{},
					"",
					[]string(nil),
					[]label.Label(nil),
					[]string(nil),
					applicationWorkspace.WorkspaceIdentifier{},
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
