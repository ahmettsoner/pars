package console

import (
	"os"
	"testing"

	"parsdevkit.net/core/utils"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GroupTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	noCleanOnFail bool
	faker         *faker.Faker
}

func (suite *GroupTestSuite) SetupSuite() {

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
func (suite *GroupTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *GroupTestSuite) SetupTest() {
}
func (suite *GroupTestSuite) TearDownTest() {
}

// func (suite *GroupTestSuite) TestCreateGroupGroup() {

// 	name := suite.faker.Project.Group()

// 	commands := []string{
// 		"group",
// 		"new",
// 		name,
// 	}
// 	_, err := common.ExecuteCommandWithSelector(common.CommanderTypes.GO, suite.T(), suite.environment, commands...)
// 	require.NoError(suite.T(), err, "failed to create new group")

// 	service := services.NewGroupService(suite.environment)
// 	_, err = service.GetByName(name)
// 	require.NoError(suite.T(), err, "Failed to get group by name.")

// 	suite.T().Cleanup(func() {
// 		if !suite.noCleanOnFail || !suite.T().Failed() {
// 			common.RemoveGroup(suite.T(), name, suite.environment)
// 			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 		}
// 	})
// }

func TestGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GroupTestSuite))
}
