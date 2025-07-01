package services

import (
	"testing"

	"parsdevkit.net/core/utilities/file"

	"parsdevkit.net/application/models/label"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/models"
	group "parsdevkit.net/modules/group/group"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/platforms/dotnet/managers"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"parsdevkit.net/application/schemas"
)

func InitializeNewWorkspace(t *testing.T, wsPath, workspaceName, environment string) workspace.WorkspaceBaseStruct {

	workspace := workspace.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			workspaceName,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		workspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
	)

	workspaceService := services.NewWorkspaceService(environment)
	tempWorkspace, err := workspaceService.Save(workspace)
	require.NoError(t, err, "Failed to save workspace")
	assert.Equal(t, workspace, *tempWorkspace)

	return workspace
}

func RemoveWorkspace(t *testing.T, workspaceName, environment string) {
	workspaceService := services.NewWorkspaceService(environment)
	_, err := workspaceService.Remove(workspaceName, true, true)
	require.NoError(t, err, "Failed to delete workspace")
}
func CreateGroup(t *testing.T, groupName, path, environment string) group.GroupBaseStruct {

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

func CreateNewTestProject(t *testing.T, name, wsPath, workspaceName string) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		applicationGroup.GroupIdentifier{},
		"",
		[]string(nil),
		[]label.Label(nil),
		file.PathToArray(name),
		applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
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

func CreateNewTestProjectWithLayer(t *testing.T, name, wsPath, workspaceName string, layers []applicationProject.Layer) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		"",
		workspaceName,
		models.ProjectTypes.Library,
		applicationGroup.GroupIdentifier{},
		"",
		[]string(nil),
		[]label.Label(nil),
		file.PathToArray(name),
		applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
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
		applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		file.PathToArray(name),
		applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
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
func CreateNewTestProjectWithGroupAndLayers(t *testing.T, name, wsPath, workspaceName, groupName, groupPath string, layers []applicationProject.Layer) applicationproject.ProjectSpecification {

	project := applicationproject.NewProjectSpecification(
		0,
		name,
		groupName,
		workspaceName,
		models.ProjectTypes.Library,
		applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		file.PathToArray(name),
		applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
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
		applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
		"",
		[]string(nil),
		[]label.Label(nil),
		file.PathToArray(path),
		applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, wsPath),
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

func GetPackages(index int, count int, withVersion bool) []applicationProject.Package {
	packages := []applicationProject.Package{
		applicationProject.NewPackage("Microsoft.Extensions.DependencyInjection", "8.0.0"),
		applicationProject.NewPackage("Microsoft.Extensions.Logging", "8.0.0"),
		applicationProject.NewPackage("Microsoft.EntityFrameworkCore.Design", "8.0.2"),
		applicationProject.NewPackage("Microsoft.EntityFrameworkCore.InMemory", "8.0.2"),
		applicationProject.NewPackage("Microsoft.EntityFrameworkCore.Sqlite", "8.0.2"),
		applicationProject.NewPackage("Newtonsoft.Json", "13.0.1"),
		applicationProject.NewPackage("AutoMapper", "11.0.0"),
		applicationProject.NewPackage("FluentValidation", "11.1.0"),
		applicationProject.NewPackage("Moq", "4.16.1"),
		applicationProject.NewPackage("Hangfire", "1.7.22"),
		applicationProject.NewPackage("Serilog", "2.10.0"),
	}

	if count <= 0 || count > len(packages) {
		return nil
	}

	var selectedElements []applicationProject.Package
	for _, _package := range packages[index : index+count] {
		packageVersion := ""
		if withVersion {
			packageVersion = _package.Version
		}

		selectedPackage := applicationProject.NewPackage(_package.Name, packageVersion)
		selectedElements = append(selectedElements, selectedPackage)
	}

	return selectedElements
}

func BasicGroup_WithNamePath(name, path string) *group.GroupBaseStruct {

	group := group.NewGroupBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Group,
			"",
			name,
			schemas.Metadata{
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
