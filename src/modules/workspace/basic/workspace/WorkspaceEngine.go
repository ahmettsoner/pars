package basic_workspace

import (
	"fmt"

	"github.com/sirupsen/logrus"

	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/workspace/basic_workspace/flows/create"
	remove_steps "parsdevkit.net/modules/workspace/basic_workspace/flows/remove"
	update_steps "parsdevkit.net/modules/workspace/basic_workspace/flows/update"
	describe_printer "parsdevkit.net/modules/workspace/basic_workspace/printers/describe"
	list_printer "parsdevkit.net/modules/workspace/basic_workspace/printers/list"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	_string "parsdevkit.net/pkg/utilities/string"
)

type WorkspaceEngine struct{}

func (s WorkspaceEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*basic_workspace_payload_structs.WorkspaceBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s WorkspaceEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s WorkspaceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s WorkspaceEngine) prepareToCreate(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToCreateStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, workspace := range workspaces {
		if err := s.completeInformation(ctx, &workspace); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(workspace.Header.Name)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, workspace)
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s WorkspaceEngine) create(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, workspaces)
	if err != nil {
		return err
	}

	for _, workspace := range readyToCreateStructs {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("CreateNewWorkspace").
			Step(&create_steps.SaveWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":      init,
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Create işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s WorkspaceEngine) prepareToUpdate(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToUpdateStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, workspace := range workspaces {
		if err := s.completeInformation(ctx, &workspace); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(workspace.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(workspace)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(workspace.Header.Name)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, workspace)
			}
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s WorkspaceEngine) update(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, workspaces)
	if err != nil {
		return err
	}
	for _, workspace := range readyToUpdateStructs {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("UpdateExistingWorkspace").
			Step(&update_steps.UpdateWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"init":      init,
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}
func (s WorkspaceEngine) prepareToRemove(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToRemoveStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, workspace := range workspaces {
		if err := s.completeInformation(ctx, &workspace); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(workspace.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, workspace)
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s WorkspaceEngine) remove(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, workspaces)
	if err != nil {
		return err
	}

	for _, workspace := range readyToRemoveStructs {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("RemoveExistingWorkspace").
			Step(&remove_steps.DeleteWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Remove işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s WorkspaceEngine) completeInformation(ctx *application.ApplicationContext, model *basic_workspace_payload_structs.WorkspaceBaseStruct) error {

	logrus.Debugf("filling workspace (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}

	return nil
}

func (s WorkspaceEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  basic_workspace_payload_structs.MODULE_KEY,
		Order: 1000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	r := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*basic_workspace_payload_structs.WorkspaceBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_workspace_payload_structs.WorkspaceBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
func (s WorkspaceEngine) List(ctx *application.ApplicationContext) error {
	err := s.list(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s WorkspaceEngine) prepareToList(ctx *application.ApplicationContext) ([]list_printer.ViewModel, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	var readyToListStructs []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	workspaceList, err := service.List()
	if err != nil {
		return nil, err
	}

	for _, e := range *workspaceList {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		readyToListStructs = append(readyToListStructs, resource)
	}

	return readyToListStructs, nil
}
func (s WorkspaceEngine) list(ctx *application.ApplicationContext) error {

	readyToListStructs, err := s.prepareToList(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("\n🛠️  Workspace List (%d):\n\n", len(readyToListStructs))

	printer := list_printer.ListWorkspace{Workspaces: readyToListStructs}
	printer.Print()

	return nil
}

func (s WorkspaceEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
	err := s.describe(ctx, args...)
	if err != nil {
		return err
	}

	return nil
}
func (s WorkspaceEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]describe_printer.ViewModel, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	var readyToDescribeStructs []describe_printer.ViewModel = make([]describe_printer.ViewModel, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			workspace, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if workspace != nil {

				resource := describe_printer.ViewModel{
					Name: workspace.Header.Name,
					Tags: workspace.Header.Metadata.Tags,
				}

				readyToDescribeStructs = append(readyToDescribeStructs, resource)
			} else {
				return nil, fmt.Errorf("xxx: Workspace '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Workspace argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s WorkspaceEngine) describe(ctx *application.ApplicationContext, args ...any) error {

	readyToDescribeStructs, err := s.prepareToDescribe(ctx, args...)
	if err != nil {
		return err
	}

	for _, workspace := range readyToDescribeStructs {

		fmt.Printf("\n🛠️  Details for: %s\n\n", workspace.Name)

		printer := describe_printer.DescribeWorkspace{Workspace: workspace}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Workspace Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}
