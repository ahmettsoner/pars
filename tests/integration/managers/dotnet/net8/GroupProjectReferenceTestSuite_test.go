package net8

import (
	"os"
	"testing"

	"parsdevkit.net/application"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/platforms/core"
	"parsdevkit.net/platforms/dotnet/managers"

	test "pars/tests/internal/testenv"
	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GroupProjectReferenceTestSuite struct {
	suite.Suite
	testArea      string
	environment   string
	workspace     string
	group         string
	manager       core.ApplicationPlatformManagerInterface
	projects      []application_project_payload_structs.ProjectBaseStruct
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupProjectReferenceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.workspace = suite.faker.Workspace.Name()
	suite.group = suite.faker.Project.Group()

	tempWorkingDir, err := test.CreateTempTestDirectory(testArea)
	require.NoError(suite.T(), err, "Create temporary directory failed")
	suite.testArea = tempWorkingDir
	suite.T().Logf("Creating test location at (%v)", suite.testArea)

	suite.manager = managers.NewDotnetManager()

	suite.T().Logf("Initializing New Workspace (%v)", suite.workspace)
	InitializeNewWorkspace(suite.T(), suite.testArea, suite.workspace, suite.environment)

	projectName1 := suite.faker.Project.Name()
	projectName2 := suite.faker.Project.Name()
	CreateNewTestProjectGroupAndPath(suite.T(), suite.group, suite.group, suite.testArea, suite.workspace)
	project1 := CreateNewTestProjectWithGroup(suite.T(), projectName1, suite.testArea, suite.workspace, suite.group, suite.group)
	project2 := CreateNewTestProjectWithGroup(suite.T(), projectName2, suite.testArea, suite.workspace, suite.group, suite.group)

	suite.projects = append(suite.projects, project1, project2)
	suite.T().Log("Project creation completed")
}
func (suite *GroupProjectReferenceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
		os.RemoveAll(suite.testArea)
		os.Remove(application.GetDBLocation(suite.environment))
	}
}

func (suite *GroupProjectReferenceTestSuite) SetupTest() {
}
func (suite *GroupProjectReferenceTestSuite) TearDownTest() {
}

// func (suite *GroupProjectReferenceTestSuite) Test_AddNewReferences_UnGroupedProject() {

// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProject(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ValidateReferences_UnGroupedProject() {

// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProject(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		for _, projectReference := range groupProject.Configuration.References {
// 			referenceState, err := suite.manager.HasReferenceOnProject(groupProject, projectReference)
// 			require.NoError(suite.T(), err, "failed to check reference on project")
// 			assert.True(suite.T(), referenceState)
// 		}
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ListReferences_UnGroupedProject() {

// 	projectName := suite.faker.Project.Name()
// 	for index, groupProject := range suite.projects {
// 		referenceProject := CreateNewTestProject(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		references, err := suite.manager.ListReferencesFromProject(groupProject)
// 		require.NoError(suite.T(), err, "failed to list references")

// 		assert.GreaterOrEqual(suite.T(), len(references), len(groupProject.Configuration.References))
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_AddNewReferences_OneGroupProject() {

// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, suite.group, suite.group)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ValidateReferences_OneGroupProject() {

// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, suite.group, suite.group)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		for _, projectReference := range groupProject.Configuration.References {
// 			referenceState, err := suite.manager.HasReferenceOnProject(groupProject, projectReference)
// 			require.NoError(suite.T(), err, "failed to check reference on project")
// 			assert.True(suite.T(), referenceState)
// 		}
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ListReferences_OneGroupProject() {

// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, suite.group, suite.group)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		references, err := suite.manager.ListReferencesFromProject(groupProject)
// 		require.NoError(suite.T(), err, "failed to list references")

// 		assert.GreaterOrEqual(suite.T(), len(references), len(groupProject.Configuration.References))
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_AddNewReferences_DifferentGroupProject() {

// 	groupName := suite.faker.Project.Group()
// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, groupName, groupName)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ValidateReferences_DifferentGroupProject() {

// 	groupName := suite.faker.Project.Group()
// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, groupName, groupName)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		for _, projectReference := range groupProject.Configuration.References {
// 			referenceState, err := suite.manager.HasReferenceOnProject(groupProject, projectReference)
// 			require.NoError(suite.T(), err, "failed to check reference on project")
// 			assert.True(suite.T(), referenceState)
// 		}
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

// func (suite *GroupProjectReferenceTestSuite) Test_ListReferences_DifferentGroupProject() {

// 	groupName := suite.faker.Project.Group()
// 	for index, groupProject := range suite.projects {
// 		projectName := suite.faker.Project.Name()
// 		referenceProject := CreateNewTestProjectWithGroup(suite.T(), fmt.Sprintf("%v_%v", projectName, index), suite.testArea, suite.workspace, groupName, groupName)
// 		newReferences := []application_project_payload_structs.ProjectSpecification{
// 			referenceProject,
// 		}

// 		groupProject.Configuration.References = append(groupProject.Configuration.References, newReferences...)
// 		err := suite.manager.AddReferenceToProject(groupProject, newReferences)
// 		require.NoError(suite.T(), err, "failed to add references")

// 		references, err := suite.manager.ListReferencesFromProject(groupProject)
// 		require.NoError(suite.T(), err, "failed to list references")

// 		assert.GreaterOrEqual(suite.T(), len(references), len(groupProject.Configuration.References))
// 	}

// 	suite.T().Cleanup(func() {
// 		suite.T().Logf("Test (%v) completed successfully at %v", suite.T().Name(), suite.testArea)
// 	})
// }

func TestGroupProjectReferenceTestSuite(t *testing.T) {
	suite.Run(t, new(GroupProjectReferenceTestSuite))
}
