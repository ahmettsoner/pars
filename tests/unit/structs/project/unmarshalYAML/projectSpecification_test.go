package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

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

Layers:
Dependencies:
- gopkg.in/yaml.v3@v3.0.1
References:
- Name: Logging
  Group: Core
  Workspace: pars-project
`

	// Act

	var data application_project_payload_structs.ProjectSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := application_project_payload_structs.NewProjectSpecification(
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
		application_project_payload_structs.Runtime{},
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
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
