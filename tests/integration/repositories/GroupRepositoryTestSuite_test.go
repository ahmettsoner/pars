package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GroupRepositoryTestSuite struct {
	suite.Suite
	environment   string
	repository    repositories.GroupRepository
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *GroupRepositoryTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.repository = *repositories.NewGroupRepository(suite.environment)

	suite.T().Log("Group creation completed")
}
func (suite *GroupRepositoryTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *GroupRepositoryTestSuite) SetupTest() {
}
func (suite *GroupRepositoryTestSuite) TearDownTest() {
}

func (suite *GroupRepositoryTestSuite) Test_CreateGroup() {

	groupName := suite.faker.Project.Group()
	groupEntity, _, err := CreateNewSampleGroup(groupName)
	require.NoError(suite.T(), err, "Group creation failed")

	err = suite.repository.Save(groupEntity)
	require.NoError(suite.T(), err, "Failed to save group")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(groupEntity)
		}
	})
}

func (suite *GroupRepositoryTestSuite) Test_GetByName() {

	groupName := suite.faker.Project.Group()
	groupEntity, groupStruct, err := CreateNewSampleGroup(groupName)
	require.NoError(suite.T(), err, "Group creation failed")

	err = suite.repository.Save(groupEntity)
	require.NoError(suite.T(), err, "Failed to save group")

	existingGroup, err := suite.repository.GetByName(groupName)
	require.NoError(suite.T(), err, "Failed to retrieve group by name")

	groupStructFromDB := &group.GroupBaseStruct{}
	err = json.Unmarshal([]byte(existingGroup.Document), groupStructFromDB)
	require.NoError(suite.T(), err, "Failed unmarshal group entity")

	assert.Equal(suite.T(), groupStruct, groupStructFromDB)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(groupEntity)
		}
	})
}

func (suite *GroupRepositoryTestSuite) Test_ListByPath() {

	groupName1 := suite.faker.Project.Group()
	groupPath := suite.faker.Project.Path(1)
	groupEntity1, groupStruct1, err := CreateNewSampleGroupWithSet(groupName1, groupPath)
	require.NoError(suite.T(), err, "Group creation failed")

	suite.repository.Save(groupEntity1)
	require.NoError(suite.T(), err, "Failed to save group")

	groupName2 := suite.faker.Project.Group()
	groupEntity2, groupStruct2, err := CreateNewSampleGroupWithSet(groupName2, groupPath)
	require.NoError(suite.T(), err, "Group creation failed")

	suite.repository.Save(groupEntity2)
	require.NoError(suite.T(), err, "Failed to save group")

	existingGroups, err := suite.repository.ListByPath(groupPath)
	require.NoError(suite.T(), err, "Failed to list groups by set")
	assert.Equal(suite.T(), 2, len(*existingGroups))

	for _, entity := range *existingGroups {
		groupStructFromDB := &group.GroupBaseStruct{}
		err = json.Unmarshal([]byte(entity.Document), groupStructFromDB)
		require.NoError(suite.T(), err, "Failed unmarshal group entity")

		if entity.Name == groupName1 {
			assert.Equal(suite.T(), groupStruct1, groupStructFromDB)
		} else if entity.Name == groupName2 {
			assert.Equal(suite.T(), groupStruct2, groupStructFromDB)
		}
	}

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(groupEntity1)
			suite.repository.Delete(groupEntity2)
		}
	})
}

func TestGroupRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(GroupRepositoryTestSuite))
}

func CreateNewSampleGroup(name string) (*entities.Group, *group.GroupBaseStruct, error) {

	group := BasicGroup_WithName(name)
	jsonData, err := json.Marshal(group)
	if err != nil {
		return nil, nil, err
	}

	groupEntity := entities.Group{
		Name:     name,
		Document: string(jsonData),
	}

	return &groupEntity, group, nil
}

func CreateNewSampleGroupWithSet(name, set string) (*entities.Group, *group.GroupBaseStruct, error) {

	group := BasicGroup_WithNamePath(name, set)

	jsonData, err := json.Marshal(group)
	if err != nil {
		return nil, nil, err
	}

	groupEntity := entities.Group{
		Name:     name,
		Document: string(jsonData),
	}

	return &groupEntity, group, nil
}

func BasicGroup_WithName(name string) *group.GroupBaseStruct {

	group := group.NewGroupBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Group,
			name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		group.NewGroupSpecification(0,
			name,
			"path",
			[]string{"foo", "bar"},
		),
	)
	return &group
}

func BasicGroup_WithNamePath(name, path string) *group.GroupBaseStruct {

	group := group.NewGroupBaseStruct(
		structs.NewHeader(
			structs.StructTypes.Group,
			name,
			structs.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		group.NewGroupSpecification(0,
			name,
			path,
			[]string{"foo", "bar"},
		),
	)

	return &group
}
