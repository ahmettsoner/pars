package template

import (
	"parsdevkit.net/structs/class"
	"parsdevkit.net/structs/section"

	"gopkg.in/yaml.v3"
)

type Section struct {
	Name    string
	Classes []class.Class
}

func NewSection(name string, classes []class.Class, sections []section.Section) Section {
	return Section{
		Name:    name,
		Classes: classes,
	}
}

func (s *Section) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name    string
				Classes []class.Class `yaml:"Classes"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Classes = tempObject.Classes
			}

		} else {
			return err
		}

	} else {
		s.Classes = []class.Class{class.NewClass_KeyOnly(value)}
	}

	return nil
}
