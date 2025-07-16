package models

import (
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	"parsdevkit.net/modules/resource/object_resource_contract"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	"parsdevkit.net/components/template/models/objectResources"
	objectResourceService "parsdevkit.net/components/template/services"
	"parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type FileTemplateDataContext struct {
	Workspace basic_workspace.WorkspaceComposite
	Group     basic_group.GroupComposite
	Project   application_project.ApplicationProjectComposite
	Resource  data_resource.DataResourceComposite
	Template  objectResources.FileTemplateComposite
	Layer     objectResources.DataLayerComposite
	Section   objectResources.DataSectionComposite
}

func NewFileTemplateDataContext(workspace basic_workspace_payload_structs.WorkspaceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, resource data_resource_payload_structs.ResourceBaseStruct, template file_template_payload_structs.TemplateBaseStruct, layer data_resource_payload_structs.Layer, section data_resource_payload_structs.Section) *FileTemplateDataContext {

	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)
	templateService := objectResourceService.NewObjectResourceService(manager)

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	groupService := ioc.Get[basic_group_contract.GroupInterface]()
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()
	return &FileTemplateDataContext{
		Workspace: workspaceService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(basic_workspace.WorkspaceComposite),
		Group:     groupService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(basic_group.GroupComposite),
		Project:   projectService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(application_project.ApplicationProjectComposite),
		Resource:  resourceService.Context(workspace, project, resource, template, layer.LayerIdentifier, section.SectionIdentifier).(data_resource.DataResourceComposite),
		Template: objectResources.FileTemplateComposite{
			FileTemplate: templateService.FileTemplateToModel(template),
			Original:     template,
		},
		Layer: objectResources.DataLayerComposite{
			DataLayer: templateService.DataLayerToModel(layer),
			Original:  layer,
		},
		Section: objectResources.DataSectionComposite{
			DataSection: templateService.DataSectionToModel(resource.Specifications, project.Specifications, layer.Name, template.Specifications, data_resource_payload_structs.Section{}),
			Original:    data_resource_payload_structs.Section{},
		},
	}
}
