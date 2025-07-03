package dataresource

import (
	_string "parsdevkit.net/pkg/utilities/string"
)

type ResourceConfiguration struct {
	Generate ChangeTracker
}

func NewResourceConfiguration(generate ChangeTracker) ResourceConfiguration {
	return ResourceConfiguration{
		Generate: generate,
	}
}

func (s *ResourceConfiguration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Generate ChangeTracker `yaml:"Generate"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Generate = tempObject.Generate

	}

	if _string.IsEmpty(string(s.Generate)) {
		s.Generate = ChangeTrackers.OnChange
	}

	return nil
}
