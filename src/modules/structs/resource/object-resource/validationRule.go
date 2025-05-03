package objectresource

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ValidationRuleInterface interface {
}

type ValidationRule struct {
	ValidationRuleInterface
	Type    string
	Name    string
	Message Message
}

func NewValidationRule(_type, name string, message Message) ValidationRule {
	return ValidationRule{
		Type:    _type,
		Name:    name,
		Message: message,
	}
}

func (s *ValidationRule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type    string  `yaml:"Type"`
				Name    string  `yaml:"Name"`
				Message Message `yaml:"Message"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Type = tempObject.Type
				s.Name = tempObject.Name
				s.Message = tempObject.Message
			}

		}
	} else {
		s.Type = value
	}

	if utils.IsEmpty(s.Type) {
		return &errors.ErrFieldRequired{FieldName: "ValidationRule.Type"}
	}

	return nil
}
