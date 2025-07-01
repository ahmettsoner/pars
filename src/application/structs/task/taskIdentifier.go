package task

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	_string "parsdevkit.net/core/utilities/string"
)

type TaskIdentifier struct {
	ID        int
	Name      string
	Workspace string
}

func NewTaskIdentifier(id int, name string, workspace string) TaskIdentifier {
	return TaskIdentifier{
		ID:        id,
		Name:      name,
		Workspace: workspace,
	}
}

func (e TaskIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *TaskIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name      string `yaml:"Name"`
				Workspace string `yaml:"Workspace"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
			s.Workspace = tempObject.Workspace
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
