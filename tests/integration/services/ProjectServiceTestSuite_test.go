package services

import (
	"fmt"
	"os"
	"testing"

	"parsdevkit.net/models"
	applicationproject "parsdevkit.net/structs/project/application-project"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"
	"parsdevkit.net/core/test/objects"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProjectServiceTestSuite struct {
	suite.Suite
	service       services.ApplicationProjectServiceInterface
	environment   string
	testArea      string
	workspaceName string
	workspace     workspace.WorkspaceSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ProjectServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspaceName = suite.faker.Workspace.Name()
	suite.service = services.NewApplicationProjectService(suite.environment)

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspaceName)
	workspaceBase := InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspaceName, suite.environment)
	suite.workspace = workspaceBase.Specifications

	suite.T().Log("Project creation completed")
}
func (suite *ProjectServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		RemoveWorkspace(suite.T(), suite.workspaceName, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *ProjectServiceTestSuite) SetupTest() {
}
func (suite *ProjectServiceTestSuite) TearDownTest() {
}

func (suite *ProjectServiceTestSuite) Test_ListProjects_BySetAndLayer_SingleLayer() {

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	projectSet := suite.faker.Project.Set()

	project1 := *objects.BasicProject_WithName(projectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project1.Specifications.Set = projectSet
	project1.Specifications.Configuration.Layers = append(project1.Specifications.Configuration.Layers, applicationproject.NewLayer_NameOnly("sample:layer"))
	project2 := *objects.BasicProject_WithName(projectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project2.Specifications.Set = projectSet
	project2.Specifications.Configuration.Layers = append(project2.Specifications.Configuration.Layers, applicationproject.NewLayer_NameOnly("sample:layer"))

	temp1, err := suite.service.Create(project1, false)
	require.NoError(suite.T(), err, "Failed to save project")

	utils.PrintFields2(temp1)

	temp2, err := suite.service.Create(project2, false)
	require.NoError(suite.T(), err, "Failed to save project")

	projectList, err := suite.service.ListBySetAndLayers(projectSet, "sample:layer")
	require.NoError(suite.T(), err, "Failed to retrieve projects")
	assert.Equal(suite.T(), 2, len(*projectList))

	assert.Equal(suite.T(), *temp1, (*projectList)[0])
	assert.Equal(suite.T(), *temp2, (*projectList)[1])

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			for _, project := range *projectList {
				suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			}
		}
	})
}
func (suite *ProjectServiceTestSuite) Test_ListProjects_BySetAndLayer_MultipleLayer() {

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	projectSet := suite.faker.Project.Set()

	project1 := *objects.BasicProject_WithName(projectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project1.Specifications.Set = projectSet
	project1.Specifications.Configuration.Layers = append(project1.Specifications.Configuration.Layers,
		applicationproject.NewLayer_NameOnly("sample:layer"),
		applicationproject.NewLayer_NameOnly("sample:layer2"),
		applicationproject.NewLayer_NameOnly("sample:layer3"),
	)
	project2 := *objects.BasicProject_WithName(projectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project2.Specifications.Set = projectSet
	project2.Specifications.Configuration.Layers = append(project2.Specifications.Configuration.Layers,
		applicationproject.NewLayer_NameOnly("sample:layer"),
		applicationproject.NewLayer_NameOnly("sample:layer2"),
		applicationproject.NewLayer_NameOnly("sample:layer3"),
	)

	temp1, err := suite.service.Create(project1, false)
	require.NoError(suite.T(), err, "Failed to save project")

	temp2, err := suite.service.Create(project2, false)
	require.NoError(suite.T(), err, "Failed to save project")

	projectList, err := suite.service.ListBySetAndLayers(projectSet, "sample:layer", "sample:layer2", "sample:layer3")
	require.NoError(suite.T(), err, "Failed to retrieve projects")
	assert.Equal(suite.T(), 2, len(*projectList))

	assert.Equal(suite.T(), *temp1, (*projectList)[0])
	assert.Equal(suite.T(), *temp2, (*projectList)[1])

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			for _, project := range *projectList {
				suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			}
		}
	})
}
func (suite *ProjectServiceTestSuite) Test_ListProjects_ByWorkspace() {

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()

	project1 := *objects.BasicProject_WithName(projectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project2 := *objects.BasicProject_WithName(projectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)

	_, err := suite.service.Create(project1, false)
	require.NoError(suite.T(), err, "Failed to save project")

	_, err = suite.service.Create(project2, false)
	require.NoError(suite.T(), err, "Failed to save project")

	projectList, err := suite.service.ListByWorkspace(suite.workspaceName)
	require.NoError(suite.T(), err, "Failed to retrieve projects")
	assert.LessOrEqual(suite.T(), 2, len(*projectList))

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			for _, project := range *projectList {
				suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			}
		}
	})
}

