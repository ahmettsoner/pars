package template

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
)

type TemplateIdentifier struct {
	ID        int
	Name      string
	Workspace string
}

func NewTemplateIdentifier(id int, name string, workspace string) TemplateIdentifier {
	return TemplateIdentifier{
		ID:        id,
		Name:      name,
		Workspace: workspace,
	}
}
func (e TemplateIdentifier) Validate() error {
	if utils.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *TemplateIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
