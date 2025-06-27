package objectresource

import (
	"strings"

	v "github.com/go-ozzo/ozzo-validation/v4"
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
func (e MethodArgument) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Name, v.Required),
		v.Field(&e.Value, v.Required),
	)
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
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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

			s.Value = _value
		} else {
			return &errors.InvalidFormatForLanguageError{Value: value}
		}
	}

	return nil
}
