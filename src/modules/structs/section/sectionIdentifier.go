package section

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type SectionIdentifier struct {
	Name string
}

func NewSectionIdentifier(name string) SectionIdentifier {
	return SectionIdentifier{
		Name: name,
	}
}

func (s *SectionIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
