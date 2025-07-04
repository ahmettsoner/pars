package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

	goModels "parsdevkit.net/platforms/go/models"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_ProjectBaseStruct_FullData(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Project
Kind: Application
Name:  Pars.CMD
Metadata:
  Tags: tag1, tag2
Specifications:
  Name: Utils
  Group: Common
  ProjectType: Library
  Set: Pars
  Package: pars
  Path: Utils
  Workspace: pars-project
  Platform: 
    Type: go
    Version: Go121


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

	var data applicationproject.ProjectBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.ProjectBaseStruct{
		Header: schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, "Pars.CMD", schemas.Metadata{
			Tags: []string{"tag1", "tag2"},
		},
		),
		Specifications: applicationproject.NewProjectSpecification(
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
			applicationproject.NewRuntime("", ""),
			applicationproject.NewSchema(),
			applicationproject.NewConfiguration(
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
			),
		),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
