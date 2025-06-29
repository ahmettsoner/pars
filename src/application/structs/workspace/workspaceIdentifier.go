package workspace

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
)

type WorkspaceIdentifier struct {
	ID   int
	Name string
}

func NewWorkspaceIdentifier(id int, name string) WorkspaceIdentifier {
	return WorkspaceIdentifier{
		ID:   id,
		Name: name,
	}
}
func (e WorkspaceIdentifier) Validate() error {
	if utils.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *WorkspaceIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
