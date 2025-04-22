package commontask

import (
	actionBase "parsdevkit.net/structs/task/actions"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type TaskSpecification struct {
	TaskIdentifier
	WorkspaceObject workspace.WorkspaceSpecification
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

func NewTaskSpecification(id int, name, workspace string, trigger Trigger, retry Retry, timeout int, concurrency int, parameters map[string]interface{}, tasks []actionBase.ActionInterface, workspaceObject workspace.WorkspaceSpecification) TaskSpecification {
	return TaskSpecification{
		TaskIdentifier:  NewTaskIdentifier(id, name, workspace),
		WorkspaceObject: workspaceObject,
		Trigger:         trigger,
		Retry:           retry,
		Timeout:         timeout,
		Concurrency:     concurrency,
		Parameters:      parameters,
		Tasks:           tasks,
	}
}

func (s *TaskSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		TaskIdentifier
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
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Trigger = tempObject.Trigger
		s.Retry = tempObject.Retry
		s.Timeout = tempObject.Timeout
		s.Concurrency = tempObject.Concurrency
		s.Parameters = tempObject.Parameters
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
