package models

import (
	objectResourceService "parsdevkit.net/components/template/services"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/application/ioc"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	"parsdevkit.net/application/platforms"
	"parsdevkit.net/components/template/models/objectResources"
	"parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type CodeTemplateDataContext struct {
	Workspace basic_workspace.WorkspaceComposite
	Project   application_project.ApplicationProjectComposite
	Resource  objectResources.ObjectResourceComposite
	Template  objectResources.CodeTemplateComposite
	Layer     objectResources.ObjectLayerComposite
	Section   objectResources.ObjectSectionComposite
}

func NewCodeTemplateDataContext(workspace basic_workspace_payload_structs.WorkspaceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, resource object_resource_payload_structs.ResourceBaseStruct, template code_template_payload_structs.TemplateBaseStruct, layer object_resource_payload_structs.Layer, section object_resource_payload_structs.Section) *CodeTemplateDataContext {

	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)
	templateService := objectResourceService.NewObjectResourceService(manager)

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	return &CodeTemplateDataContext{
		Workspace: workspaceService.Context(workspace).(basic_workspace.WorkspaceComposite),
		Project:   projectService.Context(project).(application_project.ApplicationProjectComposite),
		Resource: objectResources.ObjectResourceComposite{
			ObjectResource: templateService.ResourceToModel(resource, project, layer, template),
			Original:       resource,
		},
		Template: objectResources.CodeTemplateComposite{
			CodeTemplate: templateService.CodeTemplateToModel(template),
			Original:     template,
		},
		Layer: objectResources.ObjectLayerComposite{
			ObjectLayer: templateService.ObjectLayerToModel(layer),
			Original:    layer,
		},
		Section: objectResources.ObjectSectionComposite{
			ObjectSection: templateService.ObjectSectionToModel(resource, project, layer, template, section),
			Original:      section,
		},
	}
}
