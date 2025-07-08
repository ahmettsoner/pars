package object_resource

import (
	"fmt"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/resource/object_resource_payload"
	object_resource_payload_events "parsdevkit.net/modules/resource/object_resource_payload/events"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type ObjectResourceEngine struct{}

func (s ObjectResourceEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*object_resource_payload.ResourceBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s ObjectResourceEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.create(ctx, dataStruct, true)
	if err != nil {
		return err
	}
	err = s.update(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s ObjectResourceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.remove(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s ObjectResourceEngine) prepareToCreate(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct) ([]object_resource_payload.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToCreateStructs := make([]object_resource_payload.ResourceBaseStruct, 0)

	for _, resource := range resources {
		if err := s.completeInformation(ctx, &resource); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s ObjectResourceEngine) create(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct, init bool) error {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToCreateStructs, err := s.prepareToCreate(ctx, resources)
	if err != nil {
		return err
	}

	for index, resource := range readyToCreateStructs {

		logrus.Debugf("trying to create %v", resource.Header.Name)
		if _, err := service.Save(resource); err != nil {
			return err
		}

		fmt.Printf("%v (%d) Object Resource created\n", resource.Header.Name, index)

		// err := bus.SendCommand(application_project_payload_commands.CreateApplicationProject{})
		// if err != nil {
		// 	panic(err)
		// }
		bus.PublishEvent(object_resource_payload_events.ResourceCreated{
			Data: resource,
		})
	}

	return nil
}

func (s ObjectResourceEngine) prepareToUpdate(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct) ([]object_resource_payload.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToUpdateStructs := make([]object_resource_payload.ResourceBaseStruct, 0)

	for _, resource := range resources {
		if err := s.completeInformation(ctx, &resource); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(resource)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(resource.Header.Name)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, resource)
			}
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s ObjectResourceEngine) update(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct, init bool) error {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, resources)
	if err != nil {
		return err
	}
	for _, resource := range readyToUpdateStructs {
		if _, err := service.Save(resource); err != nil {
			return err
		}

		bus.PublishEvent(object_resource_payload_events.ResourceCreated{
			Data: resource,
		})
	}
	return nil
}
func (s ObjectResourceEngine) prepareToRemove(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct) ([]object_resource_payload.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToRemoveStructs := make([]object_resource_payload.ResourceBaseStruct, 0)

	for _, resource := range resources {
		if err := s.completeInformation(ctx, &resource); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s ObjectResourceEngine) remove(ctx *application.ApplicationContext, resources []object_resource_payload.ResourceBaseStruct, permanent bool) error {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	readyToRemoveStructs, err := s.prepareToRemove(ctx, resources)
	if err != nil {
		return err
	}

	for _, resource := range readyToRemoveStructs {

		if _, err := service.Remove(resource.Header.Name, resource.Specifications.Workspace, true, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", resource.Header.Name)

	}

	logrus.Debugf("'%d' resource(s) deleting", len(readyToRemoveStructs))

	return nil
}

func (s ObjectResourceEngine) completeInformation(ctx *application.ApplicationContext, model *object_resource_payload.ResourceBaseStruct) error {

	logrus.Debugf("filling model (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}

	activeWorkspace, err := s.getWorkspace(ctx, *model)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	model.Specifications.Workspace = activeWorkspace.Header.Name
	model.Specifications.WorkspaceObject = activeWorkspace.Specifications.WorkspaceIdentifier
	logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Header.Name, model.Header.Name)

	if len(model.Specifications.Layers) == 0 {
		model.Specifications.Layers = append(model.Specifications.Layers, object_resource_payload.Layer{})
	}

	return nil
}

func (s ObjectResourceEngine) getWorkspace(ctx *application.ApplicationContext, model object_resource_payload.ResourceBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {

	workspaceName := model.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *workspaceStruct.WorkspaceBaseStruct = nil

	if !_string.IsEmpty(workspaceName) {
		workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = workspace
	} else {
	}

	return result, nil
}

func (s ObjectResourceEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Resource.Object",
		Order: 3000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]object_resource_payload.ResourceBaseStruct, error) {
	r := make([]object_resource_payload.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*object_resource_payload.ResourceBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected object_resource_payload.ResourceBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
