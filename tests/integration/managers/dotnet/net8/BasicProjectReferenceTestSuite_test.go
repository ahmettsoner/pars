package net8

import (
	"os"
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/project"
	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/dotnet/managers"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BasicProjectReferenceTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	project       applicationproject.ProjectSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *BasicProjectReferenceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.manager = managers.NewDotnetManager()

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)

	projectName := suite.faker.Project.Name()
	suite.project = CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
	suite.T().Log("Project creation completed")
}
func (suite *BasicProjectReferenceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *BasicProjectReferenceTestSuite) SetupTest() {
}
func (suite *BasicProjectReferenceTestSuite) TearDownTest() {
}

func (suite *BasicProjectReferenceTestSuite) Test_AddNewReferences() {
	projectName := suite.faker.Project.Name()
	referenceProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName, structs.Metadata{}),
			referenceProject,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)

	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectReferenceTestSuite) Test_ListReferences() {

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	referenceProject := CreateNewTestProject(suite.T(), projectName1, suite.testArea, suite.workspace)
	referenceProject2 := CreateNewTestProject(suite.T(), projectName2, suite.testArea, suite.workspace)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName1, structs.Metadata{}),
			referenceProject,
		),
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName2, structs.Metadata{}),
			referenceProject2,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)
	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	packages, err := suite.manager.ListReferencesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to add reference to project")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectReferenceTestSuite) Test_ValidateReferences() {

	projectName := suite.faker.Project.Name()
	referenceProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName, structs.Metadata{}),
			referenceProject,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)
	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	for _, projectReference := range suite.project.Configuration.References {
		packageState, err := suite.manager.HasReferenceOnProject(suite.project, projectReference.Specifications)
		require.NoError(suite.T(), err, "failed to validate package on project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectReferenceTestSuite) Test_AddNewReferences_GroupedProject() {

	projectName := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()
	CreateNewTestProjectGroupAndPath(suite.T(), groupName, groupName, suite.testArea, suite.workspace)
	referenceProject := CreateNewTestProjectWithGroup(suite.T(), projectName, suite.testArea, suite.workspace, groupName, groupName)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName, structs.Metadata{}),
			referenceProject,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)
	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectReferenceTestSuite) Test_ListReferences_GroupedProject() {

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()
	CreateNewTestProjectGroupAndPath(suite.T(), groupName, groupName, suite.testArea, suite.workspace)
	referenceProject := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
	referenceProject2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName1, structs.Metadata{}),
			referenceProject,
		),
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName2, structs.Metadata{}),
			referenceProject2,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)
	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	packages, err := suite.manager.ListReferencesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to list references")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectReferenceTestSuite) Test_ValidateReferences_GroupedProject() {

	projectName := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()
	CreateNewTestProjectGroupAndPath(suite.T(), groupName, groupName, suite.testArea, suite.workspace)
	referenceProject := CreateNewTestProjectWithGroup(suite.T(), projectName, suite.testArea, suite.workspace, groupName, groupName)
	newReferences := []applicationproject.ProjectBaseStruct{
		applicationproject.NewProjectBaseStruct(
			project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, projectName, structs.Metadata{}),
			referenceProject,
		),
	}
	suite.project.Configuration.References = append(suite.project.Configuration.References, newReferences...)
	newReferencesSpecs := make([]applicationproject.ProjectSpecification, 0)
	for _, ref := range newReferences {
		newReferencesSpecs = append(newReferencesSpecs, ref.Specifications)
	}
	err := suite.manager.AddReferenceToProject(suite.project, newReferencesSpecs)
	require.NoError(suite.T(), err, "failed to add reference to project")

	for _, projectReference := range suite.project.Configuration.References {
		packageState, err := suite.manager.HasReferenceOnProject(suite.project, projectReference.Specifications)
		require.NoError(suite.T(), err, "failed to validate package on project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestBasicProjectReferenceTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProjectReferenceTestSuite))
}
