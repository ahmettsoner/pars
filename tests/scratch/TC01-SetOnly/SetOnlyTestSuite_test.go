package basic

import (
	"parsdevkit.net/application"
	"parsdevkit.net/application/models/label"

	"os"
	"testing"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SetOnlyTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	faker         *faker.Faker
	set           string
	noCleanOnFail bool
}

func (suite *SetOnlyTestSuite) SetupSuite() {

	suite.faker = faker.NewFaker()

	suite.T().Log("Preparing test suite...")

	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()
	suite.set = suite.faker.Project.Set()

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
func (suite *SetOnlyTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspace(suite.T(), suite.workspace, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *SetOnlyTestSuite) SetupTest()    {}
func (suite *SetOnlyTestSuite) TearDownTest() {}

func (suite *SetOnlyTestSuite) TestCreateBasicResource() {

	projectName := suite.faker.Project.Name()
	projectTemplateFile1 := CreateNewProjectFromTemplateFile(common.CommanderTypes.GO, suite.T(), suite.environment, suite.testArea, projectName, suite.set, []string{}, []string{}, []string{})

	resourceName := suite.faker.Project.Name()
	resourcePath := suite.faker.Project.Path(1)
	resourcePackages := resourcePath

	var structData = struct {
		Name    string
		Set     string
		Path    string
		Package string
		Layers  []string
		Tags    []string
		Labels  []label.Label
	}{
		Name:    resourceName,
		Set:     suite.set,
		Path:    resourcePath,
		Package: resourcePackages,
		Layers:  []string{},
		Tags:    []string{},
		Labels:  []label.Label{},
	}

	declarationFile := application.GetTestFileFromCurrentLocation("resources.yaml")
	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), projectTemplateFile1, suite.environment)
		os.Remove(projectTemplateFile1)

		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestSetOnlyTestSuite(t *testing.T) {
	suite.Run(t, new(SetOnlyTestSuite))
}

func CreateNewProjectFromTemplateFile(commander common.CommanderType, t *testing.T, environment, testArea, name, set string, layers []string, dependencies []string, references []string) string {

	declarationFile := application.GetTestFileFromCurrentLocation("projects.yaml")

	var structData = struct {
		Name         string
		Platform     string
		Type         string
		Set          string
		Layers       []string
		Dependencies []string
		References   []string
	}{
		Name:         name,
		Platform:     "dotnet",
		Type:         "library",
		Set:          set,
		Layers:       layers,
		References:   references,
		Dependencies: dependencies,
	}

	templateFile := common.CreateTempFileFromTemplate(t, declarationFile, testArea, structData)

	common.Apply(commander, t, templateFile, environment)

	return templateFile
}
