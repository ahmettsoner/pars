package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/application"
	"parsdevkit.net/application/models/label"
	layerPkg "parsdevkit.net/application/models/layer"
	sectionPkg "parsdevkit.net/application/models/section"
	applicationTemplate "parsdevkit.net/application/structs/template"
	templateStruct "parsdevkit.net/application/structs/template"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"parsdevkit.net/application/schemas"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/entities"

	"pars/tests/internal/testenv/common"
	"pars/tests/internal/testenv/faker"

	"parsdevkit.net/persistence/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TemplateRepositoryTestSuite struct {
	suite.Suite
	environment   string
	repository    repositories.TemplateRepository
	faker         *faker.Faker
	noCleanOnFail bool
}

func (suite *TemplateRepositoryTestSuite) SetupSuite() {

	suite.T().Log("Preparing test suite...")

	suite.faker = faker.NewFaker()
	suite.noCleanOnFail = true
	testArea := application.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	dbContext := contexts.NewDbContext(suite.environment)
	suite.repository = *repositories.NewTemplateRepository(dbContext)

	suite.T().Log("Template creation completed")
}
func (suite *TemplateRepositoryTestSuite) TearDownSuite() {
	suite.T().Log("Test suite disposing...")
	if !suite.noCleanOnFail || !suite.T().Failed() {
	}
}

func (suite *TemplateRepositoryTestSuite) SetupTest() {
}
func (suite *TemplateRepositoryTestSuite) TearDownTest() {
}

func (suite *TemplateRepositoryTestSuite) Test_CreateTemplate() {

	templateName := suite.faker.Project.Name()
	templateEntity, _, err := CreateNewSampleTemplate(templateName)
	require.NoError(suite.T(), err, "Template creation failed")

	err = suite.repository.Save(templateEntity)
	require.NoError(suite.T(), err, "Failed to save template")

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(templateEntity)
		}
	})
}

func (suite *TemplateRepositoryTestSuite) Test_GetByName() {

	templateName := suite.faker.Project.Name()
	templateEntity, templateStruct, err := CreateNewSampleTemplate(templateName)
	require.NoError(suite.T(), err, "Template creation failed")

	err = suite.repository.Save(templateEntity)
	require.NoError(suite.T(), err, "Failed to save template")

	existingTemplate, err := suite.repository.GetByName(templateName)
	require.NoError(suite.T(), err, "Failed to retrieve template by name")

	templateStructFromDB := &code_template_payload_structs.TemplateBaseStruct{}
	err = json.Unmarshal([]byte(existingTemplate.Document), templateStructFromDB)
	require.NoError(suite.T(), err, "Failed unmarshal template entity")

	assert.Equal(suite.T(), templateStruct, templateStructFromDB)

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(templateEntity)
		}
	})
}

func (suite *TemplateRepositoryTestSuite) Test_ListBySetAndLayers() {

	templateName1 := suite.faker.Project.Name()
	templateName2 := suite.faker.Project.Name()
	setName := suite.faker.Project.Set()
	templateEntity1, templateStruct1, err := CreateNewSampleTemplateWithSet(templateName1, setName)
	require.NoError(suite.T(), err, "Template creation failed")

	suite.repository.Save(templateEntity1)
	require.NoError(suite.T(), err, "Failed to save template")

	templateEntity2, templateStruct2, err := CreateNewSampleTemplateWithSet(templateName2, setName)
	require.NoError(suite.T(), err, "Template creation failed")

	suite.repository.Save(templateEntity2)
	require.NoError(suite.T(), err, "Failed to save template")

	existingTemplates, err := suite.repository.ListBySetAndLayers(setName, "service:contract")
	require.NoError(suite.T(), err, "Failed to list templates by set and layers")
	assert.Equal(suite.T(), 2, len(*existingTemplates))

	for _, entity := range *existingTemplates {
		templateStructFromDB := &code_template_payload_structs.TemplateBaseStruct{}
		err = json.Unmarshal([]byte(entity.Document), templateStructFromDB)
		require.NoError(suite.T(), err, "Failed unmarshal template entity")

		if entity.Name == templateName1 {
			assert.Equal(suite.T(), templateStruct1, templateStructFromDB)
		} else if entity.Name == templateName2 {
			assert.Equal(suite.T(), templateStruct2, templateStructFromDB)
		}
	}

	suite.T().Cleanup(func() {
		if !suite.noCleanOnFail || !suite.T().Failed() {
			suite.repository.Delete(templateEntity1)
			suite.repository.Delete(templateEntity2)
		}
	})
}

func TestTemplateRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateRepositoryTestSuite))
}

func CreateNewSampleTemplate(name string) (*entities.Template, *code_template_payload_structs.TemplateBaseStruct, error) {

	template := BasicTemplate_WithName(name)

	jsonData, err := json.Marshal(template)
	if err != nil {
		return nil, nil, err
	}

	templateEntity := entities.Template{
		Name:     name,
		Document: string(jsonData),
	}

	return &templateEntity, template, nil
}

func CreateNewSampleTemplateWithSet(name, set string) (*entities.Template, *code_template_payload_structs.TemplateBaseStruct, error) {

	template := BasicTemplate_WithNameSet(name, set)

	jsonData, err := json.Marshal(template)
	if err != nil {
		return nil, nil, err
	}

	templateEntity := entities.Template{
		Name:     name,
		Document: string(jsonData),
	}

	return &templateEntity, template, nil
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
