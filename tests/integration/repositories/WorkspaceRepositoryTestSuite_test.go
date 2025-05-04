package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WorkspaceRepositoryTestSuite struct {
	suite.Suite
	environment   string
	repository    repositories.WorkspaceRepository
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *WorkspaceRepositoryTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.repository = *repositories.NewWorkspaceRepository(suite.environment)

	suite.T().Log("Project creation completed")
}
func (suite *WorkspaceRepositoryTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *WorkspaceRepositoryTestSuite) SetupTest() {
}
func (suite *WorkspaceRepositoryTestSuite) TearDownTest() {
}

func (suite *WorkspaceRepositoryTestSuite) Test_CreateWorkspace() {

	workspaceName := suite.faker.Workspace.Name()
	workspaceEntity, _, err := CreateNewSampleWorkspace(workspaceName)
	require.NoError(suite.T(), err, "Workspace creation failed")

	err = suite.repository.Save(workspaceEntity)
	require.NoError(suite.T(), err, "Failed to save workspace")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(workspaceEntity)
		}
	})
}

func (suite *WorkspaceRepositoryTestSuite) Test_GetByName() {

	workspaceName := suite.faker.Workspace.Name()
	workspaceEntity, workspaceStruct, err := CreateNewSampleWorkspace(workspaceName)
	require.NoError(suite.T(), err, "Workspace creation failed")

	err = suite.repository.Save(workspaceEntity)
	require.NoError(suite.T(), err, "Failed to save workspace")

	existingWorkspace, err := suite.repository.GetByName(workspaceName)
	require.NoError(suite.T(), err, "Failed to retrieve workspace by name")

	workspaceStructFromDB := &workspace.WorkspaceBaseStruct{}
	err = json.Unmarshal([]byte(existingWorkspace.Document), workspaceStructFromDB)
	require.NoError(suite.T(), err, "Failed unmarshal workspace entity")

	assert.Equal(suite.T(), workspaceStruct, workspaceStructFromDB)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(workspaceEntity)
		}
	})
}

func (suite *WorkspaceRepositoryTestSuite) Test_ListByPath() {

	workspaceName1 := suite.faker.Workspace.Name()
	workspaceName2 := suite.faker.Workspace.Name()
	workspacePath := suite.faker.Project.Path(1)
	workspaceEntity1, workspaceStruct1, err := CreateNewSampleWorkspaceWithSet(workspaceName1, workspacePath)
	require.NoError(suite.T(), err, "Workspace creation failed")

	suite.repository.Save(workspaceEntity1)
	require.NoError(suite.T(), err, "Failed to save workspace")

	workspaceEntity2, workspaceStruct2, err := CreateNewSampleWorkspaceWithSet(workspaceName2, workspacePath)
	require.NoError(suite.T(), err, "Workspace creation failed")

	suite.repository.Save(workspaceEntity2)
	require.NoError(suite.T(), err, "Failed to save workspace")

	existingWorkspaces, err := suite.repository.ListByPath(workspacePath)
	require.NoError(suite.T(), err, "Failed to list workspaces by set")
	assert.Equal(suite.T(), 2, len(*existingWorkspaces))

	for _, entity := range *existingWorkspaces {
		workspaceStructFromDB := &workspace.WorkspaceBaseStruct{}
		err = json.Unmarshal([]byte(entity.Document), workspaceStructFromDB)
		require.NoError(suite.T(), err, "Failed unmarshal workspace entity")

		if entity.Name == "sample_workspace4" {
			assert.Equal(suite.T(), workspaceStruct1, workspaceStructFromDB)
		} else if entity.Name == "sample_workspace5" {
			assert.Equal(suite.T(), workspaceStruct2, workspaceStructFromDB)
		}
	}

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(workspaceEntity1)
			suite.repository.Delete(workspaceEntity2)
		}
	})
}

func TestWorkspaceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(WorkspaceRepositoryTestSuite))
}

func CreateNewSampleWorkspace(name string) (*entities.Workspace, *workspace.WorkspaceBaseStruct, error) {

	workspace := BasicWorkspace_WithName(name)
	jsonData, err := json.Marshal(workspace)
	if err != nil {
		return nil, nil, err
	}

	workspaceEntity := entities.Workspace{
		Name:     name,
		Document: string(jsonData),
	}

	return &workspaceEntity, workspace, nil
}

func CreateNewSampleWorkspaceWithSet(name, set string) (*entities.Workspace, *workspace.WorkspaceBaseStruct, error) {

	workspace := BasicWorkspace_WithNamePath(name, set)

	jsonData, err := json.Marshal(workspace)
	if err != nil {
		return nil, nil, err
	}

	workspaceEntity := entities.Workspace{
		Name:     name,
		Document: string(jsonData),
	}

	return &workspaceEntity, workspace, nil
}

func BasicWorkspace_WithName(name string) *workspace.WorkspaceBaseStruct {

	workspace := workspace.NewWorkspaceBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Workspace,
			name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		workspace.NewWorkspaceSpecification(0,
			name,
			"path",
		),
	)
	return &workspace
}

func BasicWorkspace_WithNamePath(name, path string) *workspace.WorkspaceBaseStruct {

	workspace := workspace.NewWorkspaceBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Workspace,
			name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		workspace.NewWorkspaceSpecification(0,
			name,
			path,
		),
	)

	return &workspace
}
