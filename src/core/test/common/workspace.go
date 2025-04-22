package common

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func InitializeNewWorkspace(t *testing.T, wsPath, workspaceName, environment string) {
	commands := []string{"init", workspaceName, wsPath}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
func InitializeNewWorkspaceWithService(t *testing.T, wsPath, workspaceName, environment string) workspace.WorkspaceBaseStruct {

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

func SwitchToWorkspace(t *testing.T, workspaceName, environment string) {
	commands := []string{"workspace", "--switch", workspaceName}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveWorkspace(t *testing.T, workspaceName, environment string) {
	commands := []string{"workspace", "remove", workspaceName}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
func RemoveWorkspaceWithService(t *testing.T, workspaceName, environment string) {
	workspaceService := *services.NewWorkspaceService(environment)
	_, err := workspaceService.Remove(workspaceName, true, true)
	require.NoError(t, err, "Failed to delete workspace")
}
