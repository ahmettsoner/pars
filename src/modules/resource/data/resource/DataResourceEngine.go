package data_resource

import (
	"fmt"

	data_resource_payload_events "parsdevkit.net/modules/resource/data_resource_payload/events"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/resource/data_resource/flows/create"
	remove_steps "parsdevkit.net/modules/resource/data_resource/flows/remove"
	update_steps "parsdevkit.net/modules/resource/data_resource/flows/update"
	"parsdevkit.net/modules/resource/data_resource_contract"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	"parsdevkit.net/pkg/utilities/encrypt"

	"github.com/sirupsen/logrus"
)

type DataResourceEngine struct{}

func (s DataResourceEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*data_resource_payload_structs.ResourceBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s DataResourceEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s DataResourceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s DataResourceEngine) prepareToCreate(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()
	readyToCreateStructs := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s DataResourceEngine) create(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, resources)
	if err != nil {
		return err
	}

	for _, resource := range readyToCreateStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Creating: %s.%s\n\n", resource.Header.Name, resource.GetKey())

		resourceFlow := flowx.NewFlow("CreateNewResource").
			Step(&create_steps.SaveResource{}).
			Step(&create_steps.GenerateResourceContent{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"resource": resource,
		})

		if err := resourceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Resource Create işleminde hata oluştu: %w", &err)
		}

		bus.PublishEvent(data_resource_payload_events.ResourceCreated{
			Data: resource,
		})

	}

	return nil
}

func (s DataResourceEngine) prepareToUpdate(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()
	readyToUpdateStructs := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s DataResourceEngine) update(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, resources)
	if err != nil {
		return err
	}
	for _, resource := range readyToUpdateStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Updating: %s.%s\n\n", resource.Header.Name, resource.GetKey())

		resourceFlow := flowx.NewFlow("UpdateExistingResource").
			Step(&update_steps.UpdateResource{}).
			Step(&update_steps.GenerateResourceContent{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"resource": resource,
		})

		if err := resourceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Resource Update işleminde hata oluştu: %w", &err)
		}

		bus.PublishEvent(data_resource_payload_events.ResourceCreated{
			Data: resource,
		})
	}

	return nil
}
func (s DataResourceEngine) prepareToRemove(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()
	readyToRemoveStructs := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s DataResourceEngine) remove(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, resources)
	if err != nil {
		return err
	}

	for _, resource := range readyToRemoveStructs {

		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("📦 Removing: %s.%s\n\n", resource.Header.Name, resource.GetKey())

		resourceFlow := flowx.NewFlow("RemoveExistingGroup").
			Step(&remove_steps.DeleteResource{}).
			Step(&remove_steps.ClearResourceHistory{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"resource":  resource,
		})

		if err := resourceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Resource Remove işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s DataResourceEngine) completeInformation(ctx *application.ApplicationContext, model *data_resource_payload_structs.ResourceBaseStruct) error {

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
		model.Specifications.Layers = append(model.Specifications.Layers, data_resource_payload_structs.Layer{})
	}
	return nil
}

func (s DataResourceEngine) getWorkspace(ctx *application.ApplicationContext, model data_resource_payload_structs.ResourceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	workspaceName := model.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *basic_workspace_payload_structs.WorkspaceBaseStruct = nil

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
func (s DataResourceEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Resource.Data",
		Order: 3000,
	}
}

func CastArrayToConcrate(data []schemas.SchemaInterface) ([]data_resource_payload_structs.ResourceBaseStruct, error) {
	r := make([]data_resource_payload_structs.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*data_resource_payload_structs.ResourceBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected data_resource_payload_structs.ResourceBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
