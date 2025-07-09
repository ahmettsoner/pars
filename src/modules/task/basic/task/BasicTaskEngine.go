package basic_task

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/modules/task/basic_task_contract"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/internal/flowx"
	create_steps "parsdevkit.net/modules/task/basic_task/flows/create"
	remove_steps "parsdevkit.net/modules/task/basic_task/flows/remove"
	update_steps "parsdevkit.net/modules/task/basic_task/flows/update"
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
func (s BasicTaskEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
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
func (s BasicTaskEngine) create(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, tasks)
	if err != nil {
		return err
	}

	for _, task := range readyToCreateStructs {

		fmt.Printf("════════════════════════════════════\n")
		fmt.Printf("📦 Processing: %s.%s\n", task.GetKey(), task.Header.Name)
		fmt.Printf("════════════════════════════════════\n")

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
func (s BasicTaskEngine) update(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, tasks)
	if err != nil {
		return err
	}
	for _, task := range readyToUpdateStructs {

		fmt.Printf("────────────────────────────────────\n")
		fmt.Printf("📦 Processing: %s.%s\n", task.GetKey(), task.Header.Name)
		fmt.Printf("────────────────────────────────────\n")

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
func (s BasicTaskEngine) prepareToRemove(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct) ([]basic_task_payload_structs.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToRemoveStructs := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, task := range tasks {
		if err := s.completeInformation(ctx, &task); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, task)
		}
	}
	logrus.Debugf("'%d' task(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s BasicTaskEngine) remove(ctx *application.ApplicationContext, tasks []basic_task_payload_structs.TaskBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, tasks)
	if err != nil {
		return err
	}

	for _, task := range readyToRemoveStructs {

		fmt.Printf("════════════════════════════════════\n")
		fmt.Printf("📦 Processing: %s.%s\n", task.GetKey(), task.Header.Name)
		fmt.Printf("════════════════════════════════════\n")

		taskFlow := flowx.NewFlow("RemoveExistingTask").
			Step(&remove_steps.DeleteTask{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"task":      task,
		})

		if err := taskFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Task Remove işleminde hata oluştu: %w", &err)
		}

	}

	return nil
}

func (s BasicTaskEngine) execute(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	taskService := ioc.Get[contracts.TaskServiceInterface[basic_task_payload_structs.TaskBaseStruct]]()

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
		Name:  "Task.Common",
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