func (suite *ProjectServiceTestSuite) Test_ListGroupProjects_ByOnlyGroupFullNameAndWorkspace() {

	groupName := suite.faker.Project.Group()
	groupPath := suite.faker.Project.Path(1)
	projectName1 := suite.faker.Project.Name()
	projectPath1 := suite.faker.Project.Path(1)
	projectName2 := suite.faker.Project.Name()
	projectPath2 := suite.faker.Project.Path(1)
	group := CreateGroup(suite.T(), groupName, groupPath, suite.environment)

	project1 := *objects.BasicProject_WithName(projectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project1.Specifications.Path = utils.PathToArray(projectPath1)
	project1.Specifications.ProjectIdentifier.Group = group.Name
	project1.Specifications.GroupObject = group.Specifications
	project2 := *objects.BasicProject_WithName(projectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project2.Specifications.Path = utils.PathToArray(projectPath2)
	project2.Specifications.ProjectIdentifier.Group = group.Name
	project2.Specifications.GroupObject = group.Specifications

	temp1, err := suite.service.Create(project1, true)
	require.NoError(suite.T(), err, "Failed to save project")

	temp2, err := suite.service.Create(project2, true)
	require.NoError(suite.T(), err, "Failed to save project")

	projectList, err := suite.service.ListByFullNameWorkspace(fmt.Sprintf("%v/", group.Specifications.Name), suite.workspaceName)
	require.NoError(suite.T(), err, "Failed to retrieve projects")
	assert.Equal(suite.T(), 2, len(*projectList))

	assert.Equal(suite.T(), *temp1, (*projectList)[0])
	assert.Equal(suite.T(), *temp2, (*projectList)[1])

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(fmt.Sprintf("%v/", group.Specifications.Name), suite.workspaceName, true, true)
			RemoveGroup(suite.T(), group.Specifications.Name, suite.environment)
		}
	})
}
func (suite *ProjectServiceTestSuite) Test_GetProject_ByFullNameAndWorkspace() {

	groupName := suite.faker.Project.Group()
	groupPath := suite.faker.Project.Path(1)
	projectName := suite.faker.Project.Name()
	projectPath := suite.faker.Project.Path(1)
	group := CreateGroup(suite.T(), groupName, groupPath, suite.environment)

	project1 := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project1.Specifications.Path = utils.PathToArray(projectPath)
	project1.Specifications.ProjectIdentifier.Group = group.Name
	project1.Specifications.GroupObject = group.Specifications

	temp1, err := suite.service.Create(project1, true)
	require.NoError(suite.T(), err, "Failed to save project")

	retrievedProject, err := suite.service.GetByFullNameWorkspace(project1.GetFullName(), suite.workspaceName)
	require.NoError(suite.T(), err, "Failed to retrieve projects")

	assert.Equal(suite.T(), *temp1, *retrievedProject)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(retrievedProject.GetFullName(), retrievedProject.Specifications.Workspace, true, true)
			RemoveGroup(suite.T(), group.Specifications.Name, suite.environment)
		}
	})
}

func TestProjectServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceTestSuite))
}
