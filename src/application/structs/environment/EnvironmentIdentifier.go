package environment

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type EnvironmentIdentifier struct {
	ID   int
	Name string
}

func NewEnvironmentIdentifier(id int, name string) EnvironmentIdentifier {
	return EnvironmentIdentifier{
		ID:   id,
		Name: name,
	}
}
func (e EnvironmentIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *EnvironmentIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
