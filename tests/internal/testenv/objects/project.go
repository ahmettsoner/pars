package objects

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

func BasicProject_WithName(name string, projectType models.ProjectType, platform models.PlatformType, runtime models.RuntimeType, workspace basic_workspace_payload_structs.WorkspaceSpecification) *application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Project,
			application_project_payload_structs.PROJECT_KIND,
			name,
			schemas.Metadata{
				Tags: []string(nil),
			},
		),
		application_project_payload_structs.NewProjectSpecification(0,
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
			application_project_payload_structs.NewPlatform_Basic(platform),
			application_project_payload_structs.NewRuntime_Basic(runtime),
			application_project_payload_structs.NewSchema(),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
		),
	)
	return &project
}
