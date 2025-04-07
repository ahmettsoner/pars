package services

import (
	"testing"

	"parsdevkit.net/models"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/label"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/platforms/dotnet/managers"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"parsdevkit.net/core/utils"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func InitializeNewWorkspace(t *testing.T, wsPath, workspaceName, environment string) workspace.WorkspaceBaseStruct {

	workspace := workspace.NewWorkspaceBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Workspace,
			workspaceName,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
	)

	workspaceService := *services.NewWorkspaceService(environment)
	tempWorkspace, err := workspaceService.Save(workspace)
	require.NoError(t, err, "Failed to save workspace")
	assert.Equal(t, workspace, *tempWorkspace)

	return workspace
}

func RemoveWorkspace(t *testing.T, workspaceName, environment string) {
	workspaceService := *services.NewWorkspaceService(environment)
	_, err := workspaceService.Remove(workspaceName, true, true)
	require.NoError(t, err, "Failed to delete workspace")
}
func CreateGroup(t *testing.T, groupName, path, environment string) group.GroupBaseStruct {

	group := *BasicGroup_WithNamePath(groupName, path)

	groupService := *services.NewGroupService(environment)
	tempGroup, err := groupService.Save(group)
	require.NoError(t, err, "Failed to save group")
	assert.Equal(t, group, *tempGroup)

	return group
}
func RemoveGroup(t *testing.T, groupName, environment string) {
	groupService := *services.NewGroupService(environment)
	_, err := groupService.Remove(groupName, true)
	require.NoError(t, err, "Failed to delete group")
}

func CreateNewTestProject(t *testing.T, name, wsPath, workspaceName string) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		group.GroupSpecification{},
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(name),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithLayer(t *testing.T, name, wsPath, workspaceName string, layers []applicationproject.Layer) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		group.GroupSpecification{},
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(name),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{
			Layers: layers,
		},
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectWithGroup(t *testing.T, name, wsPath, workspaceName, groupName, groupPath string) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		group.NewGroupSpecification(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(name),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectWithGroupAndLayers(t *testing.T, name, wsPath, workspaceName, groupName, groupPath string, layers []applicationproject.Layer) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		groupName,
		workspaceName,
		models.ProjectTypes.Library,
		group.NewGroupSpecification(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(name),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{
			Layers: layers,
		},
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndPath(t *testing.T, name, path, wsPath, workspaceName, groupName, groupPath string) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		groupName,
		workspaceName,
		models.ProjectTypes.Library,
		group.NewGroupSpecification(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(path),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func GetPackages(index int, count int, withVersion bool) []applicationproject.Package {
	packages := []applicationproject.Package{
		applicationproject.NewPackage("Microsoft.Extensions.DependencyInjection", "8.0.0"),
		applicationproject.NewPackage("Microsoft.Extensions.Logging", "8.0.0"),
		applicationproject.NewPackage("Microsoft.EntityFrameworkCore.Design", "8.0.2"),
		applicationproject.NewPackage("Microsoft.EntityFrameworkCore.InMemory", "8.0.2"),
		applicationproject.NewPackage("Microsoft.EntityFrameworkCore.Sqlite", "8.0.2"),
		applicationproject.NewPackage("Newtonsoft.Json", "13.0.1"),
		applicationproject.NewPackage("AutoMapper", "11.0.0"),
		applicationproject.NewPackage("FluentValidation", "11.1.0"),
		applicationproject.NewPackage("Moq", "4.16.1"),
		applicationproject.NewPackage("Hangfire", "1.7.22"),
		applicationproject.NewPackage("Serilog", "2.10.0"),
	}

	if count <= 0 || count > len(packages) {
		return nil
	}

	var selectedElements []applicationproject.Package
	for _, _package := range packages[index : index+count] {
		packageVersion := ""
		if withVersion {
			packageVersion = _package.Version
		}

		selectedPackage := applicationproject.NewPackage(_package.Name, packageVersion)
		selectedElements = append(selectedElements, selectedPackage)
	}

	return selectedElements
}

func BasicGroup_WithNamePath(name, path string) *group.GroupBaseStruct {

	group := group.NewGroupBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Group,
			name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		group.NewGroupSpecification(0,
			name,
			path,
			[]string{"foo", "bar"},
		),
	)

	return &group
}
