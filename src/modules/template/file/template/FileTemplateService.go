package file_template

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/label"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"

	"parsdevkit.net/modules/template/file_template_contract"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type FileTemplateService struct {
	templateRespository          *repositories.TemplateRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewFileTemplateService(environment string) file_template_contract.TemplateInterface {
	templateRespository := ioc.Get[*repositories.TemplateRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &FileTemplateService{
		environment:                  environment,
		templateRespository:          templateRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s FileTemplateService) GetByName(name string) (*file_template_payload_structs.TemplateBaseStruct, error) {
	var template *file_template_payload_structs.TemplateBaseStruct

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

func (s FileTemplateService) Save(model file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

	result, err := s.saveTemplateInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s FileTemplateService) List() (*([]file_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByKind(file_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template file_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s FileTemplateService) ListByWorkspace(workspace string) (*([]file_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceAndKind(workspace, file_template_payload_structs.TEMPLATE_KIND)
	if err != nil {
		return nil, err
	}

	templateList := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template file_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) ListBySetAndLayers(set string, layers ...string) (*([]file_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template file_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]file_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template file_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s FileTemplateService) ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]file_template_payload_structs.TemplateBaseStruct), error) {

	entityList, err := s.templateRespository.ListByFilter(set, workspace, layers, tags, label.ConvertLabelsToMap(labels))
	if err != nil {
		return nil, err
	}

	templateList := make([]file_template_payload_structs.TemplateBaseStruct, 0)

	for _, entity := range *entityList {
		var template file_template_payload_structs.TemplateBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s FileTemplateService) Remove(name, workspace string, permanent bool) (*file_template_payload_structs.TemplateBaseStruct, error) {
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
	var template file_template_payload_structs.TemplateBaseStruct
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

func (s FileTemplateService) saveTemplateInformation(templateModel file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

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

func (s FileTemplateService) SaveTemplate(model file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

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

func (s *FileTemplateService) UndoSaveTemplate(model file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: File Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *FileTemplateService) DeleteTemplate(model file_template_payload_structs.TemplateBaseStruct) (*file_template_payload_structs.TemplateBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: File Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}

func (s *FileTemplateService) ClearTemplateHistory(model file_template_payload_structs.TemplateBaseStruct) error {

	err := s.generationHistoryRespository.DeleteBySetAndTemplate(model.Specifications.Set, model.Header.Name)
	if err != nil {
		return err
	}

	return nil
}
