package group

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/yaml.v3"
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
	return v.ValidateStruct(&e,
		v.Field(&e.Name, v.Required),
	)
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
				return fmt.Errorf("xxx: Group Identifier Çözümlenemedi %w", err)
			}

			s.Name = tempObject.Name
		} else {
			return fmt.Errorf("xxx: Group Identifier dönüştürme hatası oluştu %w", err)
		}
	} else {
		s.Name = value
	}

	return nil
}
