package basictask

import (
	"parsdevkit.net/structs/task"
)

type TaskConfiguration struct {
	Selectors task.Selectors
}

func NewTaskConfiguration(selectors task.Selectors) TaskConfiguration {
	return TaskConfiguration{
		Selectors: selectors,
	}
}

func (s *TaskConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Selectors task.Selectors `yaml:"Selectors"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Selectors = tempObject.Selectors

	}

	return nil
}
