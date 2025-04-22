package services

import (
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/template"
	templateStruct "parsdevkit.net/structs/template"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TemplateServiceTestSuite struct {
	suite.Suite
	service       services.CodeTemplateServiceInterface
	environment   string
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *TemplateServiceTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.service = *services.NewCodeTemplateService(suite.environment)

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
			suite.service.Remove(template.Name, template.Specifications.Workspace, true)
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
			suite.service.Remove(template.Name, template.Specifications.Workspace, true)
		}
	})
}

func TestTemplateServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateServiceTestSuite))
}

func BasicTemplate_WithName(name string) *codetemplate.TemplateBaseStruct {

	template := codetemplate.NewTemplateBaseStruct(
		templateStruct.NewHeader(structs.StructTypes.Template, templateStruct.StructKinds.Code, name, structs.Metadata{}),
		codetemplate.NewTemplateSpecification(
			0,
			name,
			"",
			"bar",
			"sample_template",
			codetemplate.NewOutput("sample.cs"),
			[]string{"pack", "age"},
			[]label.Label{label.NewLabel("foo", "bar")},
			[]codetemplate.Layer{
				codetemplate.NewLayer(0, "service:contract", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "service", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "presentation:view", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "presentation:viewmodel", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:repository", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:entity", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:migration", []templateStruct.Section(nil)),
			},
			codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.Code, "test-code-content"),
			workspace.WorkspaceSpecification{},
		),
		codetemplate.NewTemplateConfiguration(codetemplate.ChangeTrackers.OnChange, template.Selectors{}),
	)

	return &template
}

func BasicTemplate_WithNameSet(name, set string) *codetemplate.TemplateBaseStruct {

	template := codetemplate.NewTemplateBaseStruct(
		templateStruct.NewHeader(structs.StructTypes.Template, templateStruct.StructKinds.Code, name, structs.Metadata{}),
		codetemplate.NewTemplateSpecification(
			0,
			name,
			"",
			set,
			"sample_template",
			codetemplate.NewOutput("sample.cs"),
			[]string{"pack", "age"},
			[]label.Label{label.NewLabel("foo", "bar")},
			[]codetemplate.Layer{
				codetemplate.NewLayer(0, "service:contract", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "service", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "presentation:view", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "presentation:viewmodel", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:repository", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:entity", []templateStruct.Section(nil)),
				codetemplate.NewLayer(0, "persistence:database:migration", []templateStruct.Section(nil)),
			},
			codetemplate.NewTemplate(codetemplate.TemplateSourceTypes.Code, "test-code-content"),
			workspace.WorkspaceSpecification{},
		),
		codetemplate.NewTemplateConfiguration(codetemplate.ChangeTrackers.OnChange, template.Selectors{}),
	)

	return &template
}
