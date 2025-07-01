package codetemplate

import (
	_string "parsdevkit.net/core/utilities/string"
	"parsdevkit.net/structs/template"
)

type TemplateConfiguration struct {
	Generate  ChangeTracker
	Selectors template.Selectors
}

func NewTemplateConfiguration(generate ChangeTracker, selectors template.Selectors) TemplateConfiguration {
	return TemplateConfiguration{
		Generate:  generate,
		Selectors: selectors,
	}
}

func (s *TemplateConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Generate  ChangeTracker      `yaml:"Generate"`
		Selectors template.Selectors `yaml:"Selectors"`
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
