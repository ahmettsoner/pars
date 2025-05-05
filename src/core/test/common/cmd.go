package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"parsdevkit.net/cmd"
	"parsdevkit.net/core/utils"

	"github.com/stretchr/testify/require"
)

type CommanderType int

var CommanderTypes = struct {
	GO    CommanderType
	Cobra CommanderType
	Pars  CommanderType
}{
	GO:    1,
	Cobra: 2,
	Pars:  3,
}

func ExecuteCommandWithSelector(commander CommanderType, t *testing.T, environment string, commands ...string) ([]string, error) {
	return ExecuteCommandWithSelectorOnPath(commander, t, environment, utils.GetSourceLocation(), commands...)
}
func ExecuteCommandWithSelectorOnPath(commander CommanderType, t *testing.T, environment string, path string, commands ...string) ([]string, error) {
	switch commander {
	case CommanderTypes.GO:
		return ExecuteCommandOnGoOnPath(t, environment, path, commands...)
	case CommanderTypes.Cobra:
		return ExecuteCommandOnCobraOnPath(t, environment, path, commands...)
	case CommanderTypes.Pars:
		return ExecuteCommandOnParsOnPath(t, environment, path, commands...)
	default:
		return ExecuteCommandOnParsOnPath(t, environment, path, commands...)
	}
}

func ExecuteCommand(t *testing.T, environment string, commands ...string) ([]string, error) {
	return ExecuteCommandOnPath(t, environment, utils.GetSourceLocation(), commands...)

}

func ExecuteCommandOnPath(t *testing.T, environment string, path string, commands ...string) ([]string, error) {
	return ExecuteCommandWithSelectorOnPath(CommanderTypes.GO, t, environment, path, commands...)

}

func ExecuteCommandOnCobra(t *testing.T, environment string, commands ...string) ([]string, error) {
	return ExecuteCommandOnCobraOnPath(t, environment, utils.GetSourceLocation(), commands...)
}
func ExecuteCommandOnCobraOnPath(t *testing.T, environment string, path string, commands ...string) ([]string, error) {

	fullCommand := []string{}
	fullCommand = append(fullCommand, commands...)
	fullCommand = append(fullCommand, []string{"--log-level", "verbose", "--env", environment}...)

	t.Logf("executing command %v", fullCommand)
	commander := *(cmd.RootCmd)
	commander.SetArgs(fullCommand)

	existingWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if err := os.Chdir(path); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	err = commander.Execute()
	require.NoErrorf(t, err, "Failed to execute command %v", fullCommand)

	if err := os.Chdir(existingWD); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	return fullCommand, err
}

func ExecuteCommandOnGo(t *testing.T, environment string, commands ...string) ([]string, error) {
	return ExecuteCommandOnGoOnPath(t, environment, utils.GetSourceLocation(), commands...)
}
func ExecuteCommandOnGoOnPath(t *testing.T, environment string, path string, commands ...string) ([]string, error) {
	fullCommand := []string{"run", filepath.Join(utils.GetSourceLocation(), "pars.go")}
	fullCommand = append(fullCommand, commands...)
	fullCommand = append(fullCommand, []string{"--log-level", "verbose", "--env", environment}...)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	t.Logf("executing command %v", fullCommand[2:])
	cmd := exec.Command("go", fullCommand...)
	cmd.Dir = path

	output, err := cmd.CombinedOutput()
	t.Logf(">>> %v", string(output))

	return fullCommand[2:], err
}

func ExecuteCommandOnPars(t *testing.T, environment string, commands ...string) ([]string, error) {
	return ExecuteCommandOnParsOnPath(t, environment, utils.GetSourceLocation(), commands...)
}
func ExecuteCommandOnParsOnPath(t *testing.T, environment string, path string, commands ...string) ([]string, error) {
	fullCommand := []string{}
	fullCommand = append(fullCommand, commands...)
	fullCommand = append(fullCommand, []string{"--log-level", "verbose", "--env", environment}...)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	t.Logf("executing command %v", fullCommand[2:])
	cmd := exec.Command("pars", fullCommand...)
	cmd.Dir = path

	output, err := cmd.CombinedOutput()
	t.Logf(">>> %v", string(output))

	return fullCommand[2:], err
}
