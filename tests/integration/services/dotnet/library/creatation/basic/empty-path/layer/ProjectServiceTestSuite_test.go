package services

import (
	"os"
	"testing"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"
	"parsdevkit.net/core/test/objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProjectServiceTestSuite struct {
	suite.Suite
	service       services.ApplicationProjectServiceInterface
	environment   string
	testArea      string
	workspaceName string
	workspace     workspace.WorkspaceSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ProjectServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspaceName = suite.faker.Workspace.Name()
	suite.service = services.NewApplicationProjectService(suite.environment)

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspaceName)
	workspaceBase := common.InitializeNewWorkspaceWithService(suite.T(), suite.testArea, suite.workspaceName, suite.environment)
	suite.workspace = workspaceBase.Specifications

	suite.T().Log("Project creation completed")
}
func (suite *ProjectServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspaceWithService(suite.T(), suite.workspaceName, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *ProjectServiceTestSuite) SetupTest() {
}
func (suite *ProjectServiceTestSuite) TearDownTest() {
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_WithDefaultPath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnlyDefaultPath(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_WithEmptyPath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnly(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_SinglePath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(1), []string{}),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_MultiplePaths() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(3), []string{}),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_WithDefaultPath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnlyDefaultPath(suite.faker.Project.Layer()),
		applicationproject.NewLayer_NameOnlyDefaultPath(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_WithEmptyPath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnly(suite.faker.Project.Layer()),
		applicationproject.NewLayer_NameOnly(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_SinglePath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(1), []string{}),
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(1), []string{}),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_MultiplePaths() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(3), []string{}),
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(3), []string{}),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_VariousPath() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnly(suite.faker.Project.Layer()),
		applicationproject.NewLayer_NameOnlyDefaultPath(suite.faker.Project.Layer()),
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(1), []string{}),
		applicationproject.NewLayer_Basic(0, suite.faker.Project.Layer(), suite.faker.Project.Path(3), []string{}),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_DefaultPackage() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnlyDefaultPackage(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleLayer_WithEmptyPackage() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnly(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}
func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleLayers_WithDefaultPackage() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.Layers = []applicationproject.Layer{
		applicationproject.NewLayer_NameOnlyDefaultPackage(suite.faker.Project.Layer()),
		applicationproject.NewLayer_NameOnlyDefaultPath(suite.faker.Project.Layer()),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validProjectStructure, err := suite.service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validProjectStructure)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func TestProjectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceTestSuite))
}
