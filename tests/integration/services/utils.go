package services

import (
	"testing"

	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/application/models/label"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/models"
	group "parsdevkit.net/modules/group/basic_group"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/platforms/dotnet/managers"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"parsdevkit.net/application/schemas"
)

func InitializeNewWorkspace(t *testing.T, wsPath, workspaceName, environment string) basic_workspace_payload_structs.WorkspaceBaseStruct {

	workspace := basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			workspaceName,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		basic_workspace_payload_structs.NewWorkspaceSpecification(0, workspaceName, wsPath),
	)

	workspaceService := workspaceWorkspace.NewWorkspaceService(environment)
	tempWorkspace, err := workspaceService.Save(workspace)
	require.NoError(t, err, "Failed to save workspace")
	assert.Equal(t, workspace, *tempWorkspace)

	return workspace
}

func RemoveWorkspace(t *testing.T, workspaceName, environment string) {
	workspaceService := workspaceWorkspace.NewWorkspaceService(environment)
	_, err := workspaceService.Remove(workspaceName, true, true)
	require.NoError(t, err, "Failed to delete workspace")
}
func CreateGroup(t *testing.T, groupName, path, environment string) basic_group_payload_structs.GroupBaseStruct {

	groupStruct := *BasicGroup_WithNamePath(groupName, path)

	groupService := group.NewGroupService(environment)
	tempGroup, err := groupService.Save(groupStruct)
	require.NoError(t, err, "Failed to save group")
	assert.Equal(t, groupStruct, *tempGroup)

	return groupStruct
}
func RemoveGroup(t *testing.T, groupName, environment string) {
	groupService := group.NewGroupService(environment)
	_, err := groupService.Remove(groupName, true)
	require.NoError(t, err, "Failed to delete group")
}

func CreateNewTestProject(t *testing.T, name, wsPath, workspaceName string) application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, name, schemas.Metadata{}),
		application_project_payload_structs.NewProjectSpecification(
			0,
			name,
			"",
			workspaceName,
			applicationGroup.GroupIdentifier{},
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(name),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
			application_project_payload_structs.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
		),
		application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.Runtime{},
			application_project_payload_structs.Schema{},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithLayer(t *testing.T, name, wsPath, workspaceName string, layers []applicationProject.Layer) application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, name, schemas.Metadata{}),
		application_project_payload_structs.NewProjectSpecification(
			0,
			name,
			"",
			workspaceName,
			applicationGroup.GroupIdentifier{},
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(name),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
			application_project_payload_structs.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			layers,
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
		),
		application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.Runtime{},
			application_project_payload_structs.Schema{},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectWithGroup(t *testing.T, name, wsPath, workspaceName, groupName, groupPath string) application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, name, schemas.Metadata{}),
		application_project_payload_structs.NewProjectSpecification(
			0,
			name,
			"",
			workspaceName,
			applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(name),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
			application_project_payload_structs.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
		),
		application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.Runtime{},
			application_project_payload_structs.Schema{},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectWithGroupAndLayers(t *testing.T, name, wsPath, workspaceName, groupName, groupPath string, layers []applicationProject.Layer) application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, name, schemas.Metadata{}),
		application_project_payload_structs.NewProjectSpecification(
			0,
			name,
			groupName,
			workspaceName,
			applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(name),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
			application_project_payload_structs.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			layers,
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
		),
		application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.Runtime{},
			application_project_payload_structs.Schema{},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndPath(t *testing.T, name, path, wsPath, workspaceName, groupName, groupPath string) application_project_payload_structs.ProjectBaseStruct {

	project := application_project_payload_structs.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, application_project_payload_structs.PROJECT_KIND, name, schemas.Metadata{}),
		application_project_payload_structs.NewProjectSpecification(
			0,
			name,
			groupName,
			workspaceName,
			applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(path),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
			application_project_payload_structs.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			[]applicationProject.Layer(nil),
			[]applicationProject.Dependency(nil),
			[]application_project_payload_structs.ProjectBaseStruct(nil),
		),
		application_project_payload_structs.NewApplication(
			models.ProjectTypes.Library,
			application_project_payload_structs.Runtime{},
			application_project_payload_structs.Schema{},
			[]string(nil),
			[]string(nil),
			[]string(nil),
			[]string(nil),
			"",
			[]string(nil),
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func GetDependencies(index int, count int, withVersion bool) []applicationProject.Dependency {
	packages := []applicationProject.Dependency{
		applicationProject.NewDependency("Microsoft.Extensions.DependencyInjection", "8.0.0"),
		applicationProject.NewDependency("Microsoft.Extensions.Logging", "8.0.0"),
		applicationProject.NewDependency("Microsoft.EntityFrameworkCore.Design", "8.0.2"),
		applicationProject.NewDependency("Microsoft.EntityFrameworkCore.InMemory", "8.0.2"),
		applicationProject.NewDependency("Microsoft.EntityFrameworkCore.Sqlite", "8.0.2"),
		applicationProject.NewDependency("Newtonsoft.Json", "13.0.1"),
		applicationProject.NewDependency("AutoMapper", "11.0.0"),
		applicationProject.NewDependency("FluentValidation", "11.1.0"),
		applicationProject.NewDependency("Moq", "4.16.1"),
		applicationProject.NewDependency("Hangfire", "1.7.22"),
		applicationProject.NewDependency("Serilog", "2.10.0"),
	}

	if count <= 0 || count > len(packages) {
		return nil
	}

	var selectedElements []applicationProject.Dependency
	for _, _package := range packages[index : index+count] {
		packageVersion := ""
		if withVersion {
			packageVersion = _package.Version
		}

		selectedDependency := applicationProject.NewDependency(_package.Name, packageVersion)
		selectedElements = append(selectedElements, selectedDependency)
	}

	return selectedElements
}

func BasicGroup_WithNamePath(name, path string) *basic_group_payload_structs.GroupBaseStruct {

	group := basic_group_payload_structs.NewGroupBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Group,
			"",
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		applicationGroup.NewGroupSpecification(0,
			name,
			path,
			[]string{"foo", "bar"},
		),
	)

	return &group
}
