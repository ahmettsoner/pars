package basic

import (
	"os"

	"parsdevkit.net/application"

	"testing"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

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
	testArea := application.GenerateTestArea()
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
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *DefaultWorkspaceTestSuite) SetupTest() {
}
func (suite *DefaultWorkspaceTestSuite) TearDownTest() {
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject() {
	groupsDeclarationFile := application.GetTestFileFromCurrentLocation("projects.yaml")
	common.Apply(common.CommanderTypes.GO, suite.T(), groupsDeclarationFile, suite.environment)

	projectsDeclarationFile := application.GetTestFileFromCurrentLocation("projects.yaml")
	common.Apply(common.CommanderTypes.GO, suite.T(), projectsDeclarationFile, suite.environment)

	resourcesDeclarationFile := application.GetTestFileFromCurrentLocation("resources.yaml")
	common.Apply(common.CommanderTypes.GO, suite.T(), resourcesDeclarationFile, suite.environment)

	templatesDeclarationFile := application.GetTestFileFromCurrentLocation("templates.yaml")
	common.Apply(common.CommanderTypes.GO, suite.T(), templatesDeclarationFile, suite.environment)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			common.Destroy(common.CommanderTypes.GO, suite.T(), resourcesDeclarationFile, suite.environment)

			common.Destroy(common.CommanderTypes.GO, suite.T(), templatesDeclarationFile, suite.environment)

			// common.Destroy(common.CommanderTypes.GO, suite.T(), projectsDeclarationFile, suite.environment)

			// common.Destroy(common.CommanderTypes.GO, suite.T(), groupsDeclarationFile, suite.environment)

			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
		}
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
