package structs

import (
	templateStruct "parsdevkit.net/application/structs/template"
)

type TemplateConfiguration struct {
	Selectors templateStruct.Selectors
}

func NewTemplateConfiguration(selectors templateStruct.Selectors) TemplateConfiguration {
	return TemplateConfiguration{
		Selectors: selectors,
	}
}

func (s *TemplateConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Selectors templateStruct.Selectors `yaml:"Selectors"`
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
