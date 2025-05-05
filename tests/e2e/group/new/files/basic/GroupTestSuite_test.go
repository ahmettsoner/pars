package basic

import (
	"os"
	"testing"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GroupTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
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

func (suite *GroupTestSuite) TestCreateBasicGroup() {
	declarationFile := utils.GetTestFileFromCurrentLocation("basic_group.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	name := suite.faker.Project.Group()
	var structData = struct {
		Name string
	}{
		Name: name,
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewGroupService(suite.environment)
	_, err := service.GetByName(name)
	require.NoError(suite.T(), err, "Failed to get group by name.")

	suite.T().Cleanup(func() {
		common.RemoveGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GroupTestSuite))
}
