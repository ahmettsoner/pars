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

type GroupProjectLayerTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	manager       core.ManagerInterface
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupProjectLayerTestSuite) SetupSuite() {

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
func (suite *GroupProjectLayerTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(utils.GetDBLocation(suite.environment))
	}
}

func (suite *GroupProjectLayerTestSuite) SetupTest() {
}
func (suite *GroupProjectLayerTestSuite) TearDownTest() {
}

// func (suite *GroupProjectLayerTestSuite) Test_AddNewLayers_Basic() {

// 	projectName1 := suite.faker.Project.Name()
// 	projectName2 := suite.faker.Project.Name()
// 	groupName := suite.faker.Project.Group()
// 	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
// 	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)

// 	testProjects := make([]applicationproject.ProjectSpecification, 0)
// 	testProjects = append(testProjects, project1, project2)

// 	suite.T().Log("Project creation completed")

// 	for _, groupProject := range testProjects {
// 		newLayers := []applicationproject.Layer{
// 			applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 		}
// 		groupProject.Configuration.Layers = append(groupProject.Configuration.Layers, newLayers...)
// 		err := suite.manager.CreateLayerFolder(groupProject, newLayers...)
// 		require.NoError(suite.T(), err, "failed to create folder")
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectLayerTestSuite) Test_AddNewLayers_InnerLayers() {

// 	projectName1 := suite.faker.Project.Name()
// 	projectName2 := suite.faker.Project.Name()
// 	groupName := suite.faker.Project.Group()
// 	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
// 	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)

// 	testProjects := make([]applicationproject.ProjectSpecification, 0)
// 	testProjects = append(testProjects, project1, project2)

// 	suite.T().Log("Project creation completed")

// 	for _, groupProject := range testProjects {
// 		newLayers := []applicationproject.Layer{
// 			applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 			applicationproject.NewLayer_Empty_Basic("sample-layer:sub", "sample-layer/sub-layer"),
// 			applicationproject.NewLayer_Empty_Basic("sample-layer:sub:deep", "sample-layer/sub-layer/deep-layer"),
// 		}
// 		groupProject.Configuration.Layers = append(groupProject.Configuration.Layers, newLayers...)

// 		groupProject.Configuration.Layers = append(groupProject.Configuration.Layers, newLayers...)
// 		err := suite.manager.CreateLayerFolder(groupProject, newLayers...)
// 		require.NoError(suite.T(), err, "failed to create folder")
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectLayerTestSuite) Test_ValidateLayers_Basic() {

// 	projectName1 := suite.faker.Project.Name()
// 	projectName2 := suite.faker.Project.Name()
// 	groupName := suite.faker.Project.Group()
// 	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
// 	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)

// 	testProjects := make([]applicationproject.ProjectSpecification, 0)
// 	testProjects = append(testProjects, project1, project2)

// 	suite.T().Log("Project creation completed")

// 	for _, groupProject := range testProjects {
// 		newLayers := []applicationproject.Layer{
// 			applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 		}
// 		groupProject.Configuration.Layers = append(groupProject.Configuration.Layers, newLayers...)
// 		err := suite.manager.CreateLayerFolder(groupProject, newLayers...)
// 		require.NoError(suite.T(), err, "failed to create folder")

// 		for _, projectLayer := range groupProject.Configuration.Layers {
// 			layerFolderState, folderErr := suite.manager.IsLayerFolderExists(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), folderErr, "failed to validate layer folder state")

// 			layerState, err := suite.manager.HasLayerOnProject(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), err, "failed to validate layer state")

// 			assert.True(suite.T(), layerFolderState)
// 			assert.True(suite.T(), layerState)
// 		}
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectLayerTestSuite) Test_ValidateLayers_InnerLayers() {

// 	projectName1 := suite.faker.Project.Name()
// 	projectName2 := suite.faker.Project.Name()
// 	groupName := suite.faker.Project.Group()
// 	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName)
// 	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName)

// 	testProjects := make([]applicationproject.ProjectSpecification, 0)
// 	testProjects = append(testProjects, project1, project2)

// 	suite.T().Log("Project creation completed")

// 	for _, groupProject := range testProjects {
// 		newLayers := []applicationproject.Layer{
// 			applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 			applicationproject.NewLayer_Empty_Basic("sample-layer:sub", "sample-layer/sub-layer"),
// 			applicationproject.NewLayer_Empty_Basic("sample-layer:sub:deep", "sample-layer/sub-layer/deep-layer"),
// 		}
// 		groupProject.Configuration.Layers = append(groupProject.Configuration.Layers, newLayers...)
// 		err := suite.manager.CreateLayerFolder(groupProject, newLayers...)
// 		require.NoError(suite.T(), err, "failed to create folder")

// 		for _, projectLayer := range groupProject.Configuration.Layers {
// 			layerFolderState, folderErr := suite.manager.IsLayerFolderExists(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), folderErr, "failed to validate layer folder state")

// 			layerState, err := suite.manager.HasLayerOnProject(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), err, "failed to validate layer state")

// 			assert.True(suite.T(), layerFolderState)
// 			assert.True(suite.T(), layerState)
// 		}
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectLayerTestSuite) Test_ValidateProjectCreateWithLayers_Basic() {
// 	newLayers := []applicationproject.Layer{
// 		applicationproject.NewLayer_Empty_Basic("sample-layer", "sample-layer"),
// 	}

// 	projectName1 := suite.faker.Project.Name()
// 	projectName2 := suite.faker.Project.Name()
// 	groupName := suite.faker.Project.Group()
// 	project1 := CreateNewTestProjectWithGroupAndLayers(suite.T(), projectName1, suite.testArea, suite.workspace, groupName, groupName, newLayers)
// 	project2 := CreateNewTestProjectWithGroupAndLayers(suite.T(), projectName2, suite.testArea, suite.workspace, groupName, groupName, newLayers)

// 	testProjects := make([]applicationproject.ProjectSpecification, 0)
// 	testProjects = append(testProjects, project1, project2)

// 	suite.T().Log("Project creation completed")

// 	for _, groupProject := range testProjects {

// 		for _, projectLayer := range groupProject.Configuration.Layers {
// 			layerFolderState, folderErr := suite.manager.IsLayerFolderExists(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), folderErr, "failed to validate layer folder state")

// 			layerState, err := suite.manager.HasLayerOnProject(groupProject, projectLayer.Name)
// 			require.NoError(suite.T(), err, "failed to validate layer state")

// 			assert.True(suite.T(), layerFolderState)
// 			assert.True(suite.T(), layerState)
// 		}
// 	}

//		suite.T().Cleanup(func() {
//			suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
//		})
//	}
func TestGroupProjectLayerTestSuite(t *testing.T) {
	suite.Run(t, new(GroupProjectLayerTestSuite))
}
