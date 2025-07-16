package models

import (
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/components/template"
	"parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/object_resource_contract"

	"parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/template/file_template"
	"parsdevkit.net/modules/template/file_template_contract"
	"parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
)

type FileTemplateDataContext struct {
	Workspace basic_workspace.WorkspaceComposite
	Group     basic_group.GroupComposite
	Project   application_project.ApplicationProjectComposite
	Resource  data_resource.DataResourceComposite
	Template  file_template.TemplateComposite
	Layer     data_resource.DataLayerComposite
	Section   data_resource.DataSectionComposite
}

func NewFileTemplateDataContext(source template.ContextProviderSource) *FileTemplateDataContext {

	workspaceContextProvider := ioc.Get[basic_workspace_contract.ContextProviderInterface]()
	groupContextProvider := ioc.Get[basic_group_contract.ContextProviderInterface]()
	projectContextProvider := ioc.Get[application_project_contract.ContextProviderInterface]()
	resourceContextProvider := ioc.Get[object_resource_contract.ContextProviderInterface]()
	templateContextProvider := ioc.Get[file_template_contract.ContextProviderInterface]()

	return &FileTemplateDataContext{
		Workspace: workspaceContextProvider.Context(source).(basic_workspace.WorkspaceComposite),
		Group:     groupContextProvider.Context(source).(basic_group.GroupComposite),
		Project:   projectContextProvider.Context(source).(application_project.ApplicationProjectComposite),
		Resource:  resourceContextProvider.Context(source).(data_resource.DataResourceComposite),
		Template:  templateContextProvider.Context(source).(file_template.TemplateComposite),
		Layer:     resourceContextProvider.LayerToModelContext(source).(data_resource.DataLayerComposite),
		Section:   resourceContextProvider.SectionToModelContext(source).(data_resource.DataSectionComposite),
	}
}
