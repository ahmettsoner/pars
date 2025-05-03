package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func SubmitGroup(t *testing.T, group string, environment string) {
	commands := []string{"group", "submit", group}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveGroup(t *testing.T, group string, environment string) {
	commands := []string{"group", "remove", group}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func SubmitGroupFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"group", "submit", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveGroupFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"group", "remove", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
