package net8

import (
	"encoding/json"
	"testing"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/template"
	templateStruct "parsdevkit.net/structs/template"
	codetemplate "parsdevkit.net/structs/template/code-template"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/entities"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/core/test/common"
	"parsdevkit.net/core/test/faker"

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
	testArea := utils.GenerateTestArea()
	suite.environment = common.GenerateEnvironment(suite.T(), testArea)
	suite.repository = *repositories.NewTemplateRepository(suite.environment)

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

	templateStructFromDB := &codetemplate.TemplateBaseStruct{}
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
		templateStructFromDB := &codetemplate.TemplateBaseStruct{}
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

func CreateNewSampleTemplate(name string) (*entities.Template, *codetemplate.TemplateBaseStruct, error) {

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

func CreateNewSampleTemplateWithSet(name, set string) (*entities.Template, *codetemplate.TemplateBaseStruct, error) {

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
