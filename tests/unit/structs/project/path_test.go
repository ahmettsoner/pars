package project

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	applicationGroup "parsdevkit.net/application/structs/group"

	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	layerPkg "parsdevkit.net/application/models/layer"
	"parsdevkit.net/structs/group"
	applicationproject "parsdevkit.net/structs/project/application-project"

	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/assert"
)

func Test_Project_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
	}
	// Act
	expected := filepath.Join(fakePath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetRelativeProjectPath())
}

func Test_Project_Absolute_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	fakeWorkspace := testFaker.Workspace.Name()

	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
		WorkspaceObject: applicationWorkspace.WorkspaceIdentifier{
			Path: fakeWorkspace,
		},
	}
	// Act
	expected := filepath.Join(fakeWorkspace, applicationWorkspace.CodeBasePath, fakePath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetAbsoluteProjectPath())
}

func Test_Project_WithGroup_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	fakeGroup := testFaker.Project.Group()

	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
		GroupObject: group.GroupSpecification{
			GroupIdentifier: applicationGroup.GroupIdentifier{
				Path: fakeGroup,
			},
		},
	}
	// Act
	expected := filepath.Join(fakeGroup, fakePath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetRelativeProjectPath())
}

func Test_Project_WithGroup_Absolute_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	fakeGroup := testFaker.Project.Group()
	fakeWorkspace := testFaker.Workspace.Name()

	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
		WorkspaceObject: applicationWorkspace.WorkspaceIdentifier{
			Path: fakeWorkspace,
		},
		GroupObject: group.GroupSpecification{
			GroupIdentifier: applicationGroup.GroupIdentifier{
				Path: fakeGroup,
			},
		},
	}
	// Act
	expected := filepath.Join(fakeWorkspace, applicationWorkspace.CodeBasePath, fakeGroup, fakePath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetAbsoluteProjectPath())
}

func Test_Project_WithGroup_Layer_Relative_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	fakeGroup := testFaker.Project.Group()
	fakeLayerPath := testFaker.Project.Path(1)

	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
		GroupObject: group.GroupSpecification{
			GroupIdentifier: applicationGroup.GroupIdentifier{
				Path: fakeGroup,
			},
		},
		Configuration: applicationproject.Configuration{
			Layers: []applicationProject.Layer{
				applicationProject.Layer{
					LayerIdentifier: layerPkg.LayerIdentifier{
						Name: "layer-identifier",
					},
					Path: fakeLayerPath,
				},
			},
		},
	}
	// Act
	expected := filepath.Join(fakeGroup, fakePath, fakeLayerPath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetRelativeProjectLayerPath("layer-identifier"))
}

func Test_Project_WithGroup_Layer_Absolute_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	testFaker := faker.NewFaker()
	fakePath := testFaker.Project.Path(1)
	fakeGroup := testFaker.Project.Group()
	fakeLayerPath := testFaker.Project.Path(1)
	fakeWorkspace := testFaker.Workspace.Name()

	data := applicationproject.ProjectSpecification{
		ProjectIdentifier: applicationProject.ProjectIdentifier{
			Path: utils.PathToArray(fakePath),
		},
		WorkspaceObject: applicationWorkspace.WorkspaceIdentifier{
			Path: fakeWorkspace,
		},
		GroupObject: group.GroupSpecification{
			GroupIdentifier: applicationGroup.GroupIdentifier{
				Path: fakeGroup,
			},
		},
		Configuration: applicationproject.Configuration{
			Layers: []applicationProject.Layer{
				applicationProject.Layer{
					LayerIdentifier: layerPkg.LayerIdentifier{
						Name: "layer-identifier",
					},
					Path: fakeLayerPath,
				},
			},
		},
	}
	// Act
	expected := filepath.Join(fakeWorkspace, applicationWorkspace.CodeBasePath, fakeGroup, fakePath, fakeLayerPath)
	expected = strings.TrimPrefix(expected, string(os.PathSeparator))

	// Assert
	a.Equal(expected, data.GetAbsoluteProjectLayerPath("layer-identifier"))
}
