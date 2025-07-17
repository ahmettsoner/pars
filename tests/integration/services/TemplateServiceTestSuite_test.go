package services

import (
	"testing"

	"parsdevkit.net/application"

	applicationTemplate "parsdevkit.net/application/structs/template"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	templateStruct "parsdevkit.net/application/structs/template"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	templateCode "parsdevkit.net/modules/template/code_template"
	"parsdevkit.net/modules/template/code_template_contract"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"parsdevkit.net/application/schemas"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TemplateServiceTestSuite struct {
	suite.Suite
	service       code_template_contract.TemplateInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *TemplateServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = templateCode.NewCodeTemplateService(suite.environment)

	suite.T().Log("Template creation completed")
}
func (suite *TemplateServiceTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *TemplateServiceTestSuite) SetupTest() {
}
func (suite *TemplateServiceTestSuite) TearDownTest() {
}

func (suite *TemplateServiceTestSuite) Test_CreateTemplate() {

	templateName := suite.faker.Resource.Name()
	template := *BasicTemplate_WithName(templateName)

	temp, err := suite.service.Save(template)
	require.NoError(suite.T(), err, "Failed to save template")
	assert.Equal(suite.T(), template, *temp)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(template.Header.Name, template.Specifications.Workspace, true)
		}
	})
}

func (suite *TemplateServiceTestSuite) Test_GetByName() {

	templateName := suite.faker.Resource.Name()
	template := *BasicTemplate_WithName(templateName)

	temp, err := suite.service.Save(template)
	require.NoError(suite.T(), err, "Failed to save template")
	assert.Equal(suite.T(), template, *temp)

	existingTemplate, err := suite.service.GetByName(templateName)
	require.NoError(suite.T(), err, "Failed to retrieve template by name")

	assert.Equal(suite.T(), template, *existingTemplate)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.service.Remove(template.Header.Name, template.Specifications.Workspace, true)
		}
	})
}

func TestTemplateServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateServiceTestSuite))
}

func BasicTemplate_WithName(name string) *code_template_payload_structs.TemplateBaseStruct {

	template := code_template_payload_structs.NewTemplateBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Template, code_template_payload_structs.TEMPLATE_KIND, name, schemas.Metadata{}),
		applicationTemplate.NewTemplateSpecification(
			0,
			name,
			"",
			"bar",
			"sample_template",
			applicationTemplate.NewOutput("sample.cs"),
			[]string{"pack", "age"},
			[]label.Label{label.NewLabel("foo", "bar")},
			[]layerPkg.Layer{
				layerPkg.NewLayer(0, "service:contract", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "service", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "presentation:view", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "presentation:viewmodel", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:repository", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:entity", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:migration", []sectionPkg.Section(nil)),
			},
			applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.Code, "test-code-content"),
			applicationWorkspace.WorkspaceIdentifier{},
		),
		code_template_payload_structs.NewTemplateConfiguration(code_template_payload_structs.ChangeTrackers.OnChange, templateStruct.Selectors{}),
	)

	return &template
}

func BasicTemplate_WithNameSet(name, set string) *code_template_payload_structs.TemplateBaseStruct {

	template := code_template_payload_structs.NewTemplateBaseStruct(
		schemas.NewSchemaHeader(schemas.StructTypes.Template, code_template_payload_structs.TEMPLATE_KIND, name, schemas.Metadata{}),
		applicationTemplate.NewTemplateSpecification(
			0,
			name,
			"",
			set,
			"sample_template",
			applicationTemplate.NewOutput("sample.cs"),
			[]string{"pack", "age"},
			[]label.Label{label.NewLabel("foo", "bar")},
			[]layerPkg.Layer{
				layerPkg.NewLayer(0, "service:contract", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "service", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "presentation:view", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "presentation:viewmodel", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:repository", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:entity", []sectionPkg.Section(nil)),
				layerPkg.NewLayer(0, "persistence:database:migration", []sectionPkg.Section(nil)),
			},
			applicationTemplate.NewTemplate(applicationTemplate.TemplateSourceTypes.Code, "test-code-content"),
			applicationWorkspace.WorkspaceIdentifier{},
		),
		code_template_payload_structs.NewTemplateConfiguration(code_template_payload_structs.ChangeTrackers.OnChange, templateStruct.Selectors{}),
	)

	return &template
}
