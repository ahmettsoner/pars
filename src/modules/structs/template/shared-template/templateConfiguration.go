package sharedtemplate

import (
	"gopkg.in/yaml.v3"
)

type TemplateConfiguration struct {
}

func NewTemplateConfiguration() TemplateConfiguration {
	return TemplateConfiguration{}
}

func (s *TemplateConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
	}

	return nil
}
