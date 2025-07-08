package cmd

import (
	"parsdevkit.net/application"
	"parsdevkit.net/application/bus"
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/platforms"
	"parsdevkit.net/application/schemas"
	object_resource_handler "parsdevkit.net/modules/resource/object_resource/handlers"
	object_resource_payload_commands "parsdevkit.net/modules/resource/object_resource_payload/commands"
	object_resource_payload_events "parsdevkit.net/modules/resource/object_resource_payload/events"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/template/file_template_contract"
	"parsdevkit.net/modules/template/shared_template_contract"

	group "parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_contract"
	projectApplication "parsdevkit.net/modules/project/application_project"
	"parsdevkit.net/modules/project/application_project_contract"
	resourceData "parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/resource/data_resource_contract"
	resourceObject "parsdevkit.net/modules/resource/object_resource"
	"parsdevkit.net/modules/resource/object_resource_contract"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	groupGroup "parsdevkit.net/modules/group/basic_group"
	taskCommon "parsdevkit.net/modules/task/basic_task"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
	templateCode "parsdevkit.net/modules/template/code_template"
	templateFile "parsdevkit.net/modules/template/file_template"
	templateShared "parsdevkit.net/modules/template/shared_template"
	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"
	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"

	application_project_handlers "parsdevkit.net/modules/project/application_project/handlers"
	application_project_payload_commands "parsdevkit.net/modules/project/application_project_payload/commands"
	angularManager "parsdevkit.net/platforms/angular/managers"
	dotnetManager "parsdevkit.net/platforms/dotnet/managers"
	goManager "parsdevkit.net/platforms/go/managers"
	nodejsManager "parsdevkit.net/platforms/nodejs/managers"
	parsManager "parsdevkit.net/platforms/pars/managers"
)

func RegisterServices() {
	registerCommandHandlers()
	registerEventHandlers()
	registerEngines()
	registerSchemas()
	registerPlatformManager()
	registerContainers()

	// Send command
	// err := bus.SendCommand(application_project_payload_commands.CreateApplicationProject{})
	// if err != nil {
	// 	panic(err)
	// }

	// // Publish event
	// bus.PublishEvent(application_project_payload_events.CreateApplicationProjectCreated{})
}
func registerCommandHandlers() {
	bus.RegisterCommandHandler[application_project_payload_commands.CreateApplicationProject](&application_project_handlers.CreateApplicationProjectHandler{})
	bus.RegisterCommandHandler[object_resource_payload_commands.GenerateResourceContents](&object_resource_handler.GenerateResourceContentsCommandHandler{})
}
func registerEventHandlers() {
	bus.RegisterEventHandler[object_resource_payload_events.ResourceCreated](&object_resource_handler.ResourceCreatedEventHandler{})
}

func registerSchemas() {
	schemas.Register(&basic_group_payload_structs.GroupBaseStruct{})
	schemas.Register(&application_project_payload_structs.ProjectBaseStruct{})
	schemas.Register(&data_resource_payload_structs.ResourceBaseStruct{})
	schemas.Register(&object_resource_payload_structs.ResourceBaseStruct{})
	schemas.Register(&code_template_payload_structs.TemplateBaseStruct{})
	schemas.Register(&file_template_payload_structs.TemplateBaseStruct{})
	schemas.Register(&shared_template_payload_structs.TemplateBaseStruct{})
}

func registerEngines() {
	engines.Register(&group.GroupEngine{})
	engines.Register(&projectApplication.ApplicationProjectEngine{})
	engines.Register(&resourceData.DataResourceEngine{})
	engines.Register(&resourceObject.ObjectResourceEngine{})
	engines.Register(&templateCode.CodeTemplateEngine{})
	engines.Register(&templateFile.FileTemplateEngine{})
	engines.Register(&templateShared.SharedTemplateEngine{})
	engines.Register(&taskCommon.BasicTaskEngine{})
}
func registerPlatformManager() {
	platforms.Register(parsManager.NewParsManager())
	platforms.Register(dotnetManager.NewDotnetManager())
	platforms.Register(angularManager.NewAngularManager())
	platforms.Register(nodejsManager.NewNodeJSManager())
	platforms.Register(goManager.NewGoManager())
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
	ioc.RegisterInterface[basic_workspace_contract.WorkspaceInterface](func() basic_workspace_contract.WorkspaceInterface {
		return workspaceWorkspace.NewWorkspaceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[basic_group_contract.GroupInterface](func() basic_group_contract.GroupInterface {
		return groupGroup.NewGroupService(application.GetEnvironment())
	})
	ioc.RegisterInterface[application_project_contract.ProjectInterface](func() application_project_contract.ProjectInterface {
		return projectApplication.NewApplicationProjectService(application.GetEnvironment())
	})
	ioc.RegisterInterface[code_template_contract.TemplateInterface](func() code_template_contract.TemplateInterface {
		return templateCode.NewCodeTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[file_template_contract.TemplateInterface](func() file_template_contract.TemplateInterface {
		return templateFile.NewFileTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[shared_template_contract.TemplateInterface](func() shared_template_contract.TemplateInterface {
		return templateShared.NewSharedTemplateService(application.GetEnvironment())
	})
	ioc.RegisterInterface[data_resource_contract.ResourceInterface](func() data_resource_contract.ResourceInterface {
		return resourceData.NewDataResourceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[object_resource_contract.ResourceInterface](func() object_resource_contract.ResourceInterface {
		return resourceObject.NewObjectResourceService(application.GetEnvironment())
	})
	ioc.RegisterInterface[contracts.TaskServiceInterface[basic_task_payload_structs.TaskBaseStruct]](func() contracts.TaskServiceInterface[basic_task_payload_structs.TaskBaseStruct] {
		return taskCommon.NewBasicTaskService(application.GetEnvironment())
	})

}
