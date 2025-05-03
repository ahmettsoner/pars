package objectresource

import (
	"reflect"
	"strconv"

	"parsdevkit.net/core/utils"

	"gopkg.in/yaml.v3"
)

type EncapsulationSetter struct {
	Name       string
	Visibility VisibilityType
	Method     MethodIdentifier
	Available  bool
}

func NewEncapsulationSetter(name string, visibility VisibilityType, method MethodIdentifier, available bool) EncapsulationSetter {
	return EncapsulationSetter{
		Name:       name,
		Visibility: visibility,
		Method:     method,
		Available:  available,
	}
}

func (s *EncapsulationSetter) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name       string           `yaml:"Name"`
				Visibility VisibilityType   `yaml:"Visibility"`
				Method     MethodIdentifier `yaml:"RefMethod"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Visibility = tempObject.Visibility
				s.Method = tempObject.Method
			}
		}
	} else {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			s.Method = NewMethodIdentifier(value)
		} else {
			s.Available = boolValue
		}
	}

	if (!utils.IsEmpty(s.Name) || !utils.IsEmpty(string(s.Visibility)) || !reflect.DeepEqual(MethodIdentifier{}, s.Method)) {
		s.Available = true
	}

	return nil
}
