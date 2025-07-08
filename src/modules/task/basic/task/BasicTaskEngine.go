package basic_task

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/modules/task/basic_task_contract"
	commontask "parsdevkit.net/structs/task/basic-task"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/engines"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type BasicTaskEngine struct{}

func (s BasicTaskEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*commontask.TaskBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s BasicTaskEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Task.Common",
		Order: 5000,
	}
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
func (s BasicTaskEngine) prepareToCreate(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct) ([]commontask.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToCreateStructs := make([]commontask.TaskBaseStruct, 0)

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
func (s BasicTaskEngine) create(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct, init bool) error {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToCreateStructs, err := s.prepareToCreate(ctx, tasks)
	if err != nil {
		return err
	}

	for index, task := range readyToCreateStructs {

		logrus.Debugf("trying to create %v", task.Header.Name)
		if _, err := service.Save(task); err != nil {
			return err
		}

		if _, err := s.execute(task); err != nil {
			return err
		}
		fmt.Printf("%v (%d) Group created\n", task.Header.Name, index)

	}

	return nil
}

func (s BasicTaskEngine) prepareToUpdate(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct) ([]commontask.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToUpdateStructs := make([]commontask.TaskBaseStruct, 0)

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
func (s BasicTaskEngine) update(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct, init bool) error {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	readyToUpdateStructs, err := s.prepareToUpdate(ctx, tasks)
	if err != nil {
		return err
	}
	for _, task := range readyToUpdateStructs {
		if _, err := service.Save(task); err != nil {
			return err
		}
		if _, err := s.execute(task); err != nil {
			return err
		}
	}
	return nil
}
func (s BasicTaskEngine) prepareToRemove(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct) ([]commontask.TaskBaseStruct, error) {

	service := ioc.Get[basic_task_contract.TaskInterface]()
	readyToRemoveStructs := make([]commontask.TaskBaseStruct, 0)

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
func (s BasicTaskEngine) remove(ctx *application.ApplicationContext, tasks []commontask.TaskBaseStruct, permanent bool) error {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	readyToRemoveStructs, err := s.prepareToRemove(ctx, tasks)
	if err != nil {
		return err
	}

	for _, task := range readyToRemoveStructs {

		if _, err := service.Remove(task.Header.Name, task.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Group deleted\n", task.Header.Name)

	}

	logrus.Debugf("'%d' task(s) deleting", len(readyToRemoveStructs))

	return nil
}

func (s BasicTaskEngine) execute(model commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

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

func (s BasicTaskEngine) completeInformation(ctx *application.ApplicationContext, model *commontask.TaskBaseStruct) error {

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

func (s BasicTaskEngine) getWorkspace(ctx *application.ApplicationContext, model commontask.TaskBaseStruct) (*basic_workspace_payload.WorkspaceBaseStruct, error) {

	workspaceName := model.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *basic_workspace_payload.WorkspaceBaseStruct = nil

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

func CastArrayToConcrate(data []schemas.SchemaInterface) ([]commontask.TaskBaseStruct, error) {
	r := make([]commontask.TaskBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*commontask.TaskBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected commontask.TaskBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
