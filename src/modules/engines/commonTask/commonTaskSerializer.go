package commonTask

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/task"
	commontask "parsdevkit.net/structs/task/common-task"
	"parsdevkit.net/structs/workspace"

	"gopkg.in/yaml.v3"
)

type CommonTaskSerializer struct{}

func (s CommonTaskSerializer) GetTaskStructsFromString(data string) ([]commontask.TaskBaseStruct, error) {
	tasks := make([]commontask.TaskBaseStruct, 0)

	yamlLines := strings.Split(string(data), "---")

	for _, line := range yamlLines {
		var header structs.Header
		if err := yaml.Unmarshal([]byte(line), &header); err != nil {
			return nil, err
		}

		if header.Type == structs.StructTypes.Task {
			var resourceHeader task.Header
			if err := yaml.Unmarshal([]byte(line), &resourceHeader); err != nil {
				return nil, err
			}
			if resourceHeader.Kind == task.StructKinds.Common {
				var taskDefinitionStruct = commontask.TaskBaseStruct{}
				if err := yaml.Unmarshal([]byte(line), &taskDefinitionStruct); err != nil {
					return nil, err
				}

				rawTask := taskDefinitionStruct
				if err := s.CompleteTaskInformation(&rawTask); err != nil {
					return nil, err
				}

				tasks = append(tasks, rawTask)
			}
		}
	}

	return tasks, nil
}

func (s CommonTaskSerializer) GetTaskStructsFromFile(files ...string) ([]commontask.TaskBaseStruct, error) {
	tasks := make([]commontask.TaskBaseStruct, 0)
	for _, file := range files {

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		taskStructs, err := s.GetTaskStructsFromString(string(data))
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, taskStructs...)

	}
	return tasks, nil
}

func (s CommonTaskSerializer) CompleteTaskInformation(task *commontask.TaskBaseStruct) error {

	logrus.Debugf("filling task (%v) information", task.Name)

	if strings.ToLower(task.Specifications.Workspace) != strings.ToLower("None") {
		activeWorkspace, err := s.GetWorkspace(*task)
		if err != nil {
			return err
		}

		//WARN: Doğru mu oldu?
		task.Specifications.Workspace = activeWorkspace.Name
		task.Specifications.WorkspaceObject = activeWorkspace.Specifications
		logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Name, task.Name)
	}

	return nil
}

func (s CommonTaskSerializer) GetWorkspace(task commontask.TaskBaseStruct) (*workspace.WorkspaceBaseStruct, error) {

	appContext := GetContext()
	var result workspace.WorkspaceBaseStruct = *appContext.CurrentWorkspace

	workspaceName := task.Specifications.Workspace

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = *workspace
	}

	return &result, nil
}
