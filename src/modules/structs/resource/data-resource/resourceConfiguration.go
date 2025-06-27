package dataresource

import (
	"parsdevkit.net/core/utils"
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

	if utils.IsEmpty(string(s.Generate)) {
		s.Generate = ChangeTrackers.OnChange
	}

	return nil
}
