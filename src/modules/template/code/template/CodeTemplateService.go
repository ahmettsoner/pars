package code_template

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type CodeTemplateService struct {
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewCodeTemplateService(environment string) code_template_contract.TemplateInterface {
	templateRespository := ioc.Get[*repositories.TemplateRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &CodeTemplateService{
		environment:                  environment,
		templateRespository:          templateRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s CodeTemplateService) GetByName(name string) (*code_template_payload_structs.TemplateBaseStruct, error) {
	var template *code_template_payload_structs.TemplateBaseStruct

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

func (s CodeTemplateService) Save(model code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CodeTemplateService) List() (*([]code_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(code_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]code_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template code_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListByWorkspace(workspace string) (*([]code_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, code_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]code_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template code_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListBySetAndLayers(set string, layers ...string) (*([]code_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]code_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template code_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]code_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]code_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template code_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s CodeTemplateService) ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]code_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByFilter(set, workspace, layers, tags, label.ConvertLabelsToMap(labels))
	if err != nil {
		return nil, err
	}

	templateList := make([]code_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template code_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s CodeTemplateService) Remove(name, workspace string, permanent bool) (*code_template_payload_structs.TemplateBaseStruct, error) {
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
	var template code_template_payload_structs.TemplateBaseStruct
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

func (s CodeTemplateService) saveTemplateInformation(templateModel code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

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

func (s CodeTemplateService) SaveTemplate(model code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	resourceEntity := entities.Template{
		Name:     model.Header.Name,
		Document: string(jsonData),
	}

	err = s.templateRespository.Save(&resourceEntity)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *CodeTemplateService) UndoSaveTemplate(model code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *CodeTemplateService) DeleteTemplate(model code_template_payload_structs.TemplateBaseStruct) (*code_template_payload_structs.TemplateBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}
