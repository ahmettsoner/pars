package services

import (
	"encoding/json"
	"errors"

	"parsdevkit.net/structs/resource"
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type ObjectResourceService struct {
	ObjectResourceServiceInterface
	resourceRepository           *repositories.ResourceRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewObjectResourceService(environment string) *ObjectResourceService {
	return &ObjectResourceService{
		environment:                  environment,
		resourceRepository:           repositories.NewResourceRepository(environment),
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s ObjectResourceService) GetByName(name string) (*objectresource.ResourceBaseStruct, error) {
	var resource *objectresource.ResourceBaseStruct

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

func (s ObjectResourceService) Save(model objectresource.ResourceBaseStruct) (*objectresource.ResourceBaseStruct, error) {

	result, err := s.saveResourceInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s ObjectResourceService) List() (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByKind(string(resource.StructKinds.Object))
	if err != nil {
		return nil, err
	}

	resourceList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}
func (s ObjectResourceService) ListByWorkspace(workspace string) (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndKind(workspace, string(resource.StructKinds.Object))
	if err != nil {
		return nil, err
	}

	resourceList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s ObjectResourceService) ListBySet(set string) (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySet(set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s ObjectResourceService) ListByWorkspaceAndSet(workspace, set string) (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceAndSet(workspace, set)
	if err != nil {
		return nil, err
	}

	resourceList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var resource objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &resource)

		resourceList = append(resourceList, resource)
	}

	return &resourceList, nil
}

func (s ObjectResourceService) ListBySetAndLayers(set string, layers ...string) (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s ObjectResourceService) ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]objectresource.ResourceBaseStruct), error) {

	entityList, err := s.resourceRepository.ListByWorkspaceSetAndLayers(workspace, set, layers...)
	if err != nil {
		return nil, err
	}

	templateList := make([]objectresource.ResourceBaseStruct, 0)

	for _, entity := range *entityList {
		var template objectresource.ResourceBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &template)

		templateList = append(templateList, template)
	}

	return &templateList, nil
}

func (s ObjectResourceService) Remove(name, workspace string, force, permanent bool) (*objectresource.ResourceBaseStruct, error) {
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
	if err != nil {
		return nil, err
	}

	var resource objectresource.ResourceBaseStruct
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

func (s ObjectResourceService) IsExists(name, workspace string) bool {

	resourceResourceEntity, err := s.resourceRepository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false
	}
	if resourceResourceEntity == nil {
		return false
	}

	return true
}
func (s ObjectResourceService) GetHash(name string) string {

	entity, err := s.resourceRepository.GetByName(name)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s ObjectResourceService) saveResourceInformation(resourceModel objectresource.ResourceBaseStruct) (*objectresource.ResourceBaseStruct, error) {

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
