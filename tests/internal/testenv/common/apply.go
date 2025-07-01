package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Apply(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"apply", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
