package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/application"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	"parsdevkit.net/application/structs"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/application/schemas"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	applicationResource "parsdevkit.net/application/structs/resource"
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
	testArea := application.GenerateTestArea()
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

	resourceStructFromDB := &object_resource_payload_structs.ResourceBaseStruct{}
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
		resourceStructFromDB := &object_resource_payload_structs.ResourceBaseStruct{}
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

func CreateNewSampleResource(name string) (*entities.Resource, *object_resource_payload_structs.ResourceBaseStruct, error) {

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

func CreateNewSampleResourceWithSet(name, set string) (*entities.Resource, *object_resource_payload_structs.ResourceBaseStruct, error) {

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

func BasicResource_WithName(name string) *object_resource_payload_structs.ResourceBaseStruct {

	resource := object_resource_payload_structs.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			object_resource_payload_structs.RESOURCE_KIND,
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		applicationResource.NewResourceSpecification(0,
			name,
			"",
			"/foo",
			"bar",
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]layerPkg.Layer{layerPkg.NewLayer(0, "layer1", []sectionPkg.Section{}), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section{})},

			applicationWorkspace.WorkspaceIdentifier{},
		),
		object_resource_payload_structs.NewResourceObject(
			[]object_resource_payload_structs.Attribute{
				object_resource_payload_structs.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
					structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
				object_resource_payload_structs.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
					structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
			},
			[]object_resource_payload_structs.Method{
				object_resource_payload_structs.NewMethod("soe", structs.VisibilityTypeTypes.Public,
					[]object_resource_payload_structs.MethodParameter{
						object_resource_payload_structs.NewMethodParameter("ID", structs.New_Int(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
						object_resource_payload_structs.NewMethodParameter("Name", structs.New_String(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
					},
					[]structs.DataType(nil),
					object_resource_payload_structs.Message{},
					object_resource_payload_structs.Message{},
					[]option.Option(nil),
					[]label.Label(nil),
					[]object_resource_payload_structs.Annotation(nil),
					"",
					true,
				),
			},
		),
		object_resource_payload_structs.NewResourceConfiguration(object_resource_payload_structs.ChangeTrackers.OnChange),
	)

	return &resource
}

func BasicResource_WithNameSet(name, set string) *object_resource_payload_structs.ResourceBaseStruct {

	resource := object_resource_payload_structs.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			object_resource_payload_structs.RESOURCE_KIND,
			name,
			schemas.Metadata{
				Tags: []string{"tag1", "tag2"},
			},
		),
		applicationResource.NewResourceSpecification(0,
			name,
			"",
			"/foo",
			set,
			[]string{"pars", "cmd"},
			[]label.Label{
				label.NewLabel("foo", "bar"),
			},
			[]layerPkg.Layer{layerPkg.NewLayer(0, "layer1", []sectionPkg.Section{}), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section{})},
			applicationWorkspace.WorkspaceIdentifier{},
		),
		object_resource_payload_structs.NewResourceObject(
			[]object_resource_payload_structs.Attribute{
				object_resource_payload_structs.NewAttribute("yea", structs.VisibilityTypeTypes.Private,
					structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
				object_resource_payload_structs.NewAttribute("hoo", structs.VisibilityTypeTypes.Public,
					structs.NewDataType(string(structs.ValueTypes.Int), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil)),
					0, object_resource_payload_structs.AttributeGroup{}, object_resource_payload_structs.Encapsulation{}, object_resource_payload_structs.AttributeProperties{}, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil), true),
			},
			[]object_resource_payload_structs.Method{
				object_resource_payload_structs.NewMethod("soe", structs.VisibilityTypeTypes.Public,
					[]object_resource_payload_structs.MethodParameter{
						object_resource_payload_structs.NewMethodParameter("ID", structs.New_Int(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
						object_resource_payload_structs.NewMethodParameter("Name", structs.New_String(), 0, object_resource_payload_structs.Message{}, object_resource_payload_structs.Message{}, []option.Option(nil), []label.Label(nil), object_resource_payload_structs.Validation{}, []object_resource_payload_structs.Annotation(nil)),
					},
					[]structs.DataType(nil),
					object_resource_payload_structs.Message{},
					object_resource_payload_structs.Message{},
					[]option.Option(nil),
					[]label.Label(nil),
					[]object_resource_payload_structs.Annotation(nil),
					"",
					true,
				),
			},
		),
		object_resource_payload_structs.NewResourceConfiguration(object_resource_payload_structs.ChangeTrackers.OnChange),
	)

	return &resource
}
