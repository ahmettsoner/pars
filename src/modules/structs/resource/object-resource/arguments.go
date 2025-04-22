package objectresource

import (
	"strings"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

// Burda argümanların MethodIdentifier üzerinden çözümlenmesi de sağlanmalı, farklı sınıflarda tanımlı mesaj imzaları çözümlenmesine ihtiyaç var
type Arguments struct {
	Arguments []MethodArgument
	Reference MethodIdentifier
}

func NewArgument(arguments []MethodArgument, reference MethodIdentifier) Arguments {
	return Arguments{
		Arguments: arguments,
		Reference: reference,
	}
}

func (s *Arguments) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Arguments []MethodArgument `yaml:"Arguments"`
				Reference MethodIdentifier `yaml:"Reference"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Arguments = tempObject.Arguments
				s.Reference = tempObject.Reference
			}

		} else {
			return err
		}
	} else {
		var parts []string = strings.Split(strings.ReplaceAll(value, ", ", ","), ",")
		for _, part := range parts {
			var arguments []string = strings.Split(part, " ")

			if len(arguments) == 1 {
				s.Arguments = append(s.Arguments, NewMethodArgument("", arguments[0]))
			} else if len(arguments) == 2 {
				s.Arguments = append(s.Arguments, NewMethodArgument(arguments[0], arguments[1]))
			} else {
				return &errors.InvalidFormatForLanguageError{Value: part}
			}
		}
	}

	return nil
}
