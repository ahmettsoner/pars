package shared_task

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/label"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/modules/template/shared_template_contract"
)

type SharedTemplateService struct {
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewSharedTemplateService(environment string) shared_template_contract.TemplateInterface {
	templateRespository := ioc.Get[*repositories.TemplateRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &SharedTemplateService{
		environment:                  environment,
		templateRespository:          templateRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s SharedTemplateService) GetByName(name string) (*shared_template_payload_structs.TemplateBaseStruct, error) {
	var template *shared_template_payload_structs.TemplateBaseStruct

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

func (s SharedTemplateService) ListBySetAndLayers(set string, layers ...string) (*([]shared_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template shared_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s SharedTemplateService) ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]shared_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByFilter(set, workspace, layers, tags, label.ConvertLabelsToMap(labels))
	if err != nil {
		return nil, err
	}

	templateList := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template shared_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s SharedTemplateService) Save(model shared_template_payload_structs.TemplateBaseStruct) (*shared_template_payload_structs.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s SharedTemplateService) List() (*([]shared_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(shared_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template shared_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s SharedTemplateService) ListByWorkspace(workspace string) (*([]shared_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, shared_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]shared_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template shared_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s SharedTemplateService) Remove(name, workspace string, permanent bool) (*shared_template_payload_structs.TemplateBaseStruct, error) {
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
	var template shared_template_payload_structs.TemplateBaseStruct
	err = json.Unmarshal([]byte(templateTemplateEntity.Document), &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (s SharedTemplateService) IsExists(name, workspace string) (bool, error) {

	templateTemplateEntity, err := s.templateRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false, fmt.Errorf("xxx: Shared Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if templateTemplateEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s SharedTemplateService) GetHash(name string) (string, error) {

	entity, err := s.templateRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Shared Template getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Shared Template tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s SharedTemplateService) saveTemplateInformation(templateModel shared_template_payload_structs.TemplateBaseStruct) (*shared_template_payload_structs.TemplateBaseStruct, error) {

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
