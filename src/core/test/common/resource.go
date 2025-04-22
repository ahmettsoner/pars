package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func SubmitResourceFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"resource", "submit", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveResourceFromFile(t *testing.T, declarationFile, environment string) {
	commands := []string{"resource", "remove", "-f", declarationFile}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
