package objects

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/models"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"
)

func BasicProject_WithName(name string, projectType models.ProjectType, platform models.PlatformType, runtime models.RuntimeType, workspace workspace.WorkspaceSpecification) *applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Project,
			string(project.ProjectKinds.Application),
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
			group.GroupSpecification{},
			"",
			[]string(nil),
			[]label.Label(nil),
			[]string(nil),
			workspace,
			applicationproject.NewPlatform_Basic(platform),
			applicationproject.NewRuntime_Basic(runtime),
			applicationproject.NewSchema(),
			applicationproject.NewConfiguration(
				[]applicationproject.Layer(nil),
				[]applicationproject.Package(nil),
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
