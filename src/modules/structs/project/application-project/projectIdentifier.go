package applicationproject

import (
	"encoding/json"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ProjectIdentifier struct {
	ID        int
	Name      string
	Group     string
	Workspace string
}

func NewProjectIdentifier(id int, name string, group string, workspace string) ProjectIdentifier {
	return ProjectIdentifier{
		ID:        id,
		Name:      name,
		Group:     group,
		Workspace: workspace,
	}
}

func (s *ProjectIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name      string `yaml:"Name"`
				Group     string `yaml:"Group"`
				Workspace string `yaml:"Workspace"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
			s.Group = tempObject.Group
			s.Workspace = tempObject.Workspace
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}

func (s *ProjectIdentifier) UnmarshalJSON(data []byte) error {

	var tempObject struct {
		ID        int
		Name      string
		Group     string
		Workspace string
	}

	err := json.Unmarshal(data, &tempObject)
	if err != nil {
		return err
	}

	s.ID = tempObject.ID
	s.Name = tempObject.Name
	s.Group = tempObject.Group
	s.Workspace = tempObject.Workspace

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
