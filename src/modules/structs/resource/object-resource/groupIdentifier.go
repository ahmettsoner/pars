package objectresource

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type GroupIdentifier struct {
	Name string
}

func NewGroupIdentifier(name string) GroupIdentifier {
	return GroupIdentifier{
		Name: name,
	}
}

func (s *GroupIdentifier) IsNameExists() bool {
	return !utils.IsEmpty(s.Name)
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

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Group.Name"}
	}

	return nil
}
