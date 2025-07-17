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
  Set: Pars
  Platform: 
    Type: go
    Version: Go121
  Package: pars
  Path: Utils
  Workspace: pars-project
  Layers:
  Dependencies:
  - gopkg.in/yaml.v3@v3.0.1
  References:
  - Name: Logging
    Group: Core
    Workspace: pars-project
Application:
  ProjectType: Library
`

	// Act

	var data application_project_payload_structs.ProjectBaseStruct
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.ProjectBaseStruct{
		Header: schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, "Pars.CMD", schemas.Metadata{
			Tags: []string{"tag1", "tag2"},
		},
		),
		Specifications: applicationProject.NewProjectSpecification(
			0,
			"Utils",
			"Common",
			"pars-project",
			applicationGroup.GroupIdentifier{},
			"Pars",
			[]string{"pars"},
			[]label.Label(nil),
			file.PathToArray("Utils"),
			applicationWorkspace.WorkspaceIdentifier{},
			applicationProject.NewPlatform(models.PlatformTypes.GO, goModels.GoPlatformVersions.Go121.String()),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency{
				applicationProject.NewDependency("gopkg.in/yaml.v3", "v3.0.1"),
			},
			[]applicationProject.Reference{
				applicationProject.NewReference(
					schemas.NewSchemaHeader(schemas.StructTypes.Project, "", "Logging", schemas.Metadata{}),
					applicationProject.NewProjectSpecification(
						0,
						"",
						"Core",
						"pars-project",
						applicationGroup.GroupIdentifier{},
						"",
						[]string(nil),
						[]label.Label(nil),
						[]string(nil),
						applicationWorkspace.WorkspaceIdentifier{},
						applicationProject.Platform{},
						[]applicationProject.Layer(nil),
						[]applicationProject.Dependency(nil),
						[]applicationProject.Reference(nil),
					),
				),
			},
		),
		Application: application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.NewRuntime("", ""),
			application_project_payload_structs.NewSchema(),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
