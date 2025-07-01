package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/application/schemas"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/entities"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"parsdevkit.net/persistence/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ResourceRepositoryTestSuite struct {
	suite.Suite
	environment   string
	repository    repositories.ResourceRepository
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ResourceRepositoryTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	dbContext := contexts.NewDbContext(suite.environment)
	suite.repository = *repositories.NewResourceRepository(dbContext)

	suite.T().Log("Resource creation completed")
}
func (suite *ResourceRepositoryTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *ResourceRepositoryTestSuite) SetupTest() {
}
func (suite *ResourceRepositoryTestSuite) TearDownTest() {
}

func (suite *ResourceRepositoryTestSuite) Test_CreateResource() {

	resourceName := suite.faker.Resource.Name()
	resourceEntity, _, err := CreateNewSampleResource(resourceName)
	require.NoError(suite.T(), err, "Resource creation failed")

	err = suite.repository.Save(resourceEntity)
	require.NoError(suite.T(), err, "Failed to save resource")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(resourceEntity)
		}
	})
}

func (suite *ResourceRepositoryTestSuite) Test_GetByName() {

	resourceName := suite.faker.Resource.Name()
	resourceEntity, resourceStruct, err := CreateNewSampleResource(resourceName)
	require.NoError(suite.T(), err, "Resource creation failed")

	err = suite.repository.Save(resourceEntity)
	require.NoError(suite.T(), err, "Failed to save resource")

	existingResource, err := suite.repository.GetByName(resourceName)
	require.NoError(suite.T(), err, "Failed to retrieve resource by name")

	resourceStructFromDB := &objectresource.ResourceBaseStruct{}
	err = json.Unmarshal([]byte(existingResource.Document), resourceStructFromDB)
	require.NoError(suite.T(), err, "Failed unmarshal resource entity")

	assert.Equal(suite.T(), resourceStruct, resourceStructFromDB)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(resourceEntity)
		}
	})
}

func (suite *ResourceRepositoryTestSuite) Test_ListBySet() {

	resourceName1 := suite.faker.Resource.Name()
	resourceName2 := suite.faker.Resource.Name()
	resourceSet := suite.faker.Project.Set()
	resourceEntity1, resourceStruct1, err := CreateNewSampleResourceWithSet(resourceName1, resourceSet)
	require.NoError(suite.T(), err, "Resource creation failed")

	suite.repository.Save(resourceEntity1)
	require.NoError(suite.T(), err, "Failed to save resource")

	resourceEntity2, resourceStruct2, err := CreateNewSampleResourceWithSet(resourceName2, resourceSet)
	require.NoError(suite.T(), err, "Resource creation failed")

	suite.repository.Save(resourceEntity2)
	require.NoError(suite.T(), err, "Failed to save resource")

	existingResources, err := suite.repository.ListBySet(resourceSet)
	require.NoError(suite.T(), err, "Failed to list resources by set")
	assert.Equal(suite.T(), 2, len(*existingResources))

	for _, entity := range *existingResources {
		resourceStructFromDB := &objectresource.ResourceBaseStruct{}
		err = json.Unmarshal([]byte(entity.Document), resourceStructFromDB)
		require.NoError(suite.T(), err, "Failed unmarshal resource entity")

		if entity.Name == resourceName1 {
			assert.Equal(suite.T(), resourceStruct1, resourceStructFromDB)
		} else if entity.Name == resourceName2 {
			assert.Equal(suite.T(), resourceStruct2, resourceStructFromDB)
		}
	}

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(resourceEntity1)
			suite.repository.Delete(resourceEntity2)
		}
	})
}

func TestResourceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceRepositoryTestSuite))
}

func CreateNewSampleResource(name string) (*entities.Resource, *objectresource.ResourceBaseStruct, error) {

	resource := BasicResource_WithName(name)
	jsonData, err := json.Marshal(resource)
	if err != nil {
		return nil, nil, err
	}

	resourceEntity := entities.Resource{
		Name:     name,
		Document: string(jsonData),
	}

	return &resourceEntity, resource, nil
}

func CreateNewSampleResourceWithSet(name, set string) (*entities.Resource, *objectresource.ResourceBaseStruct, error) {

	resource := BasicResource_WithNameSet(name, set)

	jsonData, err := json.Marshal(resource)
	if err != nil {
		return nil, nil, err
	}

	resourceEntity := entities.Resource{
		Name:     name,
		Document: string(jsonData),
	}

	return &resourceEntity, resource, nil
}

func BasicResource_WithName(name string) *objectresource.ResourceBaseStruct {

	resource := objectresource.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			string(resource.ResourceKinds.Object),
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		objectresource.NewResourceSpecification(0,
			name,
			"",
			"/foo",
			"bar",
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]objectresource.Layer{objectresource.NewLayer(0, "layer1", []objectresource.Section{}), objectresource.NewLayer(0, "layer2", []objectresource.Section{})},
			[]objectresource.Attribute{
				objectresource.NewAttribute("yea", objectresource.VisibilityTypeTypes.Private,
					objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
				objectresource.NewAttribute("hoo", objectresource.VisibilityTypeTypes.Public,
					objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			},
			[]objectresource.Method{
				objectresource.NewMethod("soe", objectresource.VisibilityTypeTypes.Public,
					[]objectresource.MethodParameter{
						objectresource.NewMethodParameter("ID", objectresource.New_Int(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
						objectresource.NewMethodParameter("Name", objectresource.New_String(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
					},
					[]objectresource.DataType(nil),
					objectresource.Message{},
					objectresource.Message{},
					[]option.Option(nil),
					[]label.Label(nil),
					[]objectresource.Annotation(nil),
					"",
					true,
				),
			},
			applicationWorkspace.WorkspaceIdentifier{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	return &resource
}

func BasicResource_WithNameSet(name, set string) *objectresource.ResourceBaseStruct {

	resource := objectresource.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			string(resource.ResourceKinds.Object),
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		objectresource.NewResourceSpecification(0,
			name,
			"",
			"/foo",
			set,
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]objectresource.Layer{objectresource.NewLayer(0, "layer1", []objectresource.Section{}), objectresource.NewLayer(0, "layer2", []objectresource.Section{})},
			[]objectresource.Attribute{
				objectresource.NewAttribute("yea", objectresource.VisibilityTypeTypes.Private,
					objectresource.NewDataType(string(objectresource.ValueTypes.String), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
				objectresource.NewAttribute("hoo", objectresource.VisibilityTypeTypes.Public,
					objectresource.NewDataType(string(objectresource.ValueTypes.Int), objectresource.TypePackage{}, objectresource.DataTypeCategories.Value, objectresource.ModifierTypes.Object, []objectresource.DataType(nil)),
					0, objectresource.AttributeGroup{}, objectresource.Encapsulation{}, objectresource.AttributeProperties{}, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil), true),
			},
			[]objectresource.Method{
				objectresource.NewMethod("soe", objectresource.VisibilityTypeTypes.Public,
					[]objectresource.MethodParameter{
						objectresource.NewMethodParameter("ID", objectresource.New_Int(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
						objectresource.NewMethodParameter("Name", objectresource.New_String(), 0, objectresource.Message{}, objectresource.Message{}, []option.Option(nil), []label.Label(nil), objectresource.Validation{}, []objectresource.Annotation(nil)),
					},
					[]objectresource.DataType(nil),
					objectresource.Message{},
					objectresource.Message{},
					[]option.Option(nil),
					[]label.Label(nil),
					[]objectresource.Annotation(nil),
					"",
					true,
				),
			},
			applicationWorkspace.WorkspaceIdentifier{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	return &resource
}
