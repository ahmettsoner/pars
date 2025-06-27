package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewGroup(t *testing.T, group string, environment string) {
	commands := []string{"group", "new", group}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveGroup(t *testing.T, group string, environment string) {
	commands := []string{"group", "remove", group}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
