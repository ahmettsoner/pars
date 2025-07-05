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

type GroupProjectPackageTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	projects      []applicationproject.ProjectBaseStruct
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupProjectPackageTestSuite) SetupSuite() {

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

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	groupName := suite.faker.Project.Group()
	_ = CreateNewTestProjectGroupAndPath(suite.T(), groupName, groupName, suite.testArea, suite.workspace)
	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)

	suite.projects = append(suite.projects, project1, project2)
	suite.T().Log("Project creation completed")
}
func (suite *GroupProjectPackageTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *GroupProjectPackageTestSuite) SetupTest() {
}
func (suite *GroupProjectPackageTestSuite) TearDownTest() {
}

func (suite *GroupProjectPackageTestSuite) Test_AddNewPackages_WithoutVersion() {
	for _, groupProject := range suite.projects {
		newPackages := GetPackages(0, 1, false)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_AddNewPackages_WithVersion() {
	for _, groupProject := range suite.projects {
		newPackages := GetPackages(1, 1, true)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ValidatePackages_WithoutVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(2, 1, false)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		for _, projectPackage := range groupProject.Specifications.Configuration.Dependencies {
			packageState, err := suite.manager.HasDependencyOnProject(groupProject, projectPackage)
			require.NoError(suite.T(), err, "failed to check package on project")
			assert.True(suite.T(), packageState)
		}
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ValidatePackages_WithVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(3, 1, true)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		for _, projectPackage := range groupProject.Specifications.Configuration.Dependencies {
			packageState, err := suite.manager.HasDependencyOnProject(groupProject, projectPackage)
			require.NoError(suite.T(), err, "failed to check package on project")
			assert.True(suite.T(), packageState)
		}
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ListPackages_WithoutVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(4, 1, false)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		packages, err := suite.manager.ListDependenciesFromProject(groupProject)
		require.NoError(suite.T(), err, "failed to list packages")

		assert.GreaterOrEqual(suite.T(), len(packages), len(groupProject.Specifications.Configuration.Dependencies))
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ListPackages_WithVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(5, 1, true)

		groupProject.Specifications.Configuration.Dependencies = append(groupProject.Specifications.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddDependenciesToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		packages, err := suite.manager.ListDependenciesFromProject(groupProject)
		require.NoError(suite.T(), err, "failed to list packages")

		assert.GreaterOrEqual(suite.T(), len(packages), len(groupProject.Specifications.Configuration.Dependencies))
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}
func TestGroupProjectPackageTestSuite(t *testing.T) {
	suite.Run(t, new(GroupProjectPackageTestSuite))
}
