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

type BasicProjectPackageTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	project       applicationproject.ProjectSpecification
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *BasicProjectPackageTestSuite) SetupSuite() {

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
func (suite *BasicProjectPackageTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *BasicProjectPackageTestSuite) SetupTest() {
}
func (suite *BasicProjectPackageTestSuite) TearDownTest() {
}

func (suite *BasicProjectPackageTestSuite) Test_AddNewPackages_WithoutVersion() {
	newPackages := GetPackages(0, 1, false)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectPackageTestSuite) Test_AddNewPackages_WithVersion() {
	newPackages := GetPackages(1, 1, true)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectPackageTestSuite) Test_ValidatePackages_WithoutVersion() {

	newPackages := GetPackages(2, 1, false)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	for _, projectPackage := range suite.project.Configuration.Dependencies {
		packageState, err := suite.manager.HasPackageOnProject(suite.project, projectPackage)
		require.NoError(suite.T(), err, "failed to validate package on the project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectPackageTestSuite) Test_ValidatePackages_WithVersion() {

	newPackages := GetPackages(3, 1, true)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	for _, projectPackage := range suite.project.Configuration.Dependencies {
		packageState, err := suite.manager.HasPackageOnProject(suite.project, projectPackage)
		require.NoError(suite.T(), err, "failed to check package on project")
		assert.True(suite.T(), packageState)
	}

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectPackageTestSuite) Test_ListPackages_WithoutVersion() {

	newPackages := GetPackages(4, 1, false)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	packages, err := suite.manager.ListPackagesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to list packages")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func (suite *BasicProjectPackageTestSuite) Test_ListPackage_WithVersion() {

	newPackages := GetPackages(5, 1, true)
	suite.project.Configuration.Dependencies = append(suite.project.Configuration.Dependencies, newPackages...)
	err := suite.manager.AddPackageToProject(suite.project, newPackages)
	require.NoError(suite.T(), err, "failed to add packages")

	packages, err := suite.manager.ListPackagesFromProject(suite.project)
	require.NoError(suite.T(), err, "failed to list packages")

	assert.GreaterOrEqual(suite.T(), len(packages), len(suite.project.Configuration.Dependencies))

	suite.T().Cleanup(func() {
		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
	})
}

func TestBasicProjectPackageTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProjectPackageTestSuite))
}
