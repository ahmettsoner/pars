package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

	"parsdevkit.net/application/schemas"
	goModels "parsdevkit.net/platforms/go/models"

	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/utilities/file"
)

// TODO: Testler tamamlanmalı
func Test_UnMarshall_ProjectSpecificationStruct_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Utils
ProjectType: Library
Group: Common
Set: Pars
Package: pars
Path: Utils
Workspace: pars-project
Platform: go@Go121
Configuration:
  Layers:
  Dependencies:
  - gopkg.in/yaml.v3@v3.0.1
  References:
  - Name: Logging
    Group: Core
    Workspace: pars-project
`

	// Act

	var data applicationproject.ProjectSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewProjectSpecification(
		0,
		"Utils",
		"Common",
		"pars-project",
		models.ProjectTypes.Library,
		applicationGroup.GroupIdentifier{},
		"Pars",
		[]string{"pars"},
		[]label.Label(nil),
		file.PathToArray("Utils"),
		applicationWorkspace.WorkspaceIdentifier{},
		applicationproject.NewPlatform(models.PlatformTypes.GO, goModels.GoPlatformVersions.Go121.String()),
		applicationproject.Runtime{},
		applicationproject.NewSchema(),
		applicationproject.NewConfiguration(
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency{
				applicationProject.NewDependency("gopkg.in/yaml.v3", "v3.0.1"),
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
		),
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
