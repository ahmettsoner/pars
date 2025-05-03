package applicationproject

import (
	"gopkg.in/yaml.v3"
)

type Schema struct {
}

func NewSchema() Schema {
	return Schema{}
}

func (s *Schema) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
