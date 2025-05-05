package services

import (
	"encoding/json"
	"errors"

	"parsdevkit.net/structs/template"
	sharedtemplate "parsdevkit.net/structs/template/shared-template"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type SharedTemplateService struct {
	SharedTemplateServiceInterface
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewSharedTemplateService(environment string) *SharedTemplateService {
	return &SharedTemplateService{
		environment:                  environment,
		templateRespository:          repositories.NewTemplateRepository(environment),
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s SharedTemplateService) GetByName(name string) (*sharedtemplate.TemplateBaseStruct, error) {
	var template *sharedtemplate.TemplateBaseStruct

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

func (s SharedTemplateService) Save(model sharedtemplate.TemplateBaseStruct) (*sharedtemplate.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s SharedTemplateService) List() (*([]sharedtemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(string(template.StructKinds.Shared))
	if err != nil {
		return nil, err
	}

	templateList := make([]sharedtemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template sharedtemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s SharedTemplateService) ListByWorkspace(workspace string) (*([]sharedtemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, string(template.StructKinds.Shared))
	if err != nil {
		return nil, err
	}

	templateList := make([]sharedtemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template sharedtemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s SharedTemplateService) Remove(name, workspace string, permanent bool) (*sharedtemplate.TemplateBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return nil, err
	}
	if templateTemplateEntity == nil {
		return nil, errors.New("invalid template template")
	}

	logrus.Debugf("template %v deleting...", templateTemplateEntity.Name)

	err = s.templateRespository.Delete(templateTemplateEntity)
	var template sharedtemplate.TemplateBaseStruct
	err = json.Unmarshal([]byte(templateTemplateEntity.Document), &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (s SharedTemplateService) IsExists(name, workspace string) bool {

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false
	}
	if templateTemplateEntity == nil {
		return false
	}

	return true
}
func (s SharedTemplateService) GetHash(name string) string {

	entity, err := s.templateRespository.GetByName(name)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s SharedTemplateService) saveTemplateInformation(templateModel sharedtemplate.TemplateBaseStruct) (*sharedtemplate.TemplateBaseStruct, error) {

	jsonData, err := json.Marshal(templateModel)
	if err != nil {
		return nil, err
	}

	templateEntity := entities.Template{
		Name:     templateModel.Name,
		Document: string(jsonData),
	}

	err = s.templateRespository.Save(&templateEntity)
	if err != nil {
		return nil, err
	}

	return &templateModel, nil
}
