package object_resource

import (
	"fmt"

	engineOperations "parsdevkit.net/engines"
	"parsdevkit.net/modules/resource/object_resource_contract"
	objectresourceStruct "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"

	"parsdevkit.net/pkg/utilities/json"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type ObjectResourceEngine struct{}

func (s ObjectResourceEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*objectresourceStruct.ResourceBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s ObjectResourceEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	objectresourceStructs := make([]objectresourceStruct.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		objectresourceStruct, ok := item.(*objectresourceStruct.ResourceBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected objectresourceStruct.ResourceBaseStruct, got %T", item)
		}

		objectresourceStructs = append(objectresourceStructs, *objectresourceStruct)
	}

	return s.createResources(objectresourceStructs, true)
}
func (s ObjectResourceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	objectresourceStructs := make([]objectresourceStruct.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		objectresourceStruct, ok := item.(*objectresourceStruct.ResourceBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected objectresourceStruct.ResourceBaseStruct, got %T", item)
		}

		objectresourceStructs = append(objectresourceStructs, *objectresourceStruct)
	}

	return s.removeResources(objectresourceStructs, true)
}

func (s ObjectResourceEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Resource.Object",
		Order: 3000,
	}
}
func (s ObjectResourceEngine) createResources(resources []objectresourceStruct.ResourceBaseStruct, init bool) error {

	resourcesReadyToCreate := make([]objectresourceStruct.ResourceBaseStruct, 0)
	resourcesForUpdate := make([]objectresourceStruct.ResourceBaseStruct, 0)
	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()

	for _, resource := range resources {
		if err := resource.Validate(); err != nil {
			jsonObject, _ := json.ToJson(resource)
			return fmt.Errorf("resource invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, resource := range resources {
		ok, err := resourceService.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Object Resource ('%s') kontrolünde hata oluştu\n%w", resource.Header.Name, err)
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

		if _, err := s.generate(resource); err != nil {
			return err
		}

		fmt.Printf("%v Resource updated\n", resource.Header.Name)
	}

	return nil
}

func (s ObjectResourceEngine) removeResources(resources []objectresourceStruct.ResourceBaseStruct, permanent bool) error {

	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()
	resourcesReadyToDelete := make([]objectresourceStruct.ResourceBaseStruct, 0)
	for _, resource := range resources {
		ok, err := resourceService.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Object Resource ('%s') kontrolünde hata oluştu\n%w", resource.Header.Name, err)
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
func (s ObjectResourceEngine) generate(model objectresourceStruct.ResourceBaseStruct) (*objectresourceStruct.ResourceBaseStruct, error) {

	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()

	result, err := resourceService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	// TODO: Birden fazla template işlenebilmeli
	templateEngine := engineOperations.NewCodeTemplateOperations(application.GetEnvironment())
	err = templateEngine.GenerateByResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
