package models

import (
	"parsdevkit.net/application/platforms"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	"parsdevkit.net/components/template/models/objectResources"
	objectResourceService "parsdevkit.net/components/template/services"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type FileTemplateDataContext struct {
	Workspace objectResources.WorkspaceComposite
	Project   objectResources.ApplicationProjectComposite
	Resource  objectResources.DataResourceComposite
	Template  objectResources.FileTemplateComposite
	Layer     objectResources.DataLayerComposite
	Section   objectResources.DataSectionComposite
}

func NewFileTemplateDataContext(workspace basic_workspace_payload_structs.WorkspaceBaseStruct, project application_project_payload_structs.ProjectBaseStruct, resource data_resource_payload_structs.ResourceBaseStruct, template file_template_payload_structs.TemplateBaseStruct, layer data_resource_payload_structs.Layer) *FileTemplateDataContext {

	manager := platforms.Get[application_project_payload_structs.ProjectBaseStruct](project.Specifications.Platform.Type)
	templateService := objectResourceService.NewObjectResourceService(manager)

	return &FileTemplateDataContext{
		Workspace: objectResources.WorkspaceComposite{
			Workspace: templateService.WorkspaceToModel(workspace),
			Original:  workspace,
		},
		Project: objectResources.ApplicationProjectComposite{
			ApplicationProject: templateService.ApplicationProjectToModel(project),
			Original:           project,
		},
		Resource: objectResources.DataResourceComposite{
			DataResource: templateService.DataResourceToModel(resource.Specifications, resource.Data, project.Specifications, layer.Name, template.Specifications),
			Original:     resource,
		},
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
