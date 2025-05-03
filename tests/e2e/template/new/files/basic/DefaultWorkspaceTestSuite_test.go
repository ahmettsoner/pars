package basic

import (
	"os"
	"testing"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/assert"
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
		common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *DefaultWorkspaceTestSuite) SetupTest() {
}
func (suite *DefaultWorkspaceTestSuite) TearDownTest() {
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicTemplate() {
	declarationFile := utils.GetTestFileFromCurrentLocation("basic_template.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	resourceName := suite.faker.Project.Name()
	resourceSet := suite.faker.Project.Set()
	resourcePath := suite.faker.Project.Path(1)
	var structData = struct {
		Name   string
		Set    string
		Path   string
		Output string
		Layers []string
		Code   string
	}{
		Name:   resourceName,
		Set:    resourceSet,
		Path:   resourcePath,
		Output: "foo.bar",
		Layers: []string{"foo", "bar"},
		Code: `namespace {{ .Package }};

      public class {{ .Name }}
      {
        {{range .attributes}}
        {{.visibility}} {{.type}}  {{.name}} { get; set; }
        {{end}}

        public string GreetingsFromPars()
        {
          return "Pars wishes you success :)";
        }
      }`,
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitTemplateFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewCodeTemplateService(suite.environment)
	template, err := service.GetByName(structData.Name)
	require.NoError(suite.T(), err, "Failed to get template by name.")
	assert.Equal(suite.T(), structData.Name, template.Header.Name)

	suite.T().Cleanup(func() {
		common.RemoveTemplateFromFile(suite.T(), templateFile, suite.environment)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
