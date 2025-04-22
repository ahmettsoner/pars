package services

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WorkspaceServiceTestSuite struct {
	suite.Suite
	service       services.WorkspaceServiceInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *WorkspaceServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = *services.NewWorkspaceService(suite.environment)

	suite.T().Log("Workspace creation completed")
}
func (suite *WorkspaceServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *WorkspaceServiceTestSuite) SetupTest() {
}
func (suite *WorkspaceServiceTestSuite) TearDownTest() {
}

func (suite *WorkspaceServiceTestSuite) Test_CreateWorkspace() {

	workspaceName := suite.faker.Workspace.Name()
	workspace := *BasicWorkspace_WithName(workspaceName)

	temp, err := suite.service.Save(workspace)
	require.NoError(suite.T(), err, "Failed to save workspace")
	assert.Equal(suite.T(), workspace, *temp)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(workspace.Name, true, true)
		}
	})
}

func (suite *WorkspaceServiceTestSuite) Test_GetByName() {

	workspaceName := suite.faker.Workspace.Name()
	workspace := *BasicWorkspace_WithName(workspaceName)

	temp, err := suite.service.Save(workspace)
	require.NoError(suite.T(), err, "Failed to save workspace")
	assert.Equal(suite.T(), workspace, *temp)

	existingWorkspace, err := suite.service.GetByName(workspaceName)
	require.NoError(suite.T(), err, "Failed to retrieve workspace by name")

	assert.Equal(suite.T(), workspace, *existingWorkspace)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(workspace.Name, true, true)
		}
	})
}

func TestWorkspaceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WorkspaceServiceTestSuite))
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

func BasicWorkspace_WithSpecification(specifications workspace.WorkspaceSpecification) *workspace.WorkspaceBaseStruct {

	workspace := workspace.NewWorkspaceBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Workspace,
			specifications.Name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		specifications,
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
