package net8

import (
	"os"
	"path/filepath"
	"testing"

	"parsdevkit.net/models"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/label"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/platforms/dotnet/managers"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"parsdevkit.net/core/utils"

	"github.com/stretchr/testify/require"
)

func InitializeNewWorkspace(t *testing.T, testPath, workspaceName, environment string) {

	workspace := workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName))

	err := os.Mkdir(workspace.GetAbsolutePath(), os.ModePerm)
	require.NoError(t, err)

	err = os.Mkdir(workspace.GetCodeBaseFolder(), os.ModePerm)
	require.NoError(t, err)

	err = os.Mkdir(workspace.GetTemplatesFolder(), os.ModePerm)
	require.NoError(t, err)

	err = os.Mkdir(workspace.GetResourcesFolder(), os.ModePerm)
	require.NoError(t, err)
}

func RemoveWorkspace(t *testing.T, workspaceName, environment string) {
	// ExecuteCommand(t, environment, "remove", "workspace", workspaceName)
}

func CreateNewTestProject(t *testing.T, name, testPath, workspaceName string) applicationproject.ProjectSpecification {

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
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
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

func CreateNewTestProjectWithLayer(t *testing.T, name, testPath, workspaceName string, layers []applicationproject.Layer) applicationproject.ProjectSpecification {

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
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
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
func CreateNewTestProjectWithGroup(t *testing.T, name, testPath, workspaceName, groupName, groupPath string) applicationproject.ProjectSpecification {

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
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateProject(project)
	require.NoError(t, err)

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndLayers(t *testing.T, name, testPath, workspaceName, groupName, groupPath string, layers []applicationproject.Layer) applicationproject.ProjectSpecification {

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
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
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

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndPath(t *testing.T, name, path, testPath, workspaceName, groupName, groupPath string) applicationproject.ProjectSpecification {

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
		utils.PathToArray(path),
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateProject(project)
	require.NoError(t, err)

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectGroupAndPath(t *testing.T, name, path, testPath, workspaceName string) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		group.NewGroupSpecification(0, name, name, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		utils.PathToArray(path),
		workspace.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName)),
		applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
		applicationproject.Runtime{},
		applicationproject.Schema{},
		applicationproject.Configuration{},
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateGroup(project)
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
