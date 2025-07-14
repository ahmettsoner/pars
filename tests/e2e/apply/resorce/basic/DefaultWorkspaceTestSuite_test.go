package basic

import (
	"os"
	"testing"

	"parsdevkit.net/application"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	resourceObject "parsdevkit.net/modules/resource/object_resource"
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

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	common.InitializeNewWorkspace(common.CommanderTypes.GO, suite.T(), suite.testArea, suite.workspace, suite.environment)

	suite.T().Logf("Switching to workspace (%v)...", suite.workspace)
	common.SwitchToWorkspace(suite.T(), suite.workspace, suite.environment)

	suite.T().Log("Test suite setup completed")
}
func (suite *DefaultWorkspaceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *DefaultWorkspaceTestSuite) SetupTest() {
}
func (suite *DefaultWorkspaceTestSuite) TearDownTest() {
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicResource() {
	declarationFile := application.GetTestFileFromCurrentLocation("basic_resource.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	resourceName := suite.faker.Project.Name()
	resourceSet := suite.faker.Project.Set()
	resourcePath := suite.faker.Project.Path(1)
	var structData = struct {
		Name    string
		Set     string
		Path    string
		Package string
		Layers  []string
	}{
		Name:    resourceName,
		Set:     resourceSet,
		Path:    resourcePath,
		Package: "foo/bar",
		Layers:  []string{"foo", "bar"},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := resourceObject.NewObjectResourceService(suite.environment)
	resource, err := service.GetByName(structData.Name)
	require.NoError(suite.T(), err, "Failed to get resource by name.")
	assert.Equal(suite.T(), structData.Name, resource.Header.Name)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
