package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/structs/template"
	codetemplate "parsdevkit.net/structs/template/code-template"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type CodeTemplateService struct {
	CodeTemplateServiceInterface
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewCodeTemplateService(environment string) *CodeTemplateService {
	dbContext := contexts.NewDbContext(environment)
	return &CodeTemplateService{
		environment:                  environment,
		templateRespository:          repositories.NewTemplateRepository(dbContext),
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(dbContext),
	}
}

func (s CodeTemplateService) GetByName(name string) (*codetemplate.TemplateBaseStruct, error) {
	var template *codetemplate.TemplateBaseStruct

	entity, err := s.templateRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &template)
	} else {
		template = nil
	}

	return template, nil
}

func (s CodeTemplateService) Save(model codetemplate.TemplateBaseStruct) (*codetemplate.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CodeTemplateService) List() (*([]codetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(string(template.TemplateKinds.Code))
	if err != nil {
		return nil, err
	}

	templateList := make([]codetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template codetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListByWorkspace(workspace string) (*([]codetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, string(template.TemplateKinds.Code))
	if err != nil {
		return nil, err
	}

	templateList := make([]codetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template codetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListBySetAndLayers(set string, layers ...string) (*([]codetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]codetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template codetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]codetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]codetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template codetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) Remove(name, workspace string, permanent bool) (*codetemplate.TemplateBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return nil, err
	}
	if templateTemplateEntity == nil {
		return nil, errors.New("invalid template name")
	}

	logrus.Debugf("template %v deleting...", templateTemplateEntity.Name)

	err = s.templateRespository.Delete(templateTemplateEntity)
	var template codetemplate.TemplateBaseStruct
	err = json.Unmarshal([]byte(templateTemplateEntity.Document), &template)
	if err != nil {
		return nil, err
	}

	err = s.generationHistoryRespository.DeleteBySetAndTemplate(template.Specifications.Set, template.Header.Name)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (s CodeTemplateService) IsExists(name, workspace string) (bool, error) {

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false, fmt.Errorf("xxx: Code Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if templateTemplateEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s CodeTemplateService) GetHash(name string) (string, error) {

	entity, err := s.templateRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Template tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s CodeTemplateService) saveTemplateInformation(templateModel codetemplate.TemplateBaseStruct) (*codetemplate.TemplateBaseStruct, error) {

	jsonData, err := json.Marshal(templateModel)
	if err != nil {
		return nil, err
	}

	templateEntity := entities.Template{
		Name:     templateModel.Header.Name,
		Document: string(jsonData),
	}

	err = s.templateRespository.Save(&templateEntity)
	if err != nil {
		return nil, err
	}

	return &templateModel, nil
}
