package basic_task

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/contracts"
	commontask "parsdevkit.net/structs/task/basic-task"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/persistence/repositories"

	"parsdevkit.net/persistence/entities"

	"github.com/sirupsen/logrus"
)

type BasicTaskService struct {
	taskRespository              *repositories.TaskRepository
	generationHistoryRespository *repositories.GenerationHistoryRepository
	environment                  string
}

func NewBasicTaskService(environment string) contracts.TaskServiceInterface[commontask.TaskBaseStruct] {
	taskRespository := ioc.Get[*repositories.TaskRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &BasicTaskService{
		environment:                  environment,
		taskRespository:              taskRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s BasicTaskService) GetByName(name string) (*commontask.TaskBaseStruct, error) {
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

func (s BasicTaskService) Save(model commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

	result, err := s.saveTaskInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s BasicTaskService) List() (*([]commontask.TaskBaseStruct), error) {

	entityList, err := s.taskRespository.ListByKind(commontask.TASK_KIND)
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

func (s BasicTaskService) ListBySetAndLayers(set string, layers ...string) (*([]commontask.TaskBaseStruct), error) {

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

func (s BasicTaskService) Remove(name, workspace string, permanent bool) (*commontask.TaskBaseStruct, error) {
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

func (s BasicTaskService) IsExists(name, workspace string) (bool, error) {

	taskTaskEntity, err := s.taskRespository.GetByNameAndWorkspace(name, workspace)
	if err != nil {
		return false, fmt.Errorf("xxx: Task getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if taskTaskEntity == nil {
		return false, nil
	}

	return true, nil
}
func (s BasicTaskService) GetHash(name string) (string, error) {

	entity, err := s.taskRespository.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("xxx: Task getirme aşamasında beklenmeyen hata oluştu '%s'\n%w", name, err)
	}
	if entity == nil {
		return "", fmt.Errorf("xxx: Task tanımlı değil '%s' Hash bilgisi alınamıyor\n%w", name, err)
	}

	return entity.Hash, nil
}

func (s BasicTaskService) saveTaskInformation(taskMommonl commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

	jsonData, err := json.Marshal(taskMommonl)
	if err != nil {
		return nil, err
	}

	taskEntity := entities.Task{
		Name:     taskMommonl.Header.Name,
		Document: string(jsonData),
	}

	err = s.taskRespository.Save(&taskEntity)
	if err != nil {
		return nil, err
	}

	return &taskMommonl, nil
}
