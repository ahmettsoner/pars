package net8

import (
	"os"
	"path/filepath"
	"testing"

	"parsdevkit.net/application/models/label"
	applicationGroup "parsdevkit.net/application/structs/group"
	"parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/platforms/dotnet/managers"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	"github.com/stretchr/testify/require"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"
	"parsdevkit.net/pkg/utilities/file"
)

func InitializeNewWorkspace(t *testing.T, testPath, workspaceName, environment string) {

	workspace := basic_workspace_payload.NewWorkspaceSpecification(0, workspaceName, filepath.Join(testPath, workspaceName))

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

func CreateNewTestProject(t *testing.T, name, testPath, workspaceName string) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
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
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{},
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithLayer(t *testing.T, name, testPath, workspaceName string, layers []project.Layer) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
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
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{
				Layers: layers,
			},
		),
	)
	manager := managers.NewDotnetManager()
	err := manager.CreateProject(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectWithGroup(t *testing.T, name, testPath, workspaceName, groupName, groupPath string) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
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
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{},
		),
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateProject(project)
	require.NoError(t, err)

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndLayers(t *testing.T, name, testPath, workspaceName, groupName, groupPath string, layers []project.Layer) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
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
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{
				Layers: layers,
			},
		),
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateProject(project)
	require.NoError(t, err)

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}

func CreateNewTestProjectWithGroupAndPath(t *testing.T, name, path, testPath, workspaceName, groupName, groupPath string) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
			0,
			name,
			"",
			workspaceName,
			models.ProjectTypes.Library,
			applicationGroup.NewGroupIdentifier(0, groupName, groupPath, []string(nil)),
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(path),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{},
		),
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateProject(project)
	require.NoError(t, err)

	err = manager.AddToGroup(project)
	require.NoError(t, err)

	return project
}
func CreateNewTestProjectGroupAndPath(t *testing.T, name, path, testPath, workspaceName string) applicationproject.ProjectBaseStruct {

	project := applicationproject.NewProjectBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Project, applicationproject.PROJECT_KIND, name, schemas.Metadata{}),
		applicationproject.NewProjectSpecification(
			0,
			name,
			"",
			workspaceName,
			models.ProjectTypes.Library,
			applicationGroup.NewGroupIdentifier(0, name, name, []string(nil)),
			"",
			[]string(nil),
			[]label.Label(nil),
			file.PathToArray(path),
			applicationWorkspace.NewWorkspaceIdentifier(0, workspaceName, filepath.Join(testPath, workspaceName)),
			applicationproject.NewPlatform(models.PlatformTypes.Dotnet, dotnetModels.DotnetPlatformVersions.Net8.String()),
			applicationproject.Runtime{},
			applicationproject.Schema{},
			applicationproject.Configuration{},
		),
	)
	manager := managers.NewDotnetManager()

	err := manager.CreateGroup(project)
	require.NoError(t, err)

	return project
}

func GetDependencies(index int, count int, withVersion bool) []project.Dependency {
	packages := []project.Dependency{
		project.NewDependency("Microsoft.Extensions.DependencyInjection", "8.0.0"),
		project.NewDependency("Microsoft.Extensions.Logging", "8.0.0"),
		project.NewDependency("Microsoft.EntityFrameworkCore.Design", "8.0.2"),
		project.NewDependency("Microsoft.EntityFrameworkCore.InMemory", "8.0.2"),
		project.NewDependency("Microsoft.EntityFrameworkCore.Sqlite", "8.0.2"),
		project.NewDependency("Newtonsoft.Json", "13.0.1"),
		project.NewDependency("AutoMapper", "11.0.0"),
		project.NewDependency("FluentValidation", "11.1.0"),
		project.NewDependency("Moq", "4.16.1"),
		project.NewDependency("Hangfire", "1.7.22"),
		project.NewDependency("Serilog", "2.10.0"),
	}

	if count <= 0 || count > len(packages) {
		return nil
	}

	var selectedElements []project.Dependency
	for _, _package := range packages[index : index+count] {
		packageVersion := ""
		if withVersion {
			packageVersion = _package.Version
		}

		selectedDependency := project.NewDependency(_package.Name, packageVersion)
		selectedElements = append(selectedElements, selectedDependency)
	}

	return selectedElements
}
