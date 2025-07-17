package template

import (
	"parsdevkit.net/application/models/label"

	"gopkg.in/yaml.v3"
	sectionPkg "parsdevkit.net/application/models/section"
)

type Project struct {
	Name    string
	Labels  []label.Label
	Section sectionPkg.Section
}

func NewProject(name string, labels []label.Label, section sectionPkg.Section) Project {
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
				Labels  []label.Label      `yaml:"Labels"`
				Section sectionPkg.Section `yaml:"Section"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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
