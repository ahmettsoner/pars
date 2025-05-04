package common

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
	textTemplate "text/template"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func SubmitProject(commander CommanderType, t *testing.T, project string, environment string) {
	commands := []string{"project", "ubmit", project}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveProject(t *testing.T, project string, environment string) {
	commands := []string{"project", "remove", project}

	_, err := ExecuteCommand(t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func RemoveProjectWithWorkspace(commander CommanderType, t *testing.T, project, workspace, environment string) {
	commands := []string{"project", "remove", project, "--workspace", workspace}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func SubmitProjectFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"project", "submit", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}

func CreateTempFileFromTemplate(t *testing.T, declarationFile string, testArea string, data any) string {
	filename := filepath.Base(declarationFile)
	name := filename[:len(filename)-len(filepath.Ext(filename))]
	extension := filepath.Ext(filename)
	tmplContent, err := os.ReadFile(declarationFile)

	logrus.Infof("File found at: %v", declarationFile)
	if err != nil {
		logrus.Fatal(err)
	}
	var outputBuffer bytes.Buffer
	err = textTemplate.Must(textTemplate.New("ProjectFromTemplate").Parse(string(tmplContent))).Execute(&outputBuffer, data)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(filepath.Join(testArea, "test-files"), os.ModePerm)
	if err != nil {
		return ""
	}

	tmpfile, err := os.CreateTemp(filepath.Join(testArea, "test-files"), fmt.Sprintf("%v_*%v", name, extension))
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return ""
	}
	// defer os.Remove(tmpfile.Name())

	_, err = io.Copy(tmpfile, &outputBuffer)
	if err != nil {
		fmt.Println("Error copying buffer to file:", err)
		return ""
	}

	err = tmpfile.Close()
	if err != nil {
		fmt.Println("Error closing temp file:", err)
		return ""
	}
	logrus.Infof("Temp File created at: %v", tmpfile.Name())

	return tmpfile.Name()
}

func RemoveProjectFromFile(commander CommanderType, t *testing.T, declarationFile, environment string) {
	commands := []string{"project", "remove", "-f", declarationFile}

	_, err := ExecuteCommandWithSelector(commander, t, environment, commands...)
	require.NoErrorf(t, err, "Failed to execute command %v", commands)
}
