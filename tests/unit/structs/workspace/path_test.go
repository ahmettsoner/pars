package workspace

import (
	"path/filepath"
	"testing"

	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"github.com/stretchr/testify/assert"
)

func Test_Workspace_Absolute_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := applicationWorkspace.WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.WorkspaceIdentifier{
			Path: "workspace",
		},
	}

	// Act
	expected := filepath.Join("workspace")

	// Assert
	a.Equal(expected, data.GetAbsolutePath())
}

func Test_Workspace_Absolute_CodeBase_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := applicationWorkspace.WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.WorkspaceIdentifier{
			Path: "workspace",
		},
	}

	// Act
	expected := filepath.Join("workspace", applicationWorkspace.CodeBasePath)

	// Assert
	a.Equal(expected, data.GetCodeBaseFolder())
}

func Test_Workspace_Absolute_Templates_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := applicationWorkspace.WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.WorkspaceIdentifier{
			Path: "workspace",
		},
	}

	// Act
	expected := filepath.Join("workspace", applicationWorkspace.TemplatesPath)

	// Assert
	a.Equal(expected, data.GetTemplatesFolder())
}

func Test_Workspace_Absolute_Resources_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	data := applicationWorkspace.WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.WorkspaceIdentifier{
			Path: "workspace",
		},
	}

	// Act
	expected := filepath.Join("workspace", applicationWorkspace.ResourcesPath)

	// Assert
	a.Equal(expected, data.GetResourcesFolder())
}
