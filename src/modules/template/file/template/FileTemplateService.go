package file_template

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/structs/template"
	filetemplate "parsdevkit.net/structs/template/file-template"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type FileTemplateService struct {
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewFileTemplateService(environment string) contracts.TemplateServiceInterface[filetemplate.TemplateBaseStruct] {
	templateRespository := ioc.Get[*repositories.TemplateRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &FileTemplateService{
		environment:                  environment,
		templateRespository:          templateRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s FileTemplateService) GetByName(name string) (*filetemplate.TemplateBaseStruct, error) {
	var template *filetemplate.TemplateBaseStruct

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

func (s FileTemplateService) Save(model filetemplate.TemplateBaseStruct) (*filetemplate.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s FileTemplateService) List() (*([]filetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(string(template.TemplateKinds.File))
	if err != nil {
		return nil, err
	}

	templateList := make([]filetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template filetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s FileTemplateService) ListByWorkspace(workspace string) (*([]filetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, string(template.TemplateKinds.File))
	if err != nil {
		return nil, err
	}

	templateList := make([]filetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template filetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) ListBySetAndLayers(set string, layers ...string) (*([]filetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]filetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template filetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]filetemplate.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]filetemplate.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template filetemplate.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) Remove(name, workspace string, permanent bool) (*filetemplate.TemplateBaseStruct, error) {
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
	var template filetemplate.TemplateBaseStruct
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

func (s FileTemplateService) IsExists(name, workspace string) (bool, error) {

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false, fmt.Errorf("xxx: File Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if templateTemplateEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s FileTemplateService) GetHash(name string) (string, error) {

	entity, err := s.templateRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: File Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: File Template tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s FileTemplateService) saveTemplateInformation(templateModel filetemplate.TemplateBaseStruct) (*filetemplate.TemplateBaseStruct, error) {

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
