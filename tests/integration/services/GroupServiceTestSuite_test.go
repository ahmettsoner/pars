package services

import (
	"testing"

	"parsdevkit.net/application"
	group "parsdevkit.net/modules/group/basic_group"
	"parsdevkit.net/modules/group/basic_group_payload"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"parsdevkit.net/application/contracts"
)

type GroupServiceTestSuite struct {
	suite.Suite
	service       contracts.GroupServiceInterface[basic_group_payload.GroupBaseStruct]
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = group.NewGroupService(suite.environment)

	suite.T().Log("Group creation completed")
}
func (suite *GroupServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *GroupServiceTestSuite) SetupTest() {
}
func (suite *GroupServiceTestSuite) TearDownTest() {
}

func (suite *GroupServiceTestSuite) Test_CreateGroup() {

	groupName := suite.faker.Project.Group()
	groupPath := suite.faker.Project.Group()
	group := *BasicGroup_WithNamePath(groupName, groupPath)

	temp, err := suite.service.Save(group)
	require.NoError(suite.T(), err, "Failed to save group")
	assert.Equal(suite.T(), group, *temp)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(group.Header.Name, true)
		}
	})
}

func (suite *GroupServiceTestSuite) Test_GetByName() {

	groupName := suite.faker.Project.Group()
	groupPath := suite.faker.Project.Group()
	group := *BasicGroup_WithNamePath(groupName, groupPath)

	temp, err := suite.service.Save(group)
	require.NoError(suite.T(), err, "Failed to save group")
	assert.Equal(suite.T(), group, *temp)

	existingGroup, err := suite.service.GetByName(groupName)
	require.NoError(suite.T(), err, "Failed to retrieve group by name")

	assert.Equal(suite.T(), group, *existingGroup)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(group.Header.Name, true)
		}
	})
}

func TestGroupServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GroupServiceTestSuite))
}
