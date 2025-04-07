package option

import (
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Option struct {
	Key   string
	Value interface{}
}

func NewOption(key string, value interface{}) Option {
	return Option{
		Key:   key,
		Value: value,
	}
}

func (s *Option) IsKeyExists() bool {
	return !utils.IsEmpty(s.Key)
}

func (s *Option) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Key   string      `yaml:"Key"`
				Value interface{} `yaml:"Value"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Key = tempObject.Key
				s.Value = tempObject.Value
			}

		} else {
			return err
		}
	} else {
		var parts []string = strings.Split(value, "=")
		if len(parts) == 1 {
			s.Key = value
		} else if len(parts) == 2 {
			key := strings.TrimSpace(strings.ToLower(parts[0]))
			_value := strings.TrimSpace(parts[1])

			s.Key = key

			if utils.IsEmpty(s.Key) {
				return &errors.InvalidLanguageError{Value: key}
			}
			s.Value = _value
		} else {
			return &errors.InvalidFormatForLanguageError{Value: value}
		}
	}

	if utils.IsEmpty(s.Key) {
		return &errors.ErrFieldRequired{FieldName: "Option.Key"}
	}

	return nil
}
