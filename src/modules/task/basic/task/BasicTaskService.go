package basic_task

import (
	"encoding/json"
	"errors"
	"fmt"

	"parsdevkit.net/application/contracts"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"

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

func NewBasicTaskService(environment string) contracts.TaskServiceInterface[basic_task_payload_structs.TaskBaseStruct] {
	taskRespository := ioc.Get[*repositories.TaskRepository]()
	generationHistoryRespository := ioc.Get[*repositories.GenerationHistoryRepository]()

	return &BasicTaskService{
		environment:                  environment,
		taskRespository:              taskRespository,
		generationHistoryRespository: generationHistoryRespository,
	}
}

func (s BasicTaskService) GetByName(name string) (*basic_task_payload_structs.TaskBaseStruct, error) {
	var task *basic_task_payload_structs.TaskBaseStruct

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

func (s BasicTaskService) Save(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	result, err := s.saveTaskInformation(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s BasicTaskService) List() (*([]basic_task_payload_structs.TaskBaseStruct), error) {

	entityList, err := s.taskRespository.ListByKind(basic_task_payload_structs.TASK_KIND)
	if err != nil {
		return nil, err
	}

	taskList := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, entity := range *entityList {
		var task basic_task_payload_structs.TaskBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &task)

		taskList = append(taskList, task)
	}

	return &taskList, nil
}

func (s BasicTaskService) ListBySetAndLayers(set string, layers ...string) (*([]basic_task_payload_structs.TaskBaseStruct), error) {

	entityList, err := s.taskRespository.ListBySetAndLayers(set, layers...)
	if err != nil {
		return nil, err
	}

	taskList := make([]basic_task_payload_structs.TaskBaseStruct, 0)

	for _, entity := range *entityList {
		var task basic_task_payload_structs.TaskBaseStruct
		err = json.Unmarshal([]byte(entity.Document), &task)

		taskList = append(taskList, task)
	}

	return &taskList, nil
}

func (s BasicTaskService) Remove(name, workspace string, permanent bool) (*basic_task_payload_structs.TaskBaseStruct, error) {
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
	var task basic_task_payload_structs.TaskBaseStruct
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

func (s BasicTaskService) saveTaskInformation(taskMommonl basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

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

func (s BasicTaskService) SaveTask(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	jsonData, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	taskEntity := entities.Task{
		Name:     model.Header.Name,
		Document: string(jsonData),
	}

	err = s.taskRespository.Save(&taskEntity)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *BasicTaskService) UndoSaveTask(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	err := s.taskRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Task silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}

	return &model, nil
}

func (s *BasicTaskService) DeleteTask(model basic_task_payload_structs.TaskBaseStruct) (*basic_task_payload_structs.TaskBaseStruct, error) {

	logrus.Debugf("template %v removing", model.Header.Name)

	err := s.taskRespository.DeleteByName(model.Header.Name)
	if err != nil {
		return nil, fmt.Errorf("xxx: Code Task silme aşamasında beklenmeyen hata oluştu %s\n%w", model.Header.Name, err)
	}
	logrus.Debugf("template (%v) information removed", model)

	return &model, nil
}
