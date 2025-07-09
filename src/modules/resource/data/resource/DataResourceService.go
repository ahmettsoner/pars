package data_resource

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/modules/resource/data_resource_contract"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type DataResourceService struct {
	resourceRepository           *repositories.ResourceRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewDataResourceService(environment string) data_resource_contract.ResourceInterface {
	resourceRepository := ioc.Get[*repositories.ResourceRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &DataResourceService{
		environment:                  environment,
		resourceRepository:           resourceRepository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s DataResourceService) GetByName(name string) (*data_resource_payload_structs.ResourceBaseStruct, error) {
	var resource *data_resource_payload_structs.ResourceBaseStruct

	entity, err := s.resourceRepository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &resource)
	} else {
		resource = nil
	}

	return resource, nil
}

func (s DataResourceService) Save(model data_resource_payload_structs.ResourceBaseStruct) (*data_resource_payload_structs.ResourceBaseStruct, error) {

	result, err := s.SaveResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s DataResourceService) List() (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByKind(data_resource_payload_structs.RESOURCE_KIND)
	if err != nil {
		return nil, err
	}

	resourceList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}
func (s DataResourceService) ListByWorkspace(workspace string) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndKind(workspace, data_resource_payload_structs.RESOURCE_KIND)
	if err != nil {
		return nil, err
	}

	resourceList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s DataResourceService) ListBySet(set string) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySet(set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}
func (s DataResourceService) ListByWorkspaceAndSet(workspace, set string) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndSet(workspace, set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s DataResourceService) ListBySetAndLayers(set string, layers ...string) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s DataResourceService) ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s DataResourceService) ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]data_resource_payload_structs.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByFilter(set, workspace, layers, tags, label.ConvertLabelsToMap(labels))
	if err != nil {
		return nil, err
	}

	templateList := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template data_resource_payload_structs.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s DataResourceService) Remove(name, workspace string, force, permanent bool) (*data_resource_payload_structs.ResourceBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	resourceResourceEntity, err := s.resourceRepository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return nil, err
	}
	if resourceResourceEntity == nil {
		return nil, errors.New("invalid resource resource")
	}

	logrus.Debugf("resource %v deleting...", resourceResourceEntity.Name)

	err = s.resourceRepository.Delete(resourceResourceEntity)
	var resource data_resource_payload_structs.ResourceBaseStruct
	err = json.Unmarshal([]byte(resourceResourceEntity.Document), &resource)
	if err != nil {
		return nil, err
	}

	err = s.generationHistoryRespository.DeleteBySetAndResource(resource.Specifications.Set, resource.Header.Name)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}
func (s *DataResourceService) DeleteResource(model data_resource_payload_structs.ResourceBaseStruct) (*data_resource_payload_structs.ResourceBaseStruct, error) {

	logrus.Debugf("resource %v removing", model.Header.Name)

	err := s.resourceRepository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Data Resource silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("resource (%v) information removed", model)

	return &model, nil
}
func (s DataResourceService) IsExists(name, workspace string) (bool, error) {

	resourceResourceEntity, err := s.resourceRepository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false, fmt.Errorf("xxx: Data Resource getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if resourceResourceEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s DataResourceService) GetHash(name string) (string, error) {

	entity, err := s.resourceRepository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Data Resource getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Data Resource tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s DataResourceService) SaveResource(model data_resource_payload_structs.ResourceBaseStruct) (*data_resource_payload_structs.ResourceBaseStruct, error) {

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	resourceEntity := entities.Resource{
		Name:     model.Header.Name,
		Document: string(jsonData),
	}

	err = s.resourceRepository.Save(&resourceEntity)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *DataResourceService) UndoSaveResource(model data_resource_payload_structs.ResourceBaseStruct) (*data_resource_payload_structs.ResourceBaseStruct, error) {

	err := s.resourceRepository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Application Project silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *DataResourceService) ClearResourceHistory(model data_resource_payload_structs.ResourceBaseStruct) error {

	err := s.generationHistoryRespository.DeleteBySetAndResource(model.Specifications.Set, model.Header.Name)
	if err != nil {
		return err
	}

	return nil
}
