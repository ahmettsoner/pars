package structs

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type GroupIdentifier struct {
	Name string
}

func NewGroupIdentifier(name string) GroupIdentifier {
	return GroupIdentifier{
		Name: name,
	}
}
func (e GroupIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *GroupIdentifier) IsNameExists() bool {
	return !_string.IsEmpty(s.Name)
}

func (s *GroupIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
