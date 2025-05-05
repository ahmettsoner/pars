package net8

import (
	"os"
	"testing"

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

type GroupProjectPackageTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	projects      []applicationproject.ProjectSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupProjectPackageTestSuite) SetupSuite() {

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
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *GroupProjectPackageTestSuite) SetupTest() {
}
func (suite *GroupProjectPackageTestSuite) TearDownTest() {
}

func (suite *GroupProjectPackageTestSuite) Test_AddNewPackages_WithoutVersion() {
	for _, groupProject := range suite.projects {
		newPackages := GetPackages(0, 1, false)

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_AddNewPackages_WithVersion() {
	for _, groupProject := range suite.projects {
		newPackages := GetPackages(1, 1, true)

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ValidatePackages_WithoutVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(2, 1, false)

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		for _, projectPackage := range groupProject.Configuration.Dependencies {
			packageState, err := suite.manager.HasPackageOnProject(groupProject, projectPackage)
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

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		for _, projectPackage := range groupProject.Configuration.Dependencies {
			packageState, err := suite.manager.HasPackageOnProject(groupProject, projectPackage)
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

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		packages, err := suite.manager.ListPackagesFromProject(groupProject)
		require.NoError(suite.T(), err, "failed to list packages")

		assert.GreaterOrEqual(suite.T(), len(packages), len(groupProject.Configuration.Dependencies))
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *GroupProjectPackageTestSuite) Test_ListPackages_WithVersion() {

	for _, groupProject := range suite.projects {
		newPackages := GetPackages(5, 1, true)

		groupProject.Configuration.Dependencies = append(groupProject.Configuration.Dependencies, newPackages...)
		err := suite.manager.AddPackageToProject(groupProject, newPackages)
		require.NoError(suite.T(), err, "failed to add packages")

		packages, err := suite.manager.ListPackagesFromProject(groupProject)
		require.NoError(suite.T(), err, "failed to list packages")

		assert.GreaterOrEqual(suite.T(), len(packages), len(groupProject.Configuration.Dependencies))
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}
func TestGroupProjectPackageTestSuite(t *testing.T) {
	suite.Run(t, new(GroupProjectPackageTestSuite))
}
