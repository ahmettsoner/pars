package console

import (
	"os"
	"path/filepath"
	"testing"

	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"

	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/application"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type InitTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	noCleanOnFail bool
	faker         *faker.Faker
}

func (suite *InitTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Log("Test suite setup completed")
}
func (suite *InitTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *InitTestSuite) SetupTest() {
}
func (suite *InitTestSuite) TearDownTest() {
}

// func (suite *InitTestSuite) TestInitializeBasicWorkspaceDefaultNameAndPath() {

// 	name := "workspace"
// 	dirInTestArea := filepath.Join(suite.testArea, suite.faker.Project.Path(1))
// 	commands := []string{
// 		"init",
// 	}

// 	_, err := common.ExecuteCommandWithSelectorOnPath(common.InitializeNewWorkspace(common.CommanderTypes.GO, suite.T(), suite.environment, dirInTestArea, commands...)
// 	require.NoError(suite.T(), err, "failed to initialize workspace")

// 	service := services.NewWorkspaceService(suite.environment)
// 	workspace, err := service.GetByName(name)
// 	require.NoError(suite.T(), err, "Failed to get workspace by name.")
// require.NotNil(suite.T(), workspace, "Not found workspace '%s'", name)

// 	require.Equal(suite.T(), workspace.Specifications.GetAbsolutePath(), filepath.Join(dirInTestArea, name), "Workspace path is not valid")

// 	suite.T().Cleanup(func() {
// 		if !suite.noCleanOnFail || !suite.T().Failed() {
// 			common.RemoveWorkspace(suite.T(), name, suite.environment)
// 			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), dirInTestArea)
// 		}
// 	})
// }

// func (suite *InitTestSuite) TestInitializeBasicWorkspaceOnDefaultPath() {

// 	name := suite.faker.Workspace.Name()
// 	dirInTestArea := filepath.Join(suite.testArea, suite.faker.Project.Path(1))
// 	commands := []string{
// 		"init",
// 		name,
// 	}

// 	_, err := common.ExecuteCommandWithSelectorOnPath(common.InitializeNewWorkspace(common.CommanderTypes.GO, suite.T(), suite.environment, dirInTestArea, commands...)
// 	require.NoError(suite.T(), err, "failed to initialize workspace")

// 	service := services.NewWorkspaceService(suite.environment)
// 	workspace, err := service.GetByName(name)
// 	require.NoError(suite.T(), err, "Failed to get workspace by name.")
// require.NotNil(suite.T(), workspace, "Not found workspace '%s'", name)

// 	require.Equal(suite.T(), workspace.Specifications.GetAbsolutePath(), filepath.Join(dirInTestArea, name), "Workspace path is not valid")

// 	suite.T().Cleanup(func() {
// 		if !suite.noCleanOnFail || !suite.T().Failed() {
// 			common.RemoveWorkspace(suite.T(), name, suite.environment)
// 			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), dirInTestArea)
// 		}
// 	})
// }

// func (suite *InitTestSuite) TestInitializeBasicWorkspaceOnCurrentPath() {

// 	name := suite.faker.Workspace.Name()
// 	dirInTestArea := filepath.Join(suite.testArea, suite.faker.Project.Path(1))
// 	commands := []string{
// 		"init",
// 		name,
// 		".",
// 	}

// 	_, err := common.ExecuteCommandWithSelectorOnPath(common.InitializeNewWorkspace(common.CommanderTypes.GO, suite.T(), suite.environment, dirInTestArea, commands...)
// 	require.NoError(suite.T(), err, "failed to initialize workspace")

// 	service := services.NewWorkspaceService(suite.environment)
// 	workspace, err := service.GetByName(name)
// 	require.NoError(suite.T(), err, "Failed to get workspace by name.")
// require.NotNil(suite.T(), workspace, "Not found workspace '%s'", name)

// 	require.Equal(suite.T(), workspace.Specifications.GetAbsolutePath(), filepath.Join(dirInTestArea), "Workspace path is not valid")

// 	suite.T().Cleanup(func() {
// 		if !suite.noCleanOnFail || !suite.T().Failed() {
// 			common.RemoveWorkspace(suite.T(), name, suite.environment)
// 			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), dirInTestArea)
// 		}
// 	})
// }

func (suite *InitTestSuite) TestInitializeBasicWorkspaceOnRelativePath() {

	name := suite.faker.Workspace.Name()
	dirInTestArea := filepath.Join(suite.testArea, suite.faker.Project.Path(1))
	relativePath, err := file.FindRelativePath(application.GetSourceLocation(), dirInTestArea)
	require.NoError(suite.T(), err, "failed find relative path")

	commands := []string{
		"init",
		name,
		relativePath,
	}

	_, err = common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to initialize workspace")

	service := workspaceWorkspace.NewWorkspaceService(suite.environment)
	workspace, err := service.GetByName(name)
	require.NoError(suite.T(), err, "Failed to get workspace by name.")
	require.NotNil(suite.T(), workspace, "Not found workspace '%s'", name)

	require.Equal(suite.T(), workspace.Specifications.GetAbsolutePath(), filepath.Join(dirInTestArea, name), "Workspace path is not valid")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveWorkspace(suite.T(), name, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), dirInTestArea)
		}
	})
}

func (suite *InitTestSuite) TestInitializeBasicWorkspaceOnAbsolutePath() {

	name := suite.faker.Workspace.Name()
	dirInTestArea := filepath.Join(suite.testArea, suite.faker.Project.Path(1))
	relativePath := suite.faker.Project.Path(2)
	absolutePath := filepath.Join(dirInTestArea, relativePath)
	commands := []string{
		"init",
		name,
		absolutePath,
	}

	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to initialize workspace")

	service := workspaceWorkspace.NewWorkspaceService(suite.environment)
	workspace, err := service.GetByName(name)
	require.NoError(suite.T(), err, "Failed to get workspace by name.")
	require.NotNil(suite.T(), workspace, "Not found workspace '%s'", name)

	require.Equal(suite.T(), workspace.Specifications.GetAbsolutePath(), filepath.Join(absolutePath, name), "Workspace path is not valid")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveWorkspace(suite.T(), name, suite.environment)
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), dirInTestArea)
		}
	})
}

func TestInitTestSuite(t *testing.T) {
	suite.Run(t, new(InitTestSuite))
}
