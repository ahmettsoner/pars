package models

import (
	"parsdevkit.net/application/ioc"

	"parsdevkit.net/components/template"
	"parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/resource/object_resource"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/template/code_template"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
)

type CodeTemplateDataContext struct {
	Workspace basic_workspace.WorkspaceComposite
	Group     basic_group.GroupComposite
	Project   application_project.ApplicationProjectComposite
	Resource  object_resource.ObjectResourceComposite
	Template  code_template.TemplateComposite
	Layer     object_resource.ObjectLayerComposite
	Section   object_resource.ObjectSectionComposite
}

func NewCodeTemplateDataContext(source template.ContextProviderSource) *CodeTemplateDataContext {

	workspaceContextProvider := ioc.Get[basic_workspace_contract.ContextProviderInterface]()
	groupContextProvider := ioc.Get[basic_group_contract.ContextProviderInterface]()
	projectContextProvider := ioc.Get[application_project_contract.ContextProviderInterface]()
	resourceContextProvider := ioc.Get[object_resource_contract.ContextProviderInterface]()
	templateContextProvider := ioc.Get[code_template_contract.ContextProviderInterface]()

	return &CodeTemplateDataContext{
		Workspace: workspaceContextProvider.Context(source).(basic_workspace.WorkspaceComposite),
		Group:     groupContextProvider.Context(source).(basic_group.GroupComposite),
		Project:   projectContextProvider.Context(source).(application_project.ApplicationProjectComposite),
		Resource:  resourceContextProvider.Context(source).(object_resource.ObjectResourceComposite),
		Template:  templateContextProvider.Context(source).(code_template.TemplateComposite),
		Layer:     resourceContextProvider.LayerToModelContext(source).(object_resource.ObjectLayerComposite),
		Section:   resourceContextProvider.SectionToModelContext(source).(object_resource.ObjectSectionComposite),
	}
}
