package objectresource

import (
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type MethodArgument struct {
	Name  string
	Value string
}

func NewMethodArgument(name, value string) MethodArgument {
	return MethodArgument{
		Name:  name,
		Value: value,
	}
}

func (s *MethodArgument) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name  string `yaml:"Name"`
				Value string `yaml:"Value"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Value = tempObject.Value
			}

		} else {
			return err
		}
	} else {
		var parts []string = strings.Split(value, " ")
		if len(parts) == 1 {
			s.Value = value
		} else if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			_value := strings.TrimSpace(parts[1])

			s.Name = name

			if utils.IsEmpty(s.Name) {
				return &errors.InvalidLanguageError{Value: name}
			}
			s.Value = _value
		} else {
			return &errors.InvalidFormatForLanguageError{Value: value}
		}
	}

	if utils.IsEmpty(s.Value) {
		return &errors.ErrFieldRequired{FieldName: "MethodArgument.Value"}
	}

	return nil
}
