package basic

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject() {
	declarationFile := utils.GetTestFileFromCurrentLocation("basic_project.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name     string
		Platform string
		Type     string
	}{
		Name:     projectName,
		Platform: "angular",
		Type:     "library",
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject_WithLayer_NameOnly() {
	declarationFile := utils.GetTestFileFromCurrentLocation("basic_project_with_layer.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name     string
		Platform string
		Type     string
		Layers   []string
	}{
		Name:     projectName,
		Platform: "angular",
		Type:     "library",
		Layers:   []string{"foo", "bar"},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject_WithReference_NameOnly() {
	declarationFile := utils.GetTestFileFromCurrentLocation("basic_project_with_reference.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName1 := suite.faker.Project.Name()
	var structData1 = struct {
		Name       string
		Platform   string
		Type       string
		References []string
	}{
		Name:     projectName1,
		Platform: "angular",
		Type:     "library",
	}

	templateFile1 := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData1)
	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile1, suite.environment)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name       string
		Platform   string
		Type       string
		References []string
	}{
		Name:       projectName,
		Platform:   "angular",
		Type:       "library",
		References: []string{projectName1},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)
	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile1, suite.environment)
		os.Remove(templateFile1)
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateGroupProject() {
	declarationFile := utils.GetTestFileFromCurrentLocation("group_project.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()

	var structData = struct {
		Name     string
		Group    string
		Platform string
		Type     string
	}{
		Name:     NormalizeText(projectName),
		Group:    groupName,
		Platform: "angular",
		Type:     "library",
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fmt.Sprintf("%v/%v", structData.Group, structData.Name), suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		common.RemoveGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateGroupProject_WithLayer_NameOnly() {
	declarationFile := utils.GetTestFileFromCurrentLocation("group_project_with_layer.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()

	var structData = struct {
		Name     string
		Group    string
		Platform string
		Type     string
		Layers   []string
	}{
		Name:     NormalizeText(projectName),
		Group:    groupName,
		Platform: "angular",
		Type:     "library",
		Layers:   []string{"foo", "bar"},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.SubmitGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
	common.SubmitProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := services.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fmt.Sprintf("%v/%v", structData.Group, structData.Name), suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(project.Specifications)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.RemoveProjectFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		common.RemoveGroupFromFile(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
func NormalizeText(input string) string {
	input = strings.ToLower(input)

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = reg.ReplaceAllString(input, "-")

	input = strings.Trim(input, "-")

	return input
}
