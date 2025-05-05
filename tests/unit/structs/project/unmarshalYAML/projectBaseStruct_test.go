package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/models"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	goModels "parsdevkit.net/platforms/go/models"

	"parsdevkit.net/core/utils"

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
		Header: project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, "Pars.CMD", structs.Metadata{
			Tags: []string{"tag1", "tag2"},
		},
		),
		Specifications: applicationproject.NewProjectSpecification(
			0,
			"Utils",
			"Common",
			"pars-project",
			models.ProjectTypes.Library,
			group.GroupSpecification{},
			"Pars",
			[]string{"pars"},
			[]label.Label(nil),
			utils.PathToArray("Utils"),
			workspace.WorkspaceSpecification{},
			applicationproject.NewPlatform(models.PlatformTypes.GO, goModels.GoPlatformVersions.Go121.String()),
			applicationproject.NewRuntime("", ""),
			applicationproject.NewSchema(),
			applicationproject.NewConfiguration(
				[]applicationproject.Layer(nil),
				[]applicationproject.Package{
					applicationproject.NewPackage("gopkg.in/yaml.v3", "v3.0.1"),
				},
				[]applicationproject.ProjectBaseStruct{
					applicationproject.NewProjectBaseStruct(
						project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, "Logging", structs.Metadata{}),
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
			),
		),
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
