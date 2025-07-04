package objects

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"
)

func BasicProject_WithName(name string, projectType models.ProjectType, platform models.PlatformType, runtime models.RuntimeType, workspace workspace.WorkspaceSpecification) *applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Project,
			applicationproject.PROJECT_KIND,
			name,
			schemas.Metadata{
				Tags: []string(nil),
			},
		),
		applicationproject.NewProjectSpecification(0,
			name,
			"",
			workspace.Name,
			projectType,
			applicationGroup.GroupIdentifier{},
			"",
			[]string(nil),
			[]label.Label(nil),
			[]string(nil),
			workspace.WorkspaceIdentifier,
			applicationproject.NewPlatform_Basic(platform),
			applicationproject.NewRuntime_Basic(runtime),
			applicationproject.NewSchema(),
			applicationproject.NewConfiguration(
				[]applicationProject.Layer(nil),
				[]applicationProject.Package(nil),
				[]applicationproject.ProjectBaseStruct(nil),
				[]string(nil),
				[]string(nil),
				[]string(nil),
				[]string(nil),
			),
		),
	)
	return &project
}
