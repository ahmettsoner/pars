package net8

import (
	"os"
	"testing"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/dotnet/managers"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SimpleGroupProjectTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	group         string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *SimpleGroupProjectTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()
	suite.group = suite.faker.Project.Group()
	CreateNewTestProjectGroupAndPath(suite.T(), suite.group, suite.group, suite.testArea, suite.workspace)

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.manager = managers.NewDotnetManager()

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)
}
func (suite *SimpleGroupProjectTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *SimpleGroupProjectTestSuite) SetupTest() {
}
func (suite *SimpleGroupProjectTestSuite) TearDownTest() {
}

// func (suite *SimpleGroupProjectTestSuite) Test_AddNewProject_Basic() {

// 	projectName := suite.faker.Project.Name()
// 	project := CreateNewTestProjectWithGroup(common.CommanderTypes.GO, suite.T(), projectName, suite.testArea, suite.workspace, suite.group, suite.group)

// 	projectState, err := suite.manager.HasProjectOnGroup(project)
// 	require.NoError(suite.T(), err, "failed to check project on group")
// 	assert.True(suite.T(), projectState)

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *SimpleGroupProjectTestSuite) Test_AddNewProject_CustomProjectPath() {

// 	projectName := suite.faker.Project.Name()
// 	projectPath := suite.faker.Project.Path(1)
// 	project := CreateNewTestProjectWithGroupAndPath(suite.T(), projectName, projectPath, suite.testArea, suite.workspace, suite.group, suite.group)

// 	projectState, err := suite.manager.HasProjectOnGroup(project)
// 	require.NoError(suite.T(), err, "failed to check project on group")
// 	assert.True(suite.T(), projectState)

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// // TODO:? Burda kurgu doğru mu?
// func (suite *SimpleGroupProjectTestSuite) Test_AddNewProject_CustomGroupPath() {

// 	projectName := suite.faker.Project.Name()
// 	projectPath := suite.faker.Project.Path(1)
// 	groupName := suite.faker.Project.Group()
// 	groupPath := suite.faker.Project.Path(1)
// 	CreateNewTestProjectGroupAndPath(suite.T(), groupName, groupPath, suite.testArea, suite.workspace)
// 	project := CreateNewTestProjectWithGroupAndPath(suite.T(), projectName, projectPath, suite.testArea, suite.workspace, groupName, groupPath)

// 	projectState, err := suite.manager.HasProjectOnGroup(project)
// 	require.NoError(suite.T(), err, "failed to check project on group")
// 	assert.True(suite.T(), projectState)

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

func TestSimpleGroupProjectTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleGroupProjectTestSuite))
}
