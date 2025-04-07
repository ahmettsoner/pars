package commonTask

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	textTemplate "text/template"

	commontask "parsdevkit.net/structs/task/common-task"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type CommonTaskEngine struct{}

func (s CommonTaskEngine) CreateTasksFromTask(init bool, data any, taskFiles ...string) error {

	var allTasks []commontask.TaskBaseStruct = make([]commontask.TaskBaseStruct, 0)

	for _, taskFilePath := range taskFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), taskFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = textTemplate.Must(textTemplate.New("TaskFromTask").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		groupSerializer := CommonTaskSerializer{}
		tasks, err := groupSerializer.GetTaskStructsFromString(mainStr)
		if err != nil {
			return err
		}
		allTasks = append(allTasks, tasks...)
	}

	if err := s.CreateTasks(allTasks, init); err != nil {
		return err
	}

	return nil
}
func (s CommonTaskEngine) CreateTasksFromFile(init bool, files ...string) error {
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {
				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		logrus.Debugf("found %v files", len(allFiles))
		groupSerializer := CommonTaskSerializer{}
		tasksFromFile, err := groupSerializer.GetTaskStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v task", len(tasksFromFile))
		if err := s.CreateTasks(tasksFromFile, init); err != nil {
			return err
		}
	}
	return nil
}

func (s CommonTaskEngine) RemoveTasksFromFile(permanent bool, files ...string) error {
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		groupSerializer := CommonTaskSerializer{}
		tasksFromFile, err := groupSerializer.GetTaskStructsFromFile(allFiles...)
		if err != nil {
			return err
		}

		if err := s.RemoveTasks(tasksFromFile, permanent); err != nil {
			return err
		}
	}
	return nil
}
func (s CommonTaskEngine) CreateTasks(tasks []commontask.TaskBaseStruct, init bool) error {

	tasksReadyToCreate := make([]commontask.TaskBaseStruct, 0)
	tasksForUpdate := make([]commontask.TaskBaseStruct, 0)
	taskService := services.NewCommonTaskService(utils.GetEnvironment())

	for _, task := range tasks {
		if ok := taskService.IsExists(task.Name, task.Specifications.Workspace); ok {
			newMommonlHash, err := utils.CalculateHashFromObject(task)
			if err != nil {
				return err
			}
			structHash := taskService.GetHash(task.Name)

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

		if _, err := s.Execute(task); err != nil {
			return err
		}

		fmt.Printf("%v Task created\n", task.Name)
	}

	logrus.Debugf("updating %v tasks ", len(tasksForUpdate))
	for _, task := range tasksForUpdate {

		if _, err := taskService.Save(task); err != nil {
			return err
		}

		if _, err := s.Execute(task); err != nil {
			return err
		}

		fmt.Printf("%v Task updated\n", task.Name)
	}

	return nil
}

func (s CommonTaskEngine) RemoveTasks(tasks []commontask.TaskBaseStruct, permanent bool) error {

	taskService := services.NewCommonTaskService(utils.GetEnvironment())
	tasksReadyToDelete := make([]commontask.TaskBaseStruct, 0)
	for _, task := range tasks {
		if ok := taskService.IsExists(task.Name, task.Specifications.Workspace); ok {
			tasksReadyToDelete = append(tasksReadyToDelete, task)
		}
	}

	for _, task := range tasksReadyToDelete {

		if _, err := taskService.Remove(task.Name, task.Specifications.Workspace, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Task deleted\n", task.Name)

	}

	return nil
}

func (s CommonTaskEngine) Execute(model commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error) {

	taskService := services.NewCommonTaskService(utils.GetEnvironment())

	result, err := taskService.GetByName(model.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	taskEngine := NewCommonTaskOperations(utils.GetEnvironment())
	err = taskEngine.GenerateByTask(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
