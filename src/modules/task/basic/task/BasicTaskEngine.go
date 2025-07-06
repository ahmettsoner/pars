package basic_task

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/pkg/utilities/json"
	commontask "parsdevkit.net/structs/task/basic-task"
	commontaskStruct "parsdevkit.net/structs/task/basic-task"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type BasicTaskEngine struct{}

func (s BasicTaskEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s BasicTaskEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	commontasks := make([]commontaskStruct.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		commontask, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected commontaskStruct.TaskBaseStruct, got %T", item)
		}

		if err := s.completeInformation(ctx, commontask); err != nil {
			return err
		}
		commontasks = append(commontasks, *commontask)
	}

	return s.createTasks(commontasks, true)
}

func (s BasicTaskEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	commontasks := make([]commontaskStruct.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		commontask, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in ProcDestroyess: expected commontaskStruct.TaskBaseStruct, got %T", item)
		}

		if err := s.completeInformation(ctx, commontask); err != nil {
			return err
		}
		commontasks = append(commontasks, *commontask)
	}

	return s.removeTasks(commontasks, true)
}

func (s BasicTaskEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Task.Common",
		Order: 5000,
	}
}

func (s BasicTaskEngine) createTasks(tasks []commontaskStruct.TaskBaseStruct, init bool) error {

	tasksReadyToCreate := make([]commontaskStruct.TaskBaseStruct, 0)
	tasksForUpdate := make([]commontaskStruct.TaskBaseStruct, 0)
	taskService := ioc.Get[contracts.TaskServiceInterface[commontask.TaskBaseStruct]]()

	for _, task := range tasks {
		if err := task.Validate(); err != nil {
			jsonObject, _ := json.ToJson(task)
			return fmt.Errorf("task invalid data: '%s'\n%w", jsonObject, err)
		}
	}

	for _, task := range tasks {
		ok, err := taskService.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Common task ('%s') kontrolünde hata oluştu\n%w", task.Header.Name, err)
		}
		if ok {
			newMommonlHash, err := encrypt.CalculateHashFromObject(task)
			if err != nil {
				return err
			}
			structHash, err := taskService.GetHash(task.Header.Name)
			if err != nil {
				return err
			}

			if newMommonlHash != structHash {
				tasksForUpdate = append(tasksForUpdate, task)
			}
		} else {
			tasksReadyToCreate = append(tasksReadyToCreate, task)
		}
	}
	logrus.Debugf("'%d' task(s) detected that will create", len(tasksReadyToCreate))
	logrus.Debugf("'%d' task(s) detected that will update", len(tasksForUpdate))

	logrus.Debugf("creating %v new tasks ", len(tasksReadyToCreate))
	logrus.Debugf("updating %v tasks ", len(tasksForUpdate))
	for _, task := range tasksReadyToCreate {

		if _, err := taskService.Save(task); err != nil {
			return err
		}

		if _, err := s.execute(task); err != nil {
			return err
		}

		fmt.Printf("%v Task created\n", task.Header.Name)
	}

	logrus.Debugf("updating %v tasks ", len(tasksForUpdate))
	for _, task := range tasksForUpdate {

		if _, err := taskService.Save(task); err != nil {
			return err
		}

		if _, err := s.execute(task); err != nil {
			return err
		}

		fmt.Printf("%v Task updated\n", task.Header.Name)
	}

	return nil
}
func (s BasicTaskEngine) removeTasks(tasks []commontaskStruct.TaskBaseStruct, permanent bool) error {

	taskService := ioc.Get[contracts.TaskServiceInterface[commontask.TaskBaseStruct]]()
	tasksReadyToDelete := make([]commontaskStruct.TaskBaseStruct, 0)
	for _, task := range tasks {
		ok, err := taskService.IsExists(task.Header.Name, task.Specifications.Workspace)
		if err != nil {
			return fmt.Errorf("xxx: Code task ('%s') kontrolünde hata oluştu\n%w", task.Header.Name, err)
		}
		if ok {
			tasksReadyToDelete = append(tasksReadyToDelete, task)
		}
	}

	for _, task := range tasksReadyToDelete {

		if _, err := taskService.Remove(task.Header.Name, task.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Task deleted\n", task.Header.Name)

	}

	return nil
}
func (s BasicTaskEngine) execute(model commontaskStruct.TaskBaseStruct) (*commontaskStruct.TaskBaseStruct, error) {

	taskService := ioc.Get[contracts.TaskServiceInterface[commontask.TaskBaseStruct]]()

	result, err := taskService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return result, nil
}

func (s BasicTaskEngine) completeInformation(ctx *application.ApplicationContext, model *commontaskStruct.TaskBaseStruct) error {

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

func (s BasicTaskEngine) getWorkspace(ctx *application.ApplicationContext, model commontaskStruct.TaskBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {

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
