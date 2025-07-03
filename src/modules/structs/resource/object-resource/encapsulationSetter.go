package objectresource

import (
	"reflect"
	"strconv"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/structs"
	_string "parsdevkit.net/pkg/utilities/string"
)

type EncapsulationSetter struct {
	Name       string
	Visibility structs.VisibilityType
	Method     MethodIdentifier
	Available  bool
}

func NewEncapsulationSetter(name string, visibility structs.VisibilityType, method MethodIdentifier, available bool) EncapsulationSetter {
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
				Name       string                 `yaml:"Name"`
				Visibility structs.VisibilityType `yaml:"Visibility"`
				Method     MethodIdentifier       `yaml:"RefMethod"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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

	if (!_string.IsEmpty(s.Name) || !_string.IsEmpty(string(s.Visibility)) || !reflect.DeepEqual(MethodIdentifier{}, s.Method)) {
		s.Available = true
	}

	return nil
}
