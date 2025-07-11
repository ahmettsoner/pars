package object_resource

import (
	"fmt"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/resource/object_resource/flows/create"
	remove_steps "parsdevkit.net/modules/resource/object_resource/flows/remove"
	update_steps "parsdevkit.net/modules/resource/object_resource/flows/update"
	describe_printer "parsdevkit.net/modules/resource/object_resource/printers/describe"
	list_printer "parsdevkit.net/modules/resource/object_resource/printers/list"
	object_resource_payload_events "parsdevkit.net/modules/resource/object_resource_payload/events"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type ObjectResourceEngine struct{}

func (s ObjectResourceEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*object_resource_payload_structs.ResourceBaseStruct)
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
func (s ObjectResourceEngine) prepareToCreate(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct) ([]object_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToCreateStructs := make([]object_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s ObjectResourceEngine) create(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, resources)
	if err != nil {
		return err
	}

	for _, resource := range readyToCreateStructs {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", resource.Header.Name, resource.GetKey())

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

		bus.PublishEvent(object_resource_payload_events.ResourceCreated{Data: resource})
	}

	return nil
}

func (s ObjectResourceEngine) prepareToUpdate(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct) ([]object_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToUpdateStructs := make([]object_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s ObjectResourceEngine) update(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, resources)
	if err != nil {
		return err
	}
	for _, resource := range readyToUpdateStructs {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", resource.Header.Name, resource.GetKey())

		resourceFlow := flowx.NewFlow("UpdateExistingResource").
			Step(&update_steps.UpdateResource{}).
			Step(&update_steps.GenerateResourceContent{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":     init,
			"resource": resource,
		})

		if err := resourceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Resource Create işleminde hata oluştu: %w", &err)
		}

	}
	return nil
}
func (s ObjectResourceEngine) prepareToRemove(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct) ([]object_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()
	readyToRemoveStructs := make([]object_resource_payload_structs.ResourceBaseStruct, 0)

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
func (s ObjectResourceEngine) remove(ctx *application.ApplicationContext, resources []object_resource_payload_structs.ResourceBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, resources)
	if err != nil {
		return err
	}

	for _, resource := range readyToRemoveStructs {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", resource.Header.Name, resource.GetKey())

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

func (s ObjectResourceEngine) completeInformation(ctx *application.ApplicationContext, model *object_resource_payload_structs.ResourceBaseStruct) error {

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
		model.Specifications.Layers = append(model.Specifications.Layers, object_resource_payload_structs.Layer{})
	}

	return nil
}

func (s ObjectResourceEngine) getWorkspace(ctx *application.ApplicationContext, model object_resource_payload_structs.ResourceBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

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

func (s ObjectResourceEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  object_resource_payload_structs.MODULE_KEY,
		Order: 3000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]object_resource_payload_structs.ResourceBaseStruct, error) {
	r := make([]object_resource_payload_structs.ResourceBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*object_resource_payload_structs.ResourceBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected object_resource_payload_structs.ResourceBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}

func (s ObjectResourceEngine) List(ctx *application.ApplicationContext) error {
	err := s.list(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s ObjectResourceEngine) prepareToList(ctx *application.ApplicationContext) ([]list_printer.ViewModel, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	var readyToListStructs []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	groupList, err := service.List()
	if err != nil {
		return nil, err
	}

	for _, e := range *groupList {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		readyToListStructs = append(readyToListStructs, resource)
	}

	return readyToListStructs, nil
}
func (s ObjectResourceEngine) list(ctx *application.ApplicationContext) error {

	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("\n🛠️  Object Resource List (%d):\n\n", len(readyToListStructs))

	printer := list_printer.ListResource{Resources: readyToListStructs}
	printer.Print()

	return nil
}

func (s ObjectResourceEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
	err := s.describe(ctx, args...)
	if err != nil {
		return err
	}

	return nil
}
func (s ObjectResourceEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]describe_printer.ViewModel, error) {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	var readyToDescribeStructs []describe_printer.ViewModel = make([]describe_printer.ViewModel, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			group, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if group != nil {

				resource := describe_printer.ViewModel{
					Name: group.Header.Name,
					Tags: group.Header.Metadata.Tags,
				}

				readyToDescribeStructs = append(readyToDescribeStructs, resource)
			} else {
				return nil, fmt.Errorf("xxx: Object Resource '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Object Resource argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s ObjectResourceEngine) describe(ctx *application.ApplicationContext, args ...any) error {

	readyToDescribeStructs, err := s.prepareToDescribe(ctx, args...)
	if err != nil {
		return err
	}

	for _, group := range readyToDescribeStructs {

		fmt.Printf("\n🛠️  Details for: %s\n\n", group.Name)

		printer := describe_printer.DescribeResource{Resource: group}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Object Resource Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}
