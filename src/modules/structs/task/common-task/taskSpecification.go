package commontask

import (
	applicationTask "parsdevkit.net/application/structs/task"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/errors"
	_string "parsdevkit.net/core/utilities/string"
	actionBase "parsdevkit.net/structs/task/actions"
)

type TaskSpecification struct {
	applicationTask.TaskIdentifier
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	Trigger         Trigger
	Retry           Retry
	Timeout         int
	Concurrency     int
	Parameters      map[string]interface{}
	Tasks           []actionBase.ActionInterface
	// parameters:
	// - name: parameter1
	//   type: string
	//   default: "default_value"
	// - name: parameter2
	//   type: object
	//   default:
	//     product:
	//       name: no-name
	//       id: 1e001d0e-f934-55b5-b8fc-5a2eb462c684
	//       category: 3f77edcb-db03-53f0-a15c-d42801e0b51e
}

func NewTaskSpecification(id int, name, workspace string, trigger Trigger, retry Retry, timeout int, concurrency int, parameters map[string]interface{}, tasks []actionBase.ActionInterface, workspaceObject applicationWorkspace.WorkspaceIdentifier) TaskSpecification {
	return TaskSpecification{
		TaskIdentifier:  applicationTask.NewTaskIdentifier(id, name, workspace),
		WorkspaceObject: workspaceObject,
		Trigger:         trigger,
		Retry:           retry,
		Timeout:         timeout,
		Concurrency:     concurrency,
		Parameters:      parameters,
		Tasks:           tasks,
	}
}

func (e TaskSpecification) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *TaskSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		applicationTask.TaskIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.TaskIdentifier = tempIdentifierObject.TaskIdentifier
	}

	var tempObject struct {
		Trigger     Trigger                `yaml:"Trigger"`
		Retry       Retry                  `yaml:"Retry"`
		Timeout     int                    `yaml:"Timeout"`
		Concurrency int                    `yaml:"Concurrency"`
		Parameters  map[string]interface{} `yaml:"Parameters"`
		Tasks       []actionBase.Action    `yaml:"Tasks"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Trigger = tempObject.Trigger
		s.Retry = tempObject.Retry
		s.Timeout = tempObject.Timeout
		s.Concurrency = tempObject.Concurrency
		s.Parameters = tempObject.Parameters
	}

	return nil
}
