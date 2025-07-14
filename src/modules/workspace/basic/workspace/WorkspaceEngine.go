package basic_workspace

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"parsdevkit.net/modules/project/application_project_contract"
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
func (s WorkspaceEngine) create(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct, init bool) error {

	for _, workspace := range models {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("CreateNewWorkspace").
			Step(&create_steps.SaveWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
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
func (s WorkspaceEngine) update(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct, init bool) error {

	for _, workspace := range models {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("UpdateExistingWorkspace").
			Step(&update_steps.UpdateWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}

func (s WorkspaceEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s WorkspaceEngine) prepareToDestroy(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToDestroyStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, workspace := range workspaces {
		if err := s.completeInformation(ctx, &workspace); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(workspace.Header.Name)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToDestroyStructs = append(readyToDestroyStructs, workspace)
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will destroy", len(readyToDestroyStructs))

	return readyToDestroyStructs, nil
}
func (s WorkspaceEngine) remove(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct, permanent bool) error {

	for _, workspace := range models {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("DestroyExistingWorkspace").
			Step(&remove_steps.DeleteWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Destroy işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s WorkspaceEngine) Init(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate2(data)
	if err != nil {
		return err
	}

	readyToInitStructs, err := s.prepareToInit(ctx, dataStruct)
	if err != nil {
		return err
	}

	err = s.init(ctx, readyToInitStructs, true)
	if err != nil {
		return err
	}

	return nil
}
func (s WorkspaceEngine) prepareToInit(ctx *application.ApplicationContext, workspaces []basic_workspace_payload_structs.WorkspaceBaseStruct) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToInitStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, workspace := range workspaces {
		if err := s.completeInformation(ctx, &workspace); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(workspace.Header.Name)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToInitStructs = append(readyToInitStructs, workspace)
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will init", len(readyToInitStructs))

	return readyToInitStructs, nil
}
func (s WorkspaceEngine) init(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct, permanent bool) error {

	for _, workspace := range models {

		fmt.Printf("\n🛠️  Initializing: %s.%s\n\n", workspace.Header.Name, workspace.GetKey())

		workspaceFlow := flowx.NewFlow("InitializeWorkspace").
			Step(&create_steps.SaveWorkspace{})

		fc := flowx.NewContextWithData(map[string]any{
			"workspace": workspace,
		})

		if err := workspaceFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Workspace Destroy işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s WorkspaceEngine) List(ctx *application.ApplicationContext) error {
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
func (s WorkspaceEngine) prepareToList(ctx *application.ApplicationContext) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	workspaceList, err := service.List()
	if err != nil {
		return nil, err
	}

	return *workspaceList, nil
}
func (s WorkspaceEngine) list(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct) error {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)

	activeWorkspace, err := service.GetActiveWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Active Workspace\n%w", err)
	}

	selectedWorkspace, err := service.GetSelectedWorkspace()
	if err != nil {
		return fmt.Errorf("Failed to find Selected Workspace\n%w", err)
	}

	for _, e := range models {
		name := e.Header.Name
		label := name

		if activeWorkspace != nil && selectedWorkspace != nil && activeWorkspace.Header.Name == name && selectedWorkspace.Header.Name == name {
			label = fmt.Sprintf("* %v (active & selected)", name)
		} else if activeWorkspace != nil && activeWorkspace.Header.Name == name {
			label = fmt.Sprintf("* %v (active)", name)
		} else if selectedWorkspace != nil && selectedWorkspace.Header.Name == name {
			label = fmt.Sprintf("* %v (selected)", name)
		}

		resource := list_printer.ViewModel{
			Name: label,
			Tags: e.Header.Metadata.Tags,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  Workspace List (%d):\n\n", len(viewModels))

	printer := list_printer.ListWorkspace{Workspaces: viewModels}
	printer.Print()

	return nil
}

func (s WorkspaceEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
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
func (s WorkspaceEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	var readyToDescribeStructs []basic_workspace_payload_structs.WorkspaceBaseStruct = make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			workspace, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}

			if workspace != nil && workspace.Header.Kind == basic_workspace_payload_structs.WORKSPACE_KIND {

				readyToDescribeStructs = append(readyToDescribeStructs, *workspace)
			} else {
				return nil, fmt.Errorf("xxx: Workspace '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Workspace argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s WorkspaceEngine) describe(ctx *application.ApplicationContext, models []basic_workspace_payload_structs.WorkspaceBaseStruct) error {

	projectService := ioc.Get[application_project_contract.ProjectInterface]()

	for _, workspace := range models {

		projectList, err := projectService.ListByWorkspace(workspace.Specifications.Name)
		if err != nil {
			return fmt.Errorf("Failed to retrieve workspace projects '%s'\n%w", workspace.Header.Name, err)
		}

		resourceProjects := []describe_printer.ProjectViewModel{}

		for _, v := range *projectList {

			resourceProjectLabels := make([]string, 0)
			for _, v := range v.Specifications.Labels {
				resourceLabel := v.Key
				if !_string.IsEmpty(v.Value) {
					resourceLabel = fmt.Sprintf("%v=%v", v.Key, v.Value)
				}
				resourceProjectLabels = append(resourceProjectLabels, resourceLabel)
			}

			resource := describe_printer.ProjectViewModel{
				Name:   v.Header.Name,
				Set:    v.Specifications.Set,
				Group:  v.Specifications.Group,
				Tags:   v.Header.Metadata.Tags,
				Labels: resourceProjectLabels,
			}

			resourceProjects = append(resourceProjects, resource)
		}

		resource := describe_printer.ViewModel{
			Name:     workspace.Header.Name,
			Tags:     workspace.Header.Metadata.Tags,
			Path:     workspace.Specifications.Path,
			Projects: resourceProjects,
		}

		fmt.Printf("\n🛠️  Details for: %s\n\n", resource.Name)

		printer := describe_printer.DescribeWorkspace{Workspace: resource}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Workspace Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s WorkspaceEngine) Remove(ctx *application.ApplicationContext, args ...any) error {

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
func (s WorkspaceEngine) prepareToRemove(ctx *application.ApplicationContext, args ...any) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	readyToRemoveStructs := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0)

	for _, a := range args {
		workspace, err := service.GetByName(a.(string))
		if err != nil {
			return nil, err
		}
		if workspace != nil && workspace.Header.Kind == basic_workspace_payload_structs.WORKSPACE_KIND {
			readyToRemoveStructs = append(readyToRemoveStructs, *workspace)
		}
	}
	logrus.Debugf("'%d' workspace(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
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

func CastArrayToConcrate2(data []schemas.SchemaInterface) ([]basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	r := make([]basic_workspace_payload_structs.WorkspaceBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(basic_workspace_payload_structs.WorkspaceBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_workspace_payload_structs.WorkspaceBaseStruct, got %T", item)
		}

		r = append(r, model)
	}
	return r, nil
}
