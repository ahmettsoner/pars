package models

import (
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	filetemplate "parsdevkit.net/structs/template/file-template"

	"parsdevkit.net/components/template/models/objectResources"
	objectResourceService "parsdevkit.net/components/template/services"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"
	platformsCommon "parsdevkit.net/platforms/common"
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
	manager, err := platformsCommon.GetPlatformManager(project.Specifications.Platform.Type, platformsCommon.Registry)
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
