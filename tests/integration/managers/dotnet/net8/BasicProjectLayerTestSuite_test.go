package net8

import (
	"os"
	"testing"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/dotnet/managers"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test"
	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BasicProjectLayerTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *BasicProjectLayerTestSuite) SetupSuite() {

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
}
func (suite *BasicProjectLayerTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *BasicProjectLayerTestSuite) SetupTest() {
}
func (suite *BasicProjectLayerTestSuite) TearDownTest() {
}

// func (suite *BasicProjectLayerTestSuite) Test_AddNewLayers_Basic() {
// 	projectName := suite.faker.Project.Name()
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 	}

// 	testProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
// 	suite.T().Log("Project creation completed")

// 	testProject.Configuration.Layers = append(testProject.Configuration.Layers, newLayers...)
// 	err := suite.manager.CreateLayerFolder(testProject, newLayers...)
// 	require.NoError(suite.T(), err, "Could not create layer folder")

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *BasicProjectLayerTestSuite) Test_AddNewLayers_InnerLayers() {
// 	projectName := suite.faker.Project.Name()
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 		applicationproject.NewLayer_Empty_Basic("sample-layer:sub", "sample-layer/sub-layer"),
// 		applicationproject.NewLayer_Empty_Basic("sample-layer:sub:deep", "sample-layer/sub-layer/deep-layer"),
// 	}
// 	testProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
// 	suite.T().Log("Project creation completed")

// 	testProject.Configuration.Layers = append(testProject.Configuration.Layers, newLayers...)
// 	err := suite.manager.CreateLayerFolder(testProject, newLayers...)
// 	require.NoError(suite.T(), err, "CreateLayerFolder returned an error for valid layers")

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *BasicProjectLayerTestSuite) Test_ValidateLayers_Basic() {

// 	projectName := suite.faker.Project.Name()
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 	}

// 	testProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
// 	suite.T().Log("Project creation completed")

// 	testProject.Configuration.Layers = append(testProject.Configuration.Layers, newLayers...)
// 	err := suite.manager.CreateLayerFolder(testProject, newLayers...)
// 	require.NoError(suite.T(), err, "CreateLayerFolder returned an error for valid layers")

// 	for _, projectLayer := range testProject.Configuration.Layers {
// 		layerFolderState, folderErr := suite.manager.IsLayerFolderExists(testProject, projectLayer.Name)
// 		layerState, err := suite.manager.HasLayerOnProject(testProject, projectLayer.Name)

// 		require.NoError(suite.T(), folderErr, "HasLayerOnProject returned an error for existing layer "+projectLayer.Name)
// 		require.NoError(suite.T(), err, "Failed to validate layer in the project configuration")
// 		assert.True(suite.T(), layerFolderState)
// 		assert.True(suite.T(), layerState)
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *BasicProjectLayerTestSuite) Test_ValidateLayers_InnerLayers() {
// 	projectName := suite.faker.Project.Name()
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 		applicationproject.NewLayer_Empty_Basic("sample-layer:sub", "sample-layer/sub-layer"),
// 		applicationproject.NewLayer_Empty_Basic("sample-layer:sub:deep", "sample-layer/sub-layer/deep-layer"),
// 	}
// 	testProject := CreateNewTestProject(suite.T(), projectName, suite.testArea, suite.workspace)
// 	suite.T().Log("Project creation completed")

// 	testProject.Configuration.Layers = append(testProject.Configuration.Layers, newLayers...)
// 	err := suite.manager.CreateLayerFolder(testProject, newLayers...)
// 	require.NoError(suite.T(), err, "CreateLayerFolder returned an error for valid layers")

// 	for _, projectLayer := range testProject.Configuration.Layers {
// 		layerFolderState, folderErr := suite.manager.IsLayerFolderExists(testProject, projectLayer.Name)
// 		layerState, err := suite.manager.HasLayerOnProject(testProject, projectLayer.Name)

// 		require.NoError(suite.T(), folderErr, "HasLayerOnProject returned an error for existing layer "+projectLayer.Name)
// 		require.NoError(suite.T(), err, "Failed to validate layer in the project configuration")
// 		assert.True(suite.T(), layerFolderState)
// 		assert.True(suite.T(), layerState)
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *BasicProjectLayerTestSuite) Test_ValidateProjectCreateWithLayers_Basic() {

// 	projectName := suite.faker.Project.Name()
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 	}

// 	testProject := CreateNewTestProjectWithLayer(suite.T(), projectName, suite.testArea, suite.workspace, newLayers)
// 	suite.T().Log("Project creation completed")

// 	for _, projectLayer := range testProject.Configuration.Layers {
// 		layerFolderState, folderErr := suite.manager.IsLayerFolderExists(testProject, projectLayer.Name)
// 		layerState, err := suite.manager.HasLayerOnProject(testProject, projectLayer.Name)

// 		require.NoError(suite.T(), folderErr, "HasLayerOnProject returned an error for existing layer "+projectLayer.Name)
// 		require.NoError(suite.T(), err, "Failed to validate layer in the project configuration")
// 		assert.True(suite.T(), layerFolderState)
// 		assert.True(suite.T(), layerState)
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

func TestBasicProjectLayerTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProjectLayerTestSuite))
}
