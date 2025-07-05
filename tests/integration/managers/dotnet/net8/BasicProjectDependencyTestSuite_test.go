package net8

import (
	"os"
	"testing"

	"parsdevkit.net/application"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/dotnet/managers"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BasicProjectDependencyTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ApplicationPlatformManagerInterface
	project       applicationproject.ProjectBaseStruct
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *BasicProjectDependencyTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
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
func (suite *BasicProjectDependencyTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *BasicProjectDependencyTestSuite) SetupTest() {
}
func (suite *BasicProjectDependencyTestSuite) TearDownTest() {
}

func (suite *BasicProjectDependencyTestSuite) Test_AddNewDependencies_WithoutVersion() {
	newDependencies := GetDependencies(0, 1, false)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectDependencyTestSuite) Test_AddNewDependencies_WithVersion() {
	newDependencies := GetDependencies(1, 1, true)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectDependencyTestSuite) Test_ValidateDependencies_WithoutVersion() {

	newDependencies := GetDependencies(2, 1, false)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	for _, projectDependency := range suite.project.Specifications.Configuration.Dependencies {
		packageState, err := suite.manager.HasDependencyOnProject(suite.project, projectDependency)
		require.NoError(suite.T(), err, "failed to validate package on the project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectDependencyTestSuite) Test_ValidateDependencies_WithVersion() {

	newDependencies := GetDependencies(3, 1, true)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	for _, projectDependency := range suite.project.Specifications.Configuration.Dependencies {
		packageState, err := suite.manager.HasDependencyOnProject(suite.project, projectDependency)
		require.NoError(suite.T(), err, "failed to check package on project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectDependencyTestSuite) Test_ListDependencies_WithoutVersion() {

	newDependencies := GetDependencies(4, 1, false)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	packages, err := suite.manager.ListDependenciesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to list packages")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Specifications.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectDependencyTestSuite) Test_ListDependency_WithVersion() {

	newDependencies := GetDependencies(5, 1, true)
	suite.project.Specifications.Configuration.Dependencies = append(suite.project.Specifications.Configuration.Dependencies, newDependencies...)
	err := suite.manager.AddDependenciesToProject(suite.project, newDependencies)
	require.NoError(suite.T(), err, "failed to add packages")

	packages, err := suite.manager.ListDependenciesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to list packages")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Specifications.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestBasicProjectDependencyTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProjectDependencyTestSuite))
}
