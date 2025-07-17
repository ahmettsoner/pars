package data_resource

import (
	"fmt"

	layerPkg "parsdevkit.net/application/models/layer"
	data_resource_payload_events "parsdevkit.net/modules/resource/data_resource_payload/events"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/resource/data_resource/flows/create"
	remove_steps "parsdevkit.net/modules/resource/data_resource/flows/remove"
	update_steps "parsdevkit.net/modules/resource/data_resource/flows/update"
	describe_printer "parsdevkit.net/modules/resource/data_resource/printers/describe"
	list_printer "parsdevkit.net/modules/resource/data_resource/printers/list"
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

	readyToCreateStructs, err := s.prepareToCreate(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.create(ctx, readyToCreateStructs, true)
	if err != nil {
		return err
	}

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, dataStruct)
	if err != nil {
		return err
	}
	err = s.update(ctx, readyToUpdateStructs, true)
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
func (s DataResourceEngine) create(ctx *application.ApplicationContext, models []data_resource_payload_structs.ResourceBaseStruct, init bool) error {

	for _, resource := range models {

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
func (s DataResourceEngine) update(ctx *application.ApplicationContext, models []data_resource_payload_structs.ResourceBaseStruct, init bool) error {

	for _, resource := range models {

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
			return fmt.Errorf("xxx: Resource Update işleminde hata oluştu: %w", &err)
		}

		bus.PublishEvent(data_resource_payload_events.ResourceCreated{
			Data: resource,
		})
	}

	return nil
}

func (s DataResourceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	readyToDestroyStructs, err := s.prepareToDestroy(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.remove(ctx, readyToDestroyStructs, true)
	if err != nil {
		return err
	}

	return nil
}

func (s DataResourceEngine) prepareToDestroy(ctx *application.ApplicationContext, resources []data_resource_payload_structs.ResourceBaseStruct) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()
	readyToDestroyStructs := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, resource := range resources {
		if err := s.completeInformation(ctx, &resource); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(resource.Header.Name, resource.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToDestroyStructs = append(readyToDestroyStructs, resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will destroy", len(readyToDestroyStructs))

	return readyToDestroyStructs, nil
}
func (s DataResourceEngine) remove(ctx *application.ApplicationContext, models []data_resource_payload_structs.ResourceBaseStruct, permanent bool) error {

	for _, resource := range models {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", resource.Header.Name, resource.GetKey())

		resourceFlow := flowx.NewFlow("DestroyExistingResource").
			Step(&remove_steps.DeleteResource{}).
			Step(&remove_steps.ClearResourceHistory{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"resource":  resource,
		})

		if err := resourceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Resource Destroy işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s DataResourceEngine) List(ctx *application.ApplicationContext) error {
	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	err = s.list(ctx, readyToListStructs)
	if err != nil {
		return err
	}

	return nil
}
func (s DataResourceEngine) prepareToList(ctx *application.ApplicationContext) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()

	resourceList, err := service.List()
	if err != nil {
		return nil, err
	}

	return *resourceList, nil
}
func (s DataResourceEngine) list(ctx *application.ApplicationContext, models []data_resource_payload_structs.ResourceBaseStruct) error {

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)

	for _, e := range *&models {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  Data Resource List (%d):\n\n", len(viewModels))

	printer := list_printer.ListResource{Resources: viewModels}
	printer.Print()

	return nil
}

func (s DataResourceEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
	readyToDescribeStructs, err := s.prepareToDescribe(ctx, args...)
	if err != nil {
		return err
	}

	err = s.describe(ctx, readyToDescribeStructs)
	if err != nil {
		return err
	}

	return nil
}
func (s DataResourceEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()

	var readyToDescribeStructs []data_resource_payload_structs.ResourceBaseStruct = make([]data_resource_payload_structs.ResourceBaseStruct, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			resource, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if resource != nil && resource.Header.Kind == data_resource_payload_structs.RESOURCE_KIND {
				readyToDescribeStructs = append(readyToDescribeStructs, *resource)

			} else {
				return nil, fmt.Errorf("xxx: Data Resource '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Data Resource argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s DataResourceEngine) describe(ctx *application.ApplicationContext, models []data_resource_payload_structs.ResourceBaseStruct) error {

	for _, resource := range models {

		viewModel := describe_printer.ViewModel{
			Name: resource.Header.Name,
			Tags: resource.Header.Metadata.Tags,
		}

		fmt.Printf("\n🛠️  Details for: %s\n\n", viewModel.Name)

		printer := describe_printer.DescribeResource{Resource: viewModel}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Data Resource Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}
func (s DataResourceEngine) Remove(ctx *application.ApplicationContext, args ...any) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, args...)
	if err != nil {
		return err
	}

	err = s.remove(ctx, readyToRemoveStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s DataResourceEngine) prepareToRemove(ctx *application.ApplicationContext, args ...any) ([]data_resource_payload_structs.ResourceBaseStruct, error) {

	service := ioc.Get[data_resource_contract.ResourceInterface]()
	readyToRemoveStructs := make([]data_resource_payload_structs.ResourceBaseStruct, 0)

	for _, a := range args {
		resource, err := service.GetByName(a.(string))
		if err != nil {
			return nil, err
		}
		if resource != nil && resource.Header.Kind == data_resource_payload_structs.RESOURCE_KIND {
			readyToRemoveStructs = append(readyToRemoveStructs, *resource)
		}
	}
	logrus.Debugf("'%d' resource(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
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
		model.Specifications.Layers = append(model.Specifications.Layers, layerPkg.Layer{})
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
		Name:  data_resource_payload_structs.MODULE_KEY,
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
