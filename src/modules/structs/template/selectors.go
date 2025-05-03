package template

import (
	"gopkg.in/yaml.v3"
)

type Selectors struct {
	Project  Project
	Resource Resource
}

func NewSelectors(project Project, resource Resource) Selectors {
	return Selectors{
		Project:  project,
		Resource: resource,
	}
}

func (s *Selectors) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempObject struct {
		Resource Resource `yaml:"Resource"`
		Project  Project  `yaml:"Project"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Resource = tempObject.Resource
		s.Project = tempObject.Project
	}

	return nil
}
