package library

import (
	"fmt"
	"os"
	"testing"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DotnetLibraryTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	noCleanOnFail bool
	faker         *faker.Faker
}

func (suite *DotnetLibraryTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	common.InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)

	suite.T().Log("Test suite setup completed")
}

func (suite *DotnetLibraryTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *DotnetLibraryTestSuite) SetupTest()    {}
func (suite *DotnetLibraryTestSuite) TearDownTest() {}

func (suite *DotnetLibraryTestSuite) TestCreateBasicProject() {

	projectName := suite.faker.Project.Name()
	fullName := fmt.Sprintf("%v", projectName)
	commands := []string{
		"project",
		"new",
		"dotnet",
		"library",
		fullName,
		"--workspace",
		suite.workspace,
	}
	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to create new project")

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fullName, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveProjectWithWorkspace(common.CommanderTypes.GO, suite.T(), fullName, suite.workspace, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func (suite *DotnetLibraryTestSuite) TestCreateGroupProject() {

	projectGroup := suite.faker.Project.Group()
	projectName := suite.faker.Project.Name()
	fullName := fmt.Sprintf("%v/%v", projectGroup, projectName)

	common.SubmitGroup(suite.T(), projectGroup, suite.environment)

	commands := []string{
		"project",
		"new",
		"dotnet",
		"library",
		fullName,
		"--workspace",
		suite.workspace,
	}
	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to create new project")

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fullName, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveProjectWithWorkspace(common.CommanderTypes.GO, suite.T(), fullName, suite.workspace, suite.environment)
			common.RemoveGroup(suite.T(), projectGroup, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func (suite *DotnetLibraryTestSuite) TestCreateGroupLayeredProject() {

	projectGroup := suite.faker.Project.Group()
	projectName := suite.faker.Project.Name()
	fullName := fmt.Sprintf("%v/%v", projectGroup, projectName)

	common.SubmitGroup(suite.T(), projectGroup, suite.environment)

	commands := []string{
		"project",
		"new",
		"dotnet",
		"library",
		fullName,
		"--methodology",
		"layered",
		"--workspace",
		suite.workspace,
	}
	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to create new project")

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fullName, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveProjectWithWorkspace(common.CommanderTypes.GO, suite.T(), fullName, suite.workspace, suite.environment)
			common.RemoveGroup(suite.T(), projectGroup, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func (suite *DotnetLibraryTestSuite) TestCreateGroupNTierProject() {

	projectGroup := suite.faker.Project.Group()
	projectName := suite.faker.Project.Name()
	fullName := fmt.Sprintf("%v/%v", projectGroup, projectName)

	common.SubmitGroup(suite.T(), projectGroup, suite.environment)

	commands := []string{
		"project",
		"new",
		"dotnet",
		"library",
		fullName,
		"--methodology",
		"ntier",
		"--workspace",
		suite.workspace,
	}
	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to create new project")

	projectNames := []string{
		fmt.Sprintf("%v.Service", fullName),
		fmt.Sprintf("%v.Database", fullName),
		fmt.Sprintf("%v.Core", fullName),
	}

	service := services.NewApplicationProjectService(suite.environment)

	for _, projectName := range projectNames {
		project, err := service.GetByFullNameWorkspace(projectName, suite.workspace)
		require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

		projectStructure, err := service.ValidateProjectStructure(project.Specifications)
		require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
		assert.Equal(suite.T(), true, projectStructure)

		projectReference, err := service.ValidateProjectReferences(project.Specifications)
		require.NoError(suite.T(), err, "Validation of the project references failed.")
		assert.Equal(suite.T(), true, projectReference)

		projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
		require.NoError(suite.T(), err, "Validation of the project packages failed.")
		assert.Equal(suite.T(), true, projectPackages)
	}

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveProjectWithWorkspace(common.CommanderTypes.GO, suite.T(), fmt.Sprintf("%v/", projectGroup), suite.workspace, suite.environment)
			common.RemoveGroup(suite.T(), projectGroup, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func TestDotnetLibraryTestSuite(t *testing.T) {
	suite.Run(t, new(DotnetLibraryTestSuite))
}
