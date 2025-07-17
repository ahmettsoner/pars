package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/models"

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

	var data applicationProject.ProjectSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewProjectSpecification(
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
	)

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
