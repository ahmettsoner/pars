package basic_task

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/pkg/utilities/json"
	commontask "parsdevkit.net/structs/task/common-task"
	commontaskStruct "parsdevkit.net/structs/task/common-task"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/ioc"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type CommonTaskEngine struct{}

func (s CommonTaskEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s CommonTaskEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	commontasks := make([]commontaskStruct.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		commontask, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected commontaskStruct.TaskBaseStruct, got %T", item)
		}

		commontasks = append(commontasks, *commontask)
	}

	return s.createTasks(commontasks, true)
}

func (s CommonTaskEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	commontasks := make([]commontaskStruct.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		commontask, ok := item.(*commontaskStruct.TaskBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in ProcDestroyess: expected commontaskStruct.TaskBaseStruct, got %T", item)
		}

		commontasks = append(commontasks, *commontask)
	}

	return s.removeTasks(commontasks, true)
}

func (s CommonTaskEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Task.Common",
		Order: 5000,
	}
}

func (s CommonTaskEngine) createTasks(tasks []commontaskStruct.TaskBaseStruct, init bool) error {

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
func (s CommonTaskEngine) removeTasks(tasks []commontaskStruct.TaskBaseStruct, permanent bool) error {

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
func (s CommonTaskEngine) execute(model commontaskStruct.TaskBaseStruct) (*commontaskStruct.TaskBaseStruct, error) {

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
