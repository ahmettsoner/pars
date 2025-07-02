package cmd

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"

	group "parsdevkit.net/modules/group/group"
	projectApplication "parsdevkit.net/modules/project/application"
	resourceData "parsdevkit.net/modules/resource/data"
	resourceObject "parsdevkit.net/modules/resource/object"
	taskCommon "parsdevkit.net/modules/task/common"
	templateCode "parsdevkit.net/modules/template/code"
	templateFile "parsdevkit.net/modules/template/file"
	templateShared "parsdevkit.net/modules/template/shared"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	groupSchema "parsdevkit.net/modules/group/group/structs"
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
	schemas.Register(&groupSchema.GroupBaseStruct{})
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
	dbContext := contexts.NewDbContext(utils.GetEnvironment())

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
	ioc.RegisterInterface[contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct]](func() contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct] {
		return projectApplication.NewApplicationProjectService(utils.GetEnvironment(), platformsCommon.Registry)
	})
	ioc.RegisterInterface[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]](func() contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct] {
		return services.NewWorkspaceService(utils.GetEnvironment())
	})

}
