package workspace

import (
	"path/filepath"
	"testing"

	"parsdevkit.net/structs/workspace"

	"github.com/stretchr/testify/assert"
)

func Test_Workspace_Absolute_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := workspace.WorkspaceSpecification{
		Path: "workspace",
	}

	// Act
	expected := filepath.Join("workspace")

	// Assert
	a.Equal(expected, data.GetAbsolutePath())
}

func Test_Workspace_Absolute_CodeBase_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := workspace.WorkspaceSpecification{
		Path: "workspace",
	}

	// Act
	expected := filepath.Join("workspace", workspace.CodeBasePath)

	// Assert
	a.Equal(expected, data.GetCodeBaseFolder())
}

func Test_Workspace_Absolute_Templates_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := workspace.WorkspaceSpecification{
		Path: "workspace",
	}

	// Act
	expected := filepath.Join("workspace", workspace.TemplatesPath)

	// Assert
	a.Equal(expected, data.GetTemplatesFolder())
}

func Test_Workspace_Absolute_Resources_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := workspace.WorkspaceSpecification{
		Path: "workspace",
	}

	// Act
	expected := filepath.Join("workspace", workspace.ResourcesPath)

	// Assert
	a.Equal(expected, data.GetResourcesFolder())
}
