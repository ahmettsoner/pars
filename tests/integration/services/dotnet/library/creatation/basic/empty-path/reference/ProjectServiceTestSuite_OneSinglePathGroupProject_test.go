package services

import (
	"os"
	"testing"

	"parsdevkit.net/models"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"
	projectStruct "parsdevkit.net/structs/project"
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

type ProjectServiceOneSinglePathGroupProjectReferenceTestSuite struct {
	suite.Suite
	service       services.ApplicationProjectServiceInterface
	environment   string
	testArea      string
	workspaceName string
	workspace     workspace.WorkspaceSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) SetupSuite() {

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
	workspaceBase := common.InitializeNewWorkspaceWithService(suite.T(), suite.testArea, suite.workspaceName, suite.environment)
	suite.workspace = workspaceBase.Specifications

	suite.T().Log("Project creation completed")
}
func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		common.RemoveWorkspaceWithService(suite.T(), suite.workspaceName, suite.environment)
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) SetupTest() {
}
func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) TearDownTest() {
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_SingleReference_WithEmptyPath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName := suite.faker.Project.Name()
	referenceProject := *objects.BasicProject_WithName(referenceProjectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject.Specifications.ProjectIdentifier.Group = referenceProjectGroupName

	_, err := suite.service.Create(referenceProject, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName, structs.Metadata{}), referenceProject.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}
func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_MultipleReferences_WithEmptyPath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName1 := suite.faker.Project.Name()
	referenceProject1 := *objects.BasicProject_WithName(referenceProjectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject1.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject1.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err := suite.service.Create(referenceProject1, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	referenceProjectName2 := suite.faker.Project.Name()
	referenceProject2 := *objects.BasicProject_WithName(referenceProjectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject2.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject2.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err = suite.service.Create(referenceProject2, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName1, structs.Metadata{}), referenceProject1.Specifications),
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName2, structs.Metadata{}), referenceProject2.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject1.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject2.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_SingleReference_WithDefaultPath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName := suite.faker.Project.Name()
	referenceProject := *objects.BasicProject_WithName(referenceProjectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject.Specifications.Path = []string{referenceProjectName}
	referenceProject.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject.Specifications.ProjectIdentifier.Group = referenceProjectGroupName

	_, err := suite.service.Create(referenceProject, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName, structs.Metadata{}), referenceProject.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}
func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_MultipleReferences_WithDefaultPath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName1 := suite.faker.Project.Name()
	referenceProject1 := *objects.BasicProject_WithName(referenceProjectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject1.Specifications.Path = []string{referenceProjectName1}
	referenceProject1.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject1.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err := suite.service.Create(referenceProject1, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	referenceProjectName2 := suite.faker.Project.Name()
	referenceProject2 := *objects.BasicProject_WithName(referenceProjectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject2.Specifications.Path = []string{referenceProjectName2}
	referenceProject2.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject2.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err = suite.service.Create(referenceProject2, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName1, structs.Metadata{}), referenceProject1.Specifications),
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName2, structs.Metadata{}), referenceProject2.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject1.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject2.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}
func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_SingleReference_WithSinglePath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName := suite.faker.Project.Name()
	referenceProjectPath := suite.faker.Project.Path(1)
	referenceProject := *objects.BasicProject_WithName(referenceProjectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject.Specifications.Path = []string{referenceProjectPath}
	referenceProject.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject.Specifications.ProjectIdentifier.Group = referenceProjectGroupName

	_, err := suite.service.Create(referenceProject, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName, structs.Metadata{}), referenceProject.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_MultipleReferences_WithSinglePath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName1 := suite.faker.Project.Name()
	referenceProjectPath1 := suite.faker.Project.Path(1)
	referenceProject1 := *objects.BasicProject_WithName(referenceProjectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject1.Specifications.Path = []string{referenceProjectPath1}
	referenceProject1.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject1.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err := suite.service.Create(referenceProject1, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	referenceProjectName2 := suite.faker.Project.Name()
	referenceProjectPath2 := suite.faker.Project.Path(1)
	referenceProject2 := *objects.BasicProject_WithName(referenceProjectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject2.Specifications.Path = []string{referenceProjectPath2}
	referenceProject2.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject2.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err = suite.service.Create(referenceProject2, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName1, structs.Metadata{}), referenceProject1.Specifications),
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName2, structs.Metadata{}), referenceProject2.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject1.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject2.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_SingleReference_WithMultiplePath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName := suite.faker.Project.Name()
	referenceProjectPath := suite.faker.Project.Path(3)
	referenceProject := *objects.BasicProject_WithName(referenceProjectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject.Specifications.Path = []string{referenceProjectPath}
	referenceProject.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject.Specifications.ProjectIdentifier.Group = referenceProjectGroupName

	_, err := suite.service.Create(referenceProject, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName, structs.Metadata{}), referenceProject.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ProjectServiceOneSinglePathGroupProjectReferenceTestSuite) Test_CreateBasicProject_MultipleReferences_WithMultiplePath() {

	referenceProjectGroupName := suite.faker.Project.Group()
	referenceProjectGroupPath := suite.faker.Project.Path(1)
	referenceProjectName1 := suite.faker.Project.Name()
	referenceProjectPath1 := suite.faker.Project.Path(3)
	referenceProject1 := *objects.BasicProject_WithName(referenceProjectName1, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject1.Specifications.Path = []string{referenceProjectPath1}
	referenceProject1.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject1.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err := suite.service.Create(referenceProject1, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	referenceProjectName2 := suite.faker.Project.Name()
	referenceProjectPath2 := suite.faker.Project.Path(3)
	referenceProject2 := *objects.BasicProject_WithName(referenceProjectName2, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	referenceProject2.Specifications.Path = []string{referenceProjectPath2}
	referenceProject2.Specifications.GroupObject = group.NewGroupSpecification(0, referenceProjectGroupName, referenceProjectGroupPath, []string{})
	referenceProject2.Specifications.ProjectIdentifier.Group = referenceProjectGroupName
	_, err = suite.service.Create(referenceProject2, true)
	require.NoError(suite.T(), err, "Failed to save reference project")

	projectName := suite.faker.Project.Name()
	project := *objects.BasicProject_WithName(projectName, models.ProjectTypes.Library, models.PlatformTypes.Dotnet, models.RuntimeTypes.Dotnet, suite.workspace)
	project.Specifications.Configuration.References = []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName1, structs.Metadata{}), referenceProject1.Specifications),
		applicationproject.NewProjectBaseStruct(projectStruct.NewHeader(structs.StructTypes.Project, projectStruct.StructKinds.Application, referenceProjectName2, structs.Metadata{}), referenceProject2.Specifications),
	}

	temp, err := suite.service.Create(project, true)
	require.NoError(suite.T(), err, "Failed to save project")

	existingProject, err := suite.service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
	require.NoError(suite.T(), err, "Failed to retrieve project")
	assert.Equal(suite.T(), project, *existingProject)

	validateProjectDependencies, err := suite.service.ValidateProjectReferences(project.Specifications)
	require.NoError(suite.T(), err, "Failed to validate project structure")
	assert.Equal(suite.T(), project, *temp)
	assert.True(suite.T(), validateProjectDependencies)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(project.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject1.GetFullName(), project.Specifications.Workspace, true, true)
			suite.service.Remove(referenceProject2.GetFullName(), project.Specifications.Workspace, true, true)
		}
	})
}
func TestProjectServiceOneSinglePathGroupProjectReferenceTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectServiceOneSinglePathGroupProjectReferenceTestSuite))
}
