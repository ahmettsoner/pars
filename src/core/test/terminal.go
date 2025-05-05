package test

import (
	"os/exec"
	"testing"

	"parsdevkit.net/core/utils"
)

func ExecuteCommand(t *testing.T, environment string, commands ...string) ([]string, error) {
	fullCommand := []string{"run", "pars.go"}
	fullCommand = append(fullCommand, commands...)
	fullCommand = append(fullCommand, []string{"--log-level", "verbose", "--env", environment}...)

	t.Logf("executing command %v", fullCommand[2:])
	cmd := exec.Command("go", fullCommand...)
	cmd.Dir = utils.GetSourceLocation()

	output, err := cmd.CombinedOutput()
	t.Logf(">>> %v", string(output))

	return fullCommand[2:], err
}
