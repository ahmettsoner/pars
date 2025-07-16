package shared_task

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/models"
	"parsdevkit.net/modules/resource/object_resource"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
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

func (s SharedTemplateService) SaveTemplate(model shared_template_payload_structs.TemplateBaseStruct) (*shared_template_payload_structs.TemplateBaseStruct, error) {

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

func (s *SharedTemplateService) UndoSaveTemplate(model shared_template_payload_structs.TemplateBaseStruct) (*shared_template_payload_structs.TemplateBaseStruct, error) {

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Shared Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *SharedTemplateService) DeleteTemplate(model shared_template_payload_structs.TemplateBaseStruct) (*shared_template_payload_structs.TemplateBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.templateRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Shared Template silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}

func (s *SharedTemplateService) ClearTemplateHistory(model shared_template_payload_structs.TemplateBaseStruct) error {

	return nil
}

func (s *SharedTemplateService) Context(workspace, project, resource, template schemas.SchemaInterface, layer layer.LayerIdentifier, section section.SectionIdentifier) interface{} {

	if model, ok := template.(shared_template_payload_structs.TemplateBaseStruct); ok {
		if modelProject, ok := project.(application_project_payload_structs.ProjectBaseStruct); ok {
			return TemplateComposite{
				SharedTemplate: s.structToModel(modelProject.Specifications.Platform.Type, model),
				Original:       model,
			}
		}
	}

	return nil
}
func (s *SharedTemplateService) structToModel(platform models.PlatformType, model shared_template_payload_structs.TemplateBaseStruct) SharedTemplate {

	var result SharedTemplate = SharedTemplate{
		Name: model.Header.Name,
	}
	return result
}
func (s *SharedTemplateService) LabelListToModel(labels ...label.Label) []object_resource.ObjectLabel {

	var result []object_resource.ObjectLabel = make([]object_resource.ObjectLabel, 0)

	for _, label := range labels {
		result = append(result, object_resource.ObjectLabel{
			Key:   label.Key,
			Value: label.Value,
		})
	}

	return result
}

type TemplateComposite struct {
	SharedTemplate
	Original shared_template_payload_structs.TemplateBaseStruct
}

type SharedTemplate struct {
	Name string
}
