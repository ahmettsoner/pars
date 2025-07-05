package models

import (
	objectResourceService "parsdevkit.net/components/template/services"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
	objectresource "parsdevkit.net/modules/resource/object_resource_payload"
	codetemplate "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/application/platforms"
	"parsdevkit.net/components/template/models/objectResources"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"
)

type CodeTemplateDataContext struct {
	Workspace objectResources.WorkspaceComposite
	Project   objectResources.ApplicationProjectComposite
	Resource  objectResources.ObjectResourceComposite
	Template  objectResources.CodeTemplateComposite
	Layer     objectResources.ObjectLayerComposite
	Section   objectResources.ObjectSectionComposite
}

func NewCodeTemplateDataContext(workspace basic_workspace_payload.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource objectresource.ResourceBaseStruct, template codetemplate.TemplateBaseStruct, layer objectresource.Layer, section objectresource.Section) *CodeTemplateDataContext {

	manager := platforms.Get[applicationproject.ProjectBaseStruct](project.Specifications.Platform.Type)
	templateService := objectResourceService.NewObjectResourceService(manager)

	return &CodeTemplateDataContext{
		Workspace: objectResources.WorkspaceComposite{
			Workspace: templateService.WorkspaceToModel(workspace),
			Original:  workspace,
		},
		Project: objectResources.ApplicationProjectComposite{
			ApplicationProject: templateService.ApplicationProjectToModel(project),
			Original:           project,
		},
		Resource: objectResources.ObjectResourceComposite{
			ObjectResource: templateService.ResourceToModel(resource.Specifications, project.Specifications, layer.Name, template.Specifications),
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
			ObjectSection: templateService.ObjectSectionToModel(resource.Specifications, project.Specifications, layer.Name, template.Specifications, section),
			Original:      section,
		},
	}
}
