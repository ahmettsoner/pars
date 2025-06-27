package objectresource

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/yaml.v3"
)

type MethodIdentifier struct {
	Name string
}

func NewMethodIdentifier(name string) MethodIdentifier {
	return MethodIdentifier{
		Name: name,
	}
}
func (e MethodIdentifier) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Name, v.Required),
	)
}

func (s *MethodIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
