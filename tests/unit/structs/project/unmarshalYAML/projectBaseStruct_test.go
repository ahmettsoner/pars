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
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

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
  Layers:

  Dependencies:
  - gopkg.in/yaml.v3@v3.0.1
  References:
  - Name: Logging
    Group: Core
    Workspace: pars-project
`

	// Act

	var data application_project_payload_structs.ProjectBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.ProjectBaseStruct{
		Header: schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, "Pars.CMD", schemas.Metadata{
			Tags: []string{"tag1", "tag2"},
		},
		),
		Specifications: application_project_payload_structs.NewProjectSpecification(
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
			application_project_payload_structs.NewPlatform(models.PlatformTypes.GO, goModels.GoPlatformVersions.Go121.String()),
			application_project_payload_structs.NewRuntime("", ""),
			application_project_payload_structs.NewSchema(),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency{
				applicationProject.NewDependency("gopkg.in/yaml.v3", "v3.0.1"),
			},
			[]application_project_payload_structs.ProjectBaseStruct{
				application_project_payload_structs.NewProjectBaseStruct(
					schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, "Logging", schemas.Metadata{}),
					application_project_payload_structs.NewProjectSpecification(
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
						application_project_payload_structs.Platform{},
						application_project_payload_structs.Runtime{},
						application_project_payload_structs.Schema{},
						[]applicationProject.Layer(nil),
						[]applicationProject.Dependency(nil),
						[]application_project_payload_structs.ProjectBaseStruct(nil),
						[]string(nil),
						[]string(nil),
						[]string(nil),
						[]string(nil),
					),
				),
			},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
		),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
