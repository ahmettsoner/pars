package structs

import (
	templateStruct "parsdevkit.net/application/structs/template"
	_string "parsdevkit.net/pkg/utilities/string"
)

type TemplateConfiguration struct {
	Generate  ChangeTracker
	Selectors templateStruct.Selectors
}

func NewTemplateConfiguration(generate ChangeTracker, selectors templateStruct.Selectors) TemplateConfiguration {
	return TemplateConfiguration{
		Generate:  generate,
		Selectors: selectors,
	}
}

func (s *TemplateConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Generate  ChangeTracker            `yaml:"Generate"`
		Selectors templateStruct.Selectors `yaml:"Selectors"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Generate = tempObject.Generate
		s.Selectors = tempObject.Selectors

	}

	if _string.IsEmpty(string(s.Generate)) {
		s.Generate = ChangeTrackers.OnChange
	}

	return nil
}
