package common

import (
	"testing"

	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"parsdevkit.net/application/schemas"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

func InitializeNewWorkspace(commander CommanderType, t *testing.T, wsPath, workspaceName, environment string) {
	commands := []string{"init", workspaceName, wsPath}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
func InitializeNewWorkspaceWithService(t *testing.T, wsPath, workspaceName, environment string) basic_workspace_payload_structs.WorkspaceBaseStruct {

	workspace := basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			workspaceName,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		applicationWorkspace.NewWorkspaceSpecification(0, workspaceName, wsPath),
	)

	workspaceService := workspaceWorkspace.NewWorkspaceService(environment)
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

	_, err := ExecuteCommandWithSelector(CommanderTypes.GO, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
func RemoveWorkspaceWithService(t *testing.T, workspaceName, environment string) {
	workspaceService := workspaceWorkspace.NewWorkspaceService(environment)
	_, err := workspaceService.Remove(workspaceName, true, true)
	require.NoError(t, err, "Failed to delete workspace")
}
