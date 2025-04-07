package console

import (
	"os"
	"path/filepath"
	"testing"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RemoveTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	noCleanOnFail bool
	faker         *faker.Faker
}

func (suite *RemoveTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Log("Test suite setup completed")
}
func (suite *RemoveTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *RemoveTestSuite) SetupTest() {
}
func (suite *RemoveTestSuite) TearDownTest() {
}

func (suite *RemoveTestSuite) TestRemoveBasicWorkspace() {
	name := suite.faker.Workspace.Name()
	common.InitializeNewWorkspace(suite.T(), filepath.Join(suite.testArea, name), name, suite.environment)

	commands := []string{
		"workspace",
		"remove",
		name,
	}

	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to remove workspace")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func (suite *RemoveTestSuite) TestRemoveBasicMultipleWorkspace() {
	name := suite.faker.Workspace.Name()
	name1 := suite.faker.Workspace.Name()
	common.InitializeNewWorkspace(suite.T(), filepath.Join(suite.testArea, name), name, suite.environment)
	common.InitializeNewWorkspace(suite.T(), filepath.Join(suite.testArea, name1), name1, suite.environment)

	commands := []string{
		"workspace",
		"remove",
		name,
	}

	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to remove workspace")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}
func TestRemoveTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveTestSuite))
}
