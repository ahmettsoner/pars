package template

import (
	"parsdevkit.net/structs/label"

	"gopkg.in/yaml.v3"
)

type Project struct {
	Name    string
	Labels  []label.Label
	Section Section
}

func NewProject(name string, labels []label.Label, section Section) Project {
	return Project{
		Name:    name,
		Labels:  labels,
		Section: section,
	}
}

func (s *Project) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name    string
				Labels  []label.Label `yaml:"Labels"`
				Section Section       `yaml:"Section"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Labels = tempObject.Labels
				s.Section = tempObject.Section
			}

		} else {
			return err
		}

	} else {
		s.Labels = []label.Label{label.NewLabel_KeyOnly(value)}
	}

	return nil
}
