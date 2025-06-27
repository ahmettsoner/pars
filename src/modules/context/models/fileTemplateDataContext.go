package models

import (
	platformsCommon "parsdevkit.net/platforms/common"
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	filetemplate "parsdevkit.net/structs/template/file-template"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/templates/models/objectResources"
	objectResourceService "parsdevkit.net/templates/services"
)

type FileTemplateDataContext struct {
	Workspace objectResources.WorkspaceComposite
	Project   objectResources.ApplicationProjectComposite
	Resource  objectResources.DataResourceComposite
	Template  objectResources.FileTemplateComposite
	Layer     objectResources.DataLayerComposite
	Section   objectResources.DataSectionComposite
}

func NewFileTemplateDataContext(workspace workspace.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource dataresource.ResourceBaseStruct, template filetemplate.TemplateBaseStruct, layer dataresource.Layer) *FileTemplateDataContext {
	manager, err := platformsCommon.ManagerFactory(project.Specifications.Platform.Type)
	if err != nil {
		// return ObjectResourceService{}, fmt.Errorf("xxx: Yeni object resource init aşamasında, Platform Manager bulunamadı '%s'\n%w", err)
	}
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
