package services

import (
	"testing"

	"parsdevkit.net/application"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	workspaceWorkspace "parsdevkit.net/modules/workspace/basic_workspace"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"parsdevkit.net/application/schemas"
)

type WorkspaceServiceTestSuite struct {
	suite.Suite
	service       basic_workspace_contract.WorkspaceInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *WorkspaceServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = workspaceWorkspace.NewWorkspaceService(suite.environment)

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
			suite.service.Remove(workspace.Header.Name, true, true)
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
			suite.service.Remove(workspace.Header.Name, true, true)
		}
	})
}

func TestWorkspaceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WorkspaceServiceTestSuite))
}

func BasicWorkspace_WithName(name string) *basic_workspace_payload_structs.WorkspaceBaseStruct {

	workspace := basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		basic_workspace_payload_structs.NewWorkspaceSpecification(0,
			name,
			"path",
		),
	)
	return &workspace
}

func BasicWorkspace_WithSpecification(specifications basic_workspace_payload_structs.WorkspaceSpecification) *basic_workspace_payload_structs.WorkspaceBaseStruct {

	workspace := basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			specifications.Name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		specifications,
	)
	return &workspace
}

func BasicWorkspace_WithNamePath(name, path string) *basic_workspace_payload_structs.WorkspaceBaseStruct {

	workspace := basic_workspace_payload_structs.NewWorkspaceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			"",
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		basic_workspace_payload_structs.NewWorkspaceSpecification(0,
			name,
			path,
		),
	)

	return &workspace
}
