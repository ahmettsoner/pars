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

type ListTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	noCleanOnFail bool
	faker         *faker.Faker
}

func (suite *ListTestSuite) SetupSuite() {

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
func (suite *ListTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *ListTestSuite) SetupTest() {
}
func (suite *ListTestSuite) TearDownTest() {
}

func (suite *ListTestSuite) TestListCurrentWorkspace() {
	name := suite.faker.Workspace.Name()
	common.InitializeNewWorkspace(suite.T(), filepath.Join(suite.testArea, name), name, suite.environment)

	commands := []string{
		"workspace",
		"list",
	}

	listOutput, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
	require.NoError(suite.T(), err, "failed to retrieve workspace description")

	require.NotEmpty(suite.T(), listOutput, "Workspace description is not valid")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}
