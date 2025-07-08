package structs

import (
	taskStruct "parsdevkit.net/application/structs/template"
)

type TaskConfiguration struct {
	Selectors taskStruct.Selectors
}

func NewTaskConfiguration(selectors taskStruct.Selectors) TaskConfiguration {
	return TaskConfiguration{
		Selectors: selectors,
	}
}

func (s *TaskConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Selectors taskStruct.Selectors `yaml:"Selectors"`
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
