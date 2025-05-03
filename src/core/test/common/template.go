package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func SubmitTemplateFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"template", "submit", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveTemplateFromFile(t *testing.T, declarationFile, environment string) {
	commands := []string{"template", "remove", "-f", declarationFile}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
