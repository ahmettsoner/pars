package services

import (
	"testing"

	"parsdevkit.net/application"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	applicationResource "parsdevkit.net/application/structs/resource"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/application/schemas"
	resourceObject "parsdevkit.net/modules/resource/object_resource"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"parsdevkit.net/modules/resource/object_resource_contract"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ResourceServiceTestSuite struct {
	suite.Suite
	service       object_resource_contract.ResourceInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *ResourceServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = resourceObject.NewObjectResourceService(suite.environment)

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
			suite.service.Remove(resource.Header.Name, resource.Specifications.Workspace, true, true)
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
			suite.service.Remove(resource.Header.Name, resource.Specifications.Workspace, true, true)
		}
	})
}

func TestResourceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceServiceTestSuite))
}
func BasicResource(name string) *object_resource_payload_structs.ResourceBaseStruct {

	resource := object_resource_payload_structs.NewResourceBaseStruct(
		schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			object_resource_payload_structs.RESOURCE_KIND,
			name,
			schemas.Metadata{
				Tags: []string{},
			},
		),
		applicationResource.NewResourceSpecification(0,
			name,
			"",
			"",
			"",
			[]string{},
			[]label.Label{},
			[]layerPkg.Layer{},
			applicationWorkspace.WorkspaceIdentifier{},
			structs.ChangeTrackers.OnChange,
		),
		object_resource_payload_structs.NewResourceObject(
			[]object_resource_payload_structs.Attribute{},
			[]object_resource_payload_structs.Method{},
		),
		object_resource_payload_structs.NewResourceConfiguration(),
	)

	return &resource
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
			[]layerPkg.Layer{layerPkg.NewLayer(0, "presentation:view", []sectionPkg.Section{}), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section{})},

			applicationWorkspace.WorkspaceIdentifier{},
			structs.ChangeTrackers.OnChange,
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
		object_resource_payload_structs.NewResourceConfiguration(),
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
			[]layerPkg.Layer{layerPkg.NewLayer(0, "presentation:view", []sectionPkg.Section{}), layerPkg.NewLayer(0, "layer2", []sectionPkg.Section{})},
			applicationWorkspace.WorkspaceIdentifier{},
			structs.ChangeTrackers.OnChange,
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
		object_resource_payload_structs.NewResourceConfiguration(),
	)

	return &resource
}
