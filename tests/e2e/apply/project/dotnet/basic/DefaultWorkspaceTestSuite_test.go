package basic

import (
	"fmt"
	"os"
	"testing"

	"parsdevkit.net/models"
	dotnetModels "parsdevkit.net/platforms/dotnet/models"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"parsdevkit.net/application"

	"github.com/stretchr/testify/assert"
	projectApplication "parsdevkit.net/modules/project/application_project"

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
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *DefaultWorkspaceTestSuite) SetupTest() {
}
func (suite *DefaultWorkspaceTestSuite) TearDownTest() {
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject() {
	declarationFile := application.GetTestFileFromCurrentLocation("basic_project.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name     string
		Platform string
		Type     string
	}{
		Name:     projectName,
		Platform: string(models.PlatformTypes.Dotnet),
		Type:     string(dotnetModels.DotnetProjectTypes.Library),
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.Cobra, suite.T(), templateFile, suite.environment)

	service := projectApplication.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(*project)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(*project)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(*project)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.Cobra, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject_WithLayer_NameOnly() {
	declarationFile := application.GetTestFileFromCurrentLocation("basic_project_with_layer.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name     string
		Platform string
		Type     string
		Layers   []string
	}{
		Name:     projectName,
		Platform: string(models.PlatformTypes.Dotnet),
		Type:     string(dotnetModels.DotnetProjectTypes.Library),
		Layers:   []string{"foo", "bar"},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := projectApplication.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(*project)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(*project)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(*project)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateBasicProject_WithReference_NameOnly() {
	declarationFile := application.GetTestFileFromCurrentLocation("basic_project_with_reference.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName1 := suite.faker.Project.Name()
	var structData1 = struct {
		Name       string
		Platform   string
		Type       string
		References []string
	}{
		Name:     projectName1,
		Platform: string(models.PlatformTypes.Dotnet),
		Type:     string(dotnetModels.DotnetProjectTypes.Library),
	}

	templateFile1 := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData1)
	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile1, suite.environment)

	projectName := suite.faker.Project.Name()
	var structData = struct {
		Name       string
		Platform   string
		Type       string
		References []string
	}{
		Name:       projectName,
		Platform:   string(models.PlatformTypes.Dotnet),
		Type:       string(dotnetModels.DotnetProjectTypes.Library),
		References: []string{projectName1},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)
	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := projectApplication.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(structData.Name, suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name.")

	projectStructure, err := service.ValidateProjectStructure(*project)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(*project)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(*project)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile1, suite.environment)
		os.Remove(templateFile1)
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateGroupProject() {
	declarationFile := application.GetTestFileFromCurrentLocation("group_project.yaml")
	suite.T().Logf("Starting test for (%v)", declarationFile)

	projectName := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()

	var structData = struct {
		Name     string
		Group    string
		Platform string
		Type     string
	}{
		Name:     projectName,
		Group:    groupName,
		Platform: string(models.PlatformTypes.Dotnet),
		Type:     string(dotnetModels.DotnetProjectTypes.Library),
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := projectApplication.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fmt.Sprintf("%v/%v", structData.Group, structData.Name), suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(*project)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(*project)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(*project)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *DefaultWorkspaceTestSuite) TestCreateGroupProject_WithLayer_NameOnly() {
	declarationFile := application.GetTestFileFromCurrentLocation("group_project_with_layer.yaml")
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
		Name:     projectName,
		Group:    groupName,
		Platform: string(models.PlatformTypes.Dotnet),
		Type:     string(dotnetModels.DotnetProjectTypes.Library),
		Layers:   []string{"foo", "bar"},
	}

	templateFile := common.CreateTempFileFromTemplate(suite.T(), declarationFile, suite.testArea, structData)

	common.Apply(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)

	service := projectApplication.NewApplicationProjectService(suite.environment)
	project, err := service.GetByFullNameWorkspace(fmt.Sprintf("%v/%v", structData.Group, structData.Name), suite.workspace)
	require.NoError(suite.T(), err, "Failed to get project by full name and workspace name in group.")

	projectStructure, err := service.ValidateProjectStructure(*project)
	require.NoError(suite.T(), err, "Validation of the project structure in a group failed.")
	assert.Equal(suite.T(), true, projectStructure)

	projectReference, err := service.ValidateProjectReferences(*project)
	require.NoError(suite.T(), err, "Validation of the project references failed.")
	assert.Equal(suite.T(), true, projectReference)

	projectPackages, err := service.ValidateProjectDependencies(*project)
	require.NoError(suite.T(), err, "Validation of the project packages failed.")
	assert.Equal(suite.T(), true, projectPackages)

	suite.T().Cleanup(func() {
		common.Destroy(common.CommanderTypes.GO, suite.T(), templateFile, suite.environment)
		os.Remove(templateFile)
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestDefaultWorkspaceTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultWorkspaceTestSuite))
}
