package project

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
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
func (e ProjectIdentifier) Validate() error {
	if utils.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
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

	return nil
}
