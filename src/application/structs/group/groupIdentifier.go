package group

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
)

type GroupIdentifier struct {
	ID   int
	Name string
}

func NewGroupIdentifier(id int, name string) GroupIdentifier {
	return GroupIdentifier{
		ID:   id,
		Name: name,
	}
}
func (e GroupIdentifier) Validate() error {
	if utils.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
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
				return fmt.Errorf("xxx: Group Identifier Çözümlenemedi\n%w", err)
			}

			s.Name = tempObject.Name
		} else {
			return fmt.Errorf("xxx: Group Identifier dönüştürme hatası oluştu\n%w", err)
		}
	} else {
		s.Name = value
	}

	return nil
}
