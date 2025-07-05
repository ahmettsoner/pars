package models

import (
	"parsdevkit.net/application/platforms"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
	dataresource "parsdevkit.net/modules/resource/data_resource_payload"
	filetemplate "parsdevkit.net/structs/template/file-template"

	"parsdevkit.net/components/template/models/objectResources"
	objectResourceService "parsdevkit.net/components/template/services"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"
)

type FileTemplateDataContext struct {
	Workspace objectResources.WorkspaceComposite
	Project   objectResources.ApplicationProjectComposite
	Resource  objectResources.DataResourceComposite
	Template  objectResources.FileTemplateComposite
	Layer     objectResources.DataLayerComposite
	Section   objectResources.DataSectionComposite
}

func NewFileTemplateDataContext(workspace basic_workspace_payload.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource dataresource.ResourceBaseStruct, template filetemplate.TemplateBaseStruct, layer dataresource.Layer) *FileTemplateDataContext {

	manager := platforms.Get[applicationproject.ProjectBaseStruct](project.Specifications.Platform.Type)
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
			DataResource: templateService.DataResourceToModel(resource.Specifications, project.Specifications, layer.Name, template.Specifications),
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
			DataSection: templateService.DataSectionToModel(resource.Specifications, project.Specifications, layer.Name, template.Specifications, dataresource.Section{}),
			Original:    dataresource.Section{},
		},
	}
}
