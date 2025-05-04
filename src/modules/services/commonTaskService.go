package services

import (
	"encoding/json"
	"errors"

	"parsdevkit.net/structs/task"
	commontask "parsdevkit.net/structs/task/common-task"

	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type CommonTaskService struct {
	CommonTaskServiceInterface
	taskRespository              *repositories.TaskRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewCommonTaskService(environment string) *CommonTaskService {
	return &CommonTaskService{
		environment:                  environment,
		taskRespository:              repositories.NewTaskRepository(environment),
		generationHistoryRespository: repositories.NewGenerationHistoryRepository(environment),
	}
}

func (s CommonTaskService) GetByName(name string) (*commontask.TaskBaseStruct, error) {
	var task *commontask.TaskBaseStruct

	entity, err := s.taskRespository.GetByName(name)
	if err != nil {
		return nil, err
	}
	if entity != nil {
		err = json.Unmarshal([]byte(entity.Document), &task)
	} else {
		task = nil
	}

	return task, nil
}

func (s CommonTaskService) Save(model commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

	result, err := s.saveTaskInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CommonTaskService) List() (*([]commontask.TaskBaseStruct), error) {

	entityList, err := s.taskRespository.ListByKind(string(task.StructKinds.Common))
	if err != nil {
		return nil, err
	}

	taskList := make([]commontask.TaskBaseStruct, 0)

	for _, entity := range *entityList {
		var task commontask.TaskBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &task)

		taskList = append(taskList, task)
	}

	return &taskList, nil
}

func (s CommonTaskService) ListBySetAndLayers(set string, layers ...string) (*([]commontask.TaskBaseStruct), error) {

	entityList, err := s.taskRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	taskList := make([]commontask.TaskBaseStruct, 0)

	for _, entity := range *entityList {
		var task commontask.TaskBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &task)

		taskList = append(taskList, task)
	}

	return &taskList, nil
}

func (s CommonTaskService) Remove(name, workspace string, permanent bool) (*commontask.TaskBaseStruct, error) {
	//TODO: Geçici olarak tanımlandı, düzenlenecek

	taskTaskEntity, err := s.taskRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return nil, err
	}
	if taskTaskEntity == nil {
		return nil, errors.New("invalid task task")
	}

	logrus.Debugf("task %v deleting...", taskTaskEntity.Name)

	err = s.taskRespository.Delete(taskTaskEntity)
	var task commontask.TaskBaseStruct
	err = json.Unmarshal([]byte(taskTaskEntity.Document), &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s CommonTaskService) IsExists(name, workspace string) bool {

	taskTaskEntity, err := s.taskRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false
	}
	if taskTaskEntity == nil {
		return false
	}

	return true
}
func (s CommonTaskService) GetHash(name string) string {

	entity, err := s.taskRespository.GetByName(name)
	if err != nil {
		return ""
	}
	if entity == nil {
		return ""
	}

	return entity.Hash
}

func (s CommonTaskService) saveTaskInformation(taskMommonl commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

	jsonData, err := json.Marshal(taskMommonl)
	if err != nil {
		return nil, err
	}

	taskEntity := entities.Task{
		Name:     taskMommonl.Name,
		Document: string(jsonData),
	}

	err = s.taskRespository.Save(&taskEntity)
	if err != nil {
		return nil, err
	}

	return &taskMommonl, nil
}
