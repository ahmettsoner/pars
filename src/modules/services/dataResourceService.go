package services

import (
	"encoding/json"
	"errors"

	"parsdevkit.net/structs/resource"
	dataresource "parsdevkit.net/structs/resource/data-resource"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type DataResourceService struct {
	DataResourceServiceInterface
	resourceRepository           *repositories.ResourceRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewDataResourceService(environment string) *DataResourceService {
	return &DataResourceService{
		environment:                  environment,
		resourceRepository:           repositories.NewResourceRepository(environment),
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s DataResourceService) GetByName(name string) (*dataresource.ResourceBaseStruct, error) {
	var resource *dataresource.ResourceBaseStruct

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

func (s DataResourceService) Save(model dataresource.ResourceBaseStruct) (*dataresource.ResourceBaseStruct, error) {

	result, err := s.saveResourceInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s DataResourceService) List() (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByKind(string(resource.StructKinds.Data))
	if err != nil {
		return nil, err
	}

	resourceList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}
func (s DataResourceService) ListByWorkspace(workspace string) (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndKind(workspace, string(resource.StructKinds.Data))
	if err != nil {
		return nil, err
	}

	resourceList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s DataResourceService) ListBySet(set string) (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySet(set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}
func (s DataResourceService) ListByWorkspaceAndSet(workspace, set string) (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndSet(workspace, set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s DataResourceService) ListBySetAndLayers(set string, layers ...string) (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}
func (s DataResourceService) ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]dataresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]dataresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template dataresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s DataResourceService) Remove(name, workspace string, force, permanent bool) (*dataresource.ResourceBaseStruct, error) {
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
	var resource dataresource.ResourceBaseStruct
	err = json.Unmarshal([]byte(resourceResourceEntity.Document), &resource)
	if err != nil {
		return nil, err
	}

	err = s.generationHistoryRespository.DeleteBySetAndResource(resource.Specifications.Set, resource.Name)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (s DataResourceService) IsExists(name, workspace string) bool {

	resourceResourceEntity, err := s.resourceRepository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false
	}
	if resourceResourceEntity == nil {
		return false
	}

	return true
}
func (s DataResourceService) GetHash(name string) string {

	entity, err := s.resourceRepository.GetByName(name)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s DataResourceService) saveResourceInformation(resourceModel dataresource.ResourceBaseStruct) (*dataresource.ResourceBaseStruct, error) {

	jsonData, err := json.Marshal(resourceModel)
	if err != nil {
		return nil, err
	}

	resourceEntity := entities.Resource{
		Name:     resourceModel.Name,
		Document: string(jsonData),
	}

	err = s.resourceRepository.Save(&resourceEntity)
	if err != nil {
		return nil, err
	}

	return &resourceModel, nil
}
