package sharedtemplate

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type TemplateIdentifier struct {
	Name      string
	Workspace string
}

func NewTemplateIdentifier(name string, workspace string) TemplateIdentifier {
	return TemplateIdentifier{
		Name:      name,
		Workspace: workspace,
	}
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

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
