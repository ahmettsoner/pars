package basic_task

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/modules/task/basic_task_contract"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/task/basic_task/flows/create"
	remove_steps "parsdevkit.net/modules/task/basic_task/flows/remove"
	update_steps "parsdevkit.net/modules/task/basic_task/flows/update"
	describe_printer "parsdevkit.net/modules/task/basic_task/printers/describe"
	list_printer "parsdevkit.net/modules/task/basic_task/printers/list"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type BasicTaskEngine struct{}

func (s BasicTaskEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*basic_task_payload_structs.TaskBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s BasicTaskEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s BasicTaskEngine) prepareToCreate(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToCreateStructs := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, task := range tasks {
		if err := s.completeInformation(ctx, &task); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, task)
		}
	}
	logrus.Debugf("'%d' task(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s BasicTaskEngine) create(ctx *application.ApplicationContext, models []basic_task_payload_structs.TaskBaseStruct, init bool) error {

	for _, task := range models {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", task.Header.Name, task.GetKey())

		taskFlow := flowx.NewFlow("CreateNewTask").
			Step(&create_steps.SaveTask{})

		fc := flowx.NewContextWithData(map[string]any{
			"init": init,
			"task": task,
		})

		if err := taskFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Task Create işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s BasicTaskEngine) prepareToUpdate(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToUpdateStructs := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, task := range tasks {
		if err := s.completeInformation(ctx, &task); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(task)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(task.Header.Name)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, task)
			}
		}
	}
	logrus.Debugf("'%d' task(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s BasicTaskEngine) update(ctx *application.ApplicationContext, models []basic_task_payload_structs.TaskBaseStruct, init bool) error {

	for _, task := range models {

		fmt.Printf("\n🛠️  Updating: %s.%s\n\n", task.Header.Name, task.GetKey())

		taskFlow := flowx.NewFlow("UpdateExistingTask").
			Step(&update_steps.UpdateTask{})

		fc := flowx.NewContextWithData(map[string]any{
			"init": init,
			"task": task,
		})

		if err := taskFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Task Update işleminde hata oluştu: %w", &err)
		}
		if _, err := s.execute(task); err != nil {
			return err
		}
	}
	return nil
}

func (s BasicTaskEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s BasicTaskEngine) prepareToDestroy(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToDestroyStructs := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, task := range tasks {
		if err := s.completeInformation(ctx, &task); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToDestroyStructs = append(readyToDestroyStructs, task)
		}
	}
	logrus.Debugf("'%d' task(s) detected that will destroy", len(readyToDestroyStructs))

	return readyToDestroyStructs, nil
}
func (s BasicTaskEngine) remove(ctx *application.ApplicationContext, models []basic_task_payload_structs.TaskBaseStruct, permanent bool) error {

	for _, task := range models {

		fmt.Printf("\n🛠️  Removing: %s.%s\n\n", task.Header.Name, task.GetKey())

		taskFlow := flowx.NewFlow("DestroyExistingTask").
			Step(&remove_steps.DeleteTask{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"task":      task,
		})

		if err := taskFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Task Destroy işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s BasicTaskEngine) List(ctx *application.ApplicationContext) error {
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
func (s BasicTaskEngine) prepareToList(ctx *application.ApplicationContext) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	taskList, err := service.List()
	if err != nil {
		return nil, err
	}

	return *taskList, nil
}
func (s BasicTaskEngine) list(ctx *application.ApplicationContext, models []basic_task_payload_structs.TaskBaseStruct) error {

	var viewModels []list_printer.ViewModel = make([]list_printer.ViewModel, 0)
	for _, e := range models {
		resource := list_printer.ViewModel{
			Name: e.Header.Name,
			Tags: e.Header.Metadata.Tags,
		}

		viewModels = append(viewModels, resource)
	}

	fmt.Printf("\n🛠️  Task List (%d):\n\n", len(viewModels))

	printer := list_printer.ListTask{Tasks: viewModels}
	printer.Print()

	return nil
}

func (s BasicTaskEngine) Describe(ctx *application.ApplicationContext, args ...any) error {
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
func (s BasicTaskEngine) prepareToDescribe(ctx *application.ApplicationContext, args ...any) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	var readyToDescribeStructs []basic_task_payload_structs.TaskBaseStruct = make([]basic_task_payload_structs.TaskBaseStruct, 0)
	for _, v := range args {
		if a, ok := v.(string); ok {
			task, err := service.GetByName(a)
			if err != nil {
				return nil, err
			}
			if task != nil && task.Header.Kind == basic_task_payload_structs.TASK_KIND {

				readyToDescribeStructs = append(readyToDescribeStructs, *task)

			} else {
				return nil, fmt.Errorf("xxx: Task '%s' bulunamadı", a)
			}
		} else {
			return nil, fmt.Errorf("xxx: Task argümanı doğru değil")
		}
	}

	return readyToDescribeStructs, nil
}
func (s BasicTaskEngine) describe(ctx *application.ApplicationContext, models []basic_task_payload_structs.TaskBaseStruct) error {

	for _, task := range models {

		viewModel := describe_printer.ViewModel{
			Name: task.Header.Name,
			Tags: task.Header.Metadata.Tags,
		}

		fmt.Printf("\n🛠️  Details for: %s\n\n", viewModel.Name)

		printer := describe_printer.DescribeTask{Task: viewModel}

		if err := printer.Print(); err != nil {
			return fmt.Errorf("xxx: Task Print sırasında hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s BasicTaskEngine) Remove(ctx *application.ApplicationContext, args ...any) error {

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
func (s BasicTaskEngine) prepareToRemove(ctx *application.ApplicationContext, args ...any) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToRemoveStructs := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, a := range args {
		task, err := service.GetByName(a.(string))
		if err != nil {
			return nil, err
		}
		if task != nil && task.Header.Kind == basic_task_payload_structs.TASK_KIND {
			readyToRemoveStructs = append(readyToRemoveStructs, *task)
		}
	}
	logrus.Debugf("'%d' task(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s BasicTaskEngine) execute(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	taskService := ioc.Get[basic_task_contract.TaskInterface]()

	result, err := taskService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result, nil
}

func (s BasicTaskEngine) completeInformation(ctx *application.ApplicationContext, model *basic_task_payload_structs.TaskBaseStruct) error {

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

	return nil
}

func (s BasicTaskEngine) getWorkspace(ctx *application.ApplicationContext, model basic_task_payload_structs.TaskBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {

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

func (s BasicTaskEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  basic_task_payload_structs.MODULE_KEY,
		Order: 5000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]basic_task_payload_structs.TaskBaseStruct, error) {
	r := make([]basic_task_payload_structs.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*basic_task_payload_structs.TaskBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected basic_task_payload_structs.TaskBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
