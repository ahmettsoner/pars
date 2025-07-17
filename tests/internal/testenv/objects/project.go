package objects

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

func BasicProject_WithName(name string, projectType models.ProjectType, platform models.PlatformType, runtime models.RuntimeType, workspace applicationWorkspace.WorkspaceSpecification) *application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Project,
			application_project_payload_structs.PROJECT_KIND,
			name,
			schemas.Metadata{
				Tags: []string(nil),
			},
		),
		applicationProject.NewProjectSpecification(0,
			name,
			"",
			workspace.Name,
			applicationGroup.GroupIdentifier{},
			"",
			[]string(nil),
			[]label.Label(nil),
			[]string(nil),
			workspace.WorkspaceIdentifier,
			applicationProject.NewPlatform_Basic(platform),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency(nil),
			[]applicationProject.Reference(nil),
		),
		application_project_payload_structs.NewApplication(
			projectType,
			application_project_payload_structs.NewRuntime_Basic(runtime),
			application_project_payload_structs.NewSchema(),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	return &project
}
