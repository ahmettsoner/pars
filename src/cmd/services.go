package cmd

import (
	"parsdevkit.net/application"
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"

	group "parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	projectApplication "parsdevkit.net/modules/project/application_project"
	resourceData "parsdevkit.net/modules/resource/data_resource"
	resourceObject "parsdevkit.net/modules/resource/object_resource"

	groupGroup "parsdevkit.net/modules/group/basic_group"
	taskCommon "parsdevkit.net/modules/task/basic_task"
	templateCode "parsdevkit.net/modules/template/code_template"
	templateFile "parsdevkit.net/modules/template/file_template"
	templateShared "parsdevkit.net/modules/template/shared_template"
	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	commontask "parsdevkit.net/structs/task/common-task"
	codetemplate "parsdevkit.net/structs/template/code-template"
	filetemplate "parsdevkit.net/structs/template/file-template"
	sharedtemplate "parsdevkit.net/structs/template/shared-template"
	"parsdevkit.net/structs/workspace"

	group_payload "parsdevkit.net/modules/group/basic_group_payload"
	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"
	platformsCommon "parsdevkit.net/platforms/common"
	applicationProjectSchema "parsdevkit.net/structs/project/application-project"
	dataResourceSchema "parsdevkit.net/structs/resource/data-resource"
	objectResourceSchema "parsdevkit.net/structs/resource/object-resource"
	codeTemplateSchema "parsdevkit.net/structs/template/code-template"
	fileTemplateSchema "parsdevkit.net/structs/template/file-template"
	sharedTemplateSchema "parsdevkit.net/structs/template/shared-template"
)

func RegisterServices() {
	registerEngines()
	registerSchemas()
	registerContainers()
}
func registerSchemas() {
	schemas.Register(&group_payload.GroupBaseStruct{})
	schemas.Register(&applicationProjectSchema.ProjectBaseStruct{})
	schemas.Register(&dataResourceSchema.ResourceBaseStruct{})
	schemas.Register(&objectResourceSchema.ResourceBaseStruct{})
	schemas.Register(&codeTemplateSchema.TemplateBaseStruct{})
	schemas.Register(&fileTemplateSchema.TemplateBaseStruct{})
	schemas.Register(&sharedTemplateSchema.TemplateBaseStruct{})
}

func registerEngines() {
	engines.Register(&group.GroupEngine{})
	engines.Register(&projectApplication.ApplicationProjectEngine{})
	engines.Register(&resourceData.DataResourceEngine{})
	engines.Register(&resourceObject.ObjectResourceEngine{})
	engines.Register(&templateCode.CodeTemplateEngine{})
	engines.Register(&templateFile.FileTemplateEngine{})
	engines.Register(&templateShared.SharedTemplateEngine{})
	engines.Register(&taskCommon.CommonTaskEngine{})
}
func registerContainers() {
	dbContext := contexts.NewDbContext(application.GetEnvironment())

	ioc.Register(func() *repositories.WorkspaceRepository {
		return repositories.NewWorkspaceRepository(dbContext)
	})
	ioc.Register(func() *repositories.GroupRepository {
		return repositories.NewGroupRepository(dbContext)
	})
	ioc.Register(func() *repositories.ProjectRepository {
		return repositories.NewProjectRepository(dbContext)
	})
	ioc.Register(func() *repositories.ResourceRepository {
		return repositories.NewResourceRepository(dbContext)
	})
	ioc.Register(func() *repositories.TemplateRepository {
		return repositories.NewTemplateRepository(dbContext)
	})
	ioc.Register(func() *repositories.TaskRepository {
		return repositories.NewTaskRepository(dbContext)
	})
	ioc.Register(func() *repositories.SettingsRepository {
		return repositories.NewSettingsRepository(dbContext)
	})
	ioc.Register(func() *repositories.GenerationHistoryRepository {
		return repositories.NewGenerationHistoryRepository(dbContext)
	})
	ioc.RegisterInterface[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]](func() contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct] {
		return workspaceWorkspace.NewWorkspaceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[basic_group_contract.GroupInterface](func() basic_group_contract.GroupInterface {
		return groupGroup.NewGroupService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct]](func() contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct] {
		return projectApplication.NewApplicationProjectService(application.GetEnvironment(), platformsCommon.Registry)
	})
	ioc.RegisterInterface[contracts.TemplateServiceInterface[codetemplate.TemplateBaseStruct]](func() contracts.TemplateServiceInterface[codetemplate.TemplateBaseStruct] {
		return templateCode.NewCodeTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct]](func() contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct] {
		return templateFile.NewFileTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.TemplateServiceInterface[sharedtemplate.TemplateBaseStruct]](func() contracts.TemplateServiceInterface[sharedtemplate.TemplateBaseStruct] {
		return templateShared.NewSharedTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.ResourceServiceInterface[objectresource.ResourceBaseStruct]](func() contracts.ResourceServiceInterface[objectresource.ResourceBaseStruct] {
		return resourceObject.NewObjectResourceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.ResourceServiceInterface[dataresource.ResourceBaseStruct]](func() contracts.ResourceServiceInterface[dataresource.ResourceBaseStruct] {
		return resourceData.NewDataResourceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.TaskServiceInterface[commontask.TaskBaseStruct]](func() contracts.TaskServiceInterface[commontask.TaskBaseStruct] {
		return taskCommon.NewCommonTaskService(application.GetEnvironment())
	})

}
