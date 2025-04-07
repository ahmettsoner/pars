package models

import (
	applicationproject "parsdevkit.net/structs/project/application-project"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"
	objectResourceService "parsdevkit.net/templates/services"

	"parsdevkit.net/templates/models/objectResources"
)

type CodeTemplateDataContext struct {
	Workspace objectResources.WorkspaceComposite
	Project   objectResources.ApplicationProjectComposite
	Resource  objectResources.ObjectResourceComposite
	Template  objectResources.CodeTemplateComposite
	Layer     objectResources.ObjectLayerComposite
	Section   objectResources.ObjectSectionComposite
}

func NewCodeTemplateDataContext(workspace workspace.WorkspaceBaseStruct, project applicationproject.ProjectBaseStruct, resource objectresource.ResourceBaseStruct, template codetemplate.TemplateBaseStruct, layer objectresource.Layer, section objectresource.Section) *CodeTemplateDataContext {
	templateService := objectResourceService.NewObjectResourceService(project.Specifications.Platform.Type)

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
