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

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleDependency_WithDefaultVersion() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	pack, _ := suite.faker.Dotnet.Package("Net8")
	project.Specifications.Configuration.Dependencies = []applicationproject.Package{
		applicationproject.NewPackage_Basic(pack),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project dependencies")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_SingleDependency_WithVersion() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	pack, packVersion := suite.faker.Dotnet.Package("Net8")
	project.Specifications.Configuration.Dependencies = []applicationproject.Package{
		applicationproject.NewPackage(pack, packVersion),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project dependencies")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleDependencies_WithDefaultVersion() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	pack1, _ := suite.faker.Dotnet.Package("Net8")
	pack2, _ := suite.faker.Dotnet.Package("Net8")
	project.Specifications.Configuration.Dependencies = []applicationproject.Package{
		applicationproject.NewPackage_Basic(pack1),
		applicationproject.NewPackage_Basic(pack2),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project dependencies")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleDependencies_WithVersion() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	pack1, packVersion1 := suite.faker.Dotnet.Package("Net8")
	pack2, packVersion2 := suite.faker.Dotnet.Package("Net8")
	project.Specifications.Configuration.Dependencies = []applicationproject.Package{
		applicationproject.NewPackage(pack1, packVersion1),
		applicationproject.NewPackage(pack2, packVersion2),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project dependencies")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_CreateBasicProject_MultipleDependencies_WithVersionAndDefaultVersion() {

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	pack1, packVersion1 := suite.faker.Dotnet.Package("Net8")
	pack2, _ := suite.faker.Dotnet.Package("Net8")
	project.Specifications.Configuration.Dependencies = []applicationproject.Package{
		applicationproject.NewPackage(pack1, packVersion1),
		applicationproject.NewPackage_Basic(pack2),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project dependencies")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func TestProjectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceTestSuite))
}
