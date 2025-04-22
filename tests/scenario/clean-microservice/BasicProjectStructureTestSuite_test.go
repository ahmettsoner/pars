package basic

import (
	"os"
	"testing"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DefaultWorkspaceTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *DefaultWorkspaceTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")

	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()
	// suite.workspace = "sabit-ws"

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	common.InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)

	suite.T().Logf("Switching to workspace (%v)...", suite.workspace)
	common.SwitchToWorkspace(suite.T(), suite.workspace, suite.environment)

	suite.T().Log("Test suite setup completed")
}
func (suite *DefaultWorkspaceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		// common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *DefaultWorkspaceTestSuite) SetupTest() {
}
func (suite *DefaultWorkspaceTestSuite) TearDownTest() {
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject() {
	groupsDeclarationFile := utils.GetTestFileFromCurrentLocation("projects.yaml")
	common.SubmitGroupFromFile(common.CommanderTypes.GO, suite.T(), groupsDeclarationFile, suite.environment)

	projectsDeclarationFile := utils.GetTestFileFromCurrentLocation("projects.yaml")
	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), projectsDeclarationFile, suite.environment)

	resourcesDeclarationFile := utils.GetTestFileFromCurrentLocation("resources.yaml")
	common.SubmitResourceFromFile(common.CommanderTypes.GO, suite.T(), resourcesDeclarationFile, suite.environment)

	templatesDeclarationFile := utils.GetTestFileFromCurrentLocation("templates.yaml")
	common.SubmitTemplateFromFile(common.CommanderTypes.GO, suite.T(), templatesDeclarationFile, suite.environment)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.RemoveResourceFromFile(suite.T(), resourcesDeclarationFile, suite.environment)

			common.RemoveTemplateFromFile(suite.T(), templatesDeclarationFile, suite.environment)

			// common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), projectsDeclarationFile, suite.environment)

			// common.RemoveGroupFromFile(common.CommanderTypes.GO, suite.T(), groupsDeclarationFile, suite.environment)

			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
