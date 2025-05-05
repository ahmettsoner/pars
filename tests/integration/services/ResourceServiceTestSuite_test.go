package services

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ResourceServiceTestSuite struct {
	suite.Suite
	service       services.ObjectResourceServiceInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ResourceServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = *services.NewObjectResourceService(suite.environment)

	suite.T().Log("Resource creation completed")
}
func (suite *ResourceServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *ResourceServiceTestSuite) SetupTest() {
}
func (suite *ResourceServiceTestSuite) TearDownTest() {
}

func (suite *ResourceServiceTestSuite) Test_CreateResource() {

	resourceName := suite.faker.Resource.Name()
	resource := *BasicResource_WithName(resourceName)

	temp, err := suite.service.Save(resource)
	require.NoError(suite.T(), err, "Failed to save resource")
	assert.Equal(suite.T(), resource, *temp)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(resource.Name, resource.Specifications.Workspace, true, true)
		}
	})
}

func (suite *ResourceServiceTestSuite) Test_GetByName() {

	resourceName := suite.faker.Resource.Name()
	resource := *BasicResource_WithName(resourceName)

	temp, err := suite.service.Save(resource)
	require.NoError(suite.T(), err, "Failed to save resource")
	assert.Equal(suite.T(), resource, *temp)

	existingResource, err := suite.service.GetByName(resourceName)
	require.NoError(suite.T(), err, "Failed to retrieve resource by name")

	assert.Equal(suite.T(), resource, *existingResource)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(resource.Name, resource.Specifications.Workspace, true, true)
		}
	})
}

func TestResourceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceServiceTestSuite))
}
func BasicResource(name string) *objectresource.ResourceBaseStruct {

	resource := objectresource.NewResourceBaseStruct(
		resource.NewHeader(
			structs.StructTypes.Resource,
			resource.StructKinds.Object,
			name,
			structs.Metadata{
				Tags: []string{},
			},
		),
		objectresource.NewResourceSpecification(0,
			name,
			"",
			"",
			"",
			[]string{},
			[]label.Label{},
			[]objectresource.Layer{},
			[]objectresource.Attribute{},
			[]objectresource.Method{},
			workspace.WorkspaceSpecification{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	return &resource
}

func BasicResource_WithName(name string) *objectresource.ResourceBaseStruct {

	resource := objectresource.NewResourceBaseStruct(
		resource.NewHeader(
			structs.StructTypes.Resource,
			resource.StructKinds.Object,
			name,
			structs.Metadata{
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
			[]objectresource.Layer{objectresource.NewLayer(0, "presentation:view", []objectresource.Section{}), objectresource.NewLayer(0, "layer2", []objectresource.Section{})},
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
			workspace.WorkspaceSpecification{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	return &resource
}

func BasicResource_WithNameSet(name, set string) *objectresource.ResourceBaseStruct {

	resource := objectresource.NewResourceBaseStruct(
		resource.NewHeader(
			structs.StructTypes.Resource,
			resource.StructKinds.Object,
			name,
			structs.Metadata{
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
			[]objectresource.Layer{objectresource.NewLayer(0, "presentation:view", []objectresource.Section{}), objectresource.NewLayer(0, "layer2", []objectresource.Section{})},
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
			workspace.WorkspaceSpecification{},
		),
		objectresource.NewResourceConfiguration(objectresource.ChangeTrackers.OnChange),
	)

	return &resource
}
