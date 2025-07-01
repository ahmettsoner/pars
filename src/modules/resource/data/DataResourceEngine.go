package v2

import (
	"fmt"

	"parsdevkit.net/application"
	dataresourceStruct "parsdevkit.net/structs/resource/data-resource"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utilities/encrypt"
	"parsdevkit.net/core/utilities/json"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines"

	"github.com/sirupsen/logrus"
)

type DataResourceEngine struct{}

func (s DataResourceEngine) Validate(data []schemas.Schema) bool {
	for _, item := range data {
		_, ok := item.(*dataresourceStruct.ResourceBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s DataResourceEngine) Process(ctx *application.ApplicationContext, data []schemas.Schema) error {
	dataresourceStructs := make([]dataresourceStruct.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		dataresourceStruct, ok := item.(*dataresourceStruct.ResourceBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected dataresourceStruct.ResourceBaseStruct, got %T", item)
		}

		dataresourceStructs = append(dataresourceStructs, *dataresourceStruct)
	}

	return s.createResources(dataresourceStructs, true)
}
func (s DataResourceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.Schema) error {
	dataresourceStructs := make([]dataresourceStruct.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		dataresourceStruct, ok := item.(*dataresourceStruct.ResourceBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected dataresourceStruct.ResourceBaseStruct, got %T", item)
		}

		dataresourceStructs = append(dataresourceStructs, *dataresourceStruct)
	}

	return s.removeResources(dataresourceStructs, true)
}

func (s DataResourceEngine) createResources(resources []dataresourceStruct.ResourceBaseStruct, init bool) error {

	resourcesReadyToCreate := make([]dataresourceStruct.ResourceBaseStruct, 0)
	resourcesForUpdate := make([]dataresourceStruct.ResourceBaseStruct, 0)
	resourceService := services.NewDataResourceService(utils.GetEnvironment())

	for _, resource := range resources {
		if err := resource.Validate(); err != nil {
			jsonObject, _ := json.ToJson(resource)
			return fmt.Errorf("resource invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, resource := range resources {
		ok, err := resourceService.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Data Resource ('%s') kontrolünde hata oluştu\n%w", resource.Header.Name, err)
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(resource)
			if err != nil {
				return err
			}
			structHash, err := resourceService.GetHash(resource.Header.Name)
			if err != nil {
				return err
			}

			if newModelHash != structHash {
				resourcesForUpdate = append(resourcesForUpdate, resource)
			}
		} else {
			resourcesReadyToCreate = append(resourcesReadyToCreate, resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will create", len(resourcesReadyToCreate))
	logrus.Debugf("'%d' resource(s) detected that will update", len(resourcesForUpdate))

	logrus.Debugf("creating %v new resources ", len(resourcesReadyToCreate))
	logrus.Debugf("updating %v resources ", len(resourcesForUpdate))
	for _, resource := range resourcesReadyToCreate {

		if _, err := resourceService.Save(resource); err != nil {
			return err
		}

		if _, err := s.generate(resource); err != nil {
			return err
		}

		fmt.Printf("%v Resource created\n", resource.Header.Name)
	}

	logrus.Debugf("updating %v resources ", len(resourcesForUpdate))
	for _, resource := range resourcesForUpdate {

		if _, err := resourceService.Save(resource); err != nil {
			return err
		}

		fmt.Printf("%v Resource updated\n", resource.Header.Name)
	}

	return nil
}
func (s DataResourceEngine) removeResources(resources []dataresourceStruct.ResourceBaseStruct, permanent bool) error {

	resourceService := services.NewDataResourceService(utils.GetEnvironment())
	resourcesReadyToDelete := make([]dataresourceStruct.ResourceBaseStruct, 0)
	for _, resource := range resources {
		ok, err := resourceService.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Data Resource ('%s') kontrolünde hata oluştu\n%w", resource.Header.Name, err)
		}
		if ok {
			resourcesReadyToDelete = append(resourcesReadyToDelete, resource)
		}
	}

	for _, resource := range resourcesReadyToDelete {

		if _, err := resourceService.Remove(resource.Header.Name, resource.Specifications.Workspace, true, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Resource deleted\n", resource.Header.Name)

	}

	return nil
}

func (s DataResourceEngine) generate(model dataresourceStruct.ResourceBaseStruct) (*dataresourceStruct.ResourceBaseStruct, error) {

	resourceService := services.NewDataResourceService(utils.GetEnvironment())

	result, err := resourceService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	templateOperations := engines.NewFileTemplateOperations(utils.GetEnvironment())
	err = templateOperations.GenerateByResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
