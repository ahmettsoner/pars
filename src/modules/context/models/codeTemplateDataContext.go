package models

import (
	objectResourceService "parsdevkit.net/components/template/services"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/application/ioc"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	"parsdevkit.net/application/platforms"
	"parsdevkit.net/components/template/models/objectResources"
	"parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/object_resource"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type CodeTemplateDataContext struct {
	Workspace basic_workspace.WorkspaceComposite
	Group     basic_group.GroupComposite
	Project   application_project.ApplicationProjectComposite
	Resource  object_resource.ObjectResourceComposite
	Template  objectResources.CodeTemplateComposite
	Layer     objectResources.ObjectLayerComposite
	Section   objectResources.ObjectSectionComposite
}

func NewCodeTemplateDataContext(workspace basic_workspace_payload_structs.WorkspaceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, resource object_resource_payload_structs.ResourceBaseStruct, template code_template_payload_structs.TemplateBaseStruct, layer object_resource_payload_structs.Layer, section object_resource_payload_structs.Section) *CodeTemplateDataContext {

	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)
	templateService := objectResourceService.NewObjectResourceService(manager)

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	groupService := ioc.Get[basic_group_contract.GroupInterface]()
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()

	return &CodeTemplateDataContext{
		Workspace: workspaceService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(basic_workspace.WorkspaceComposite),
		Group:     groupService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(basic_group.GroupComposite),
		Project:   projectService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(application_project.ApplicationProjectComposite),
		Resource:  resourceService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(object_resource.ObjectResourceComposite),
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
