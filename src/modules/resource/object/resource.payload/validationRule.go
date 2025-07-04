package object_resource_payload

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
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
func (e ValidationRule) Validate() error {
	if _string.IsEmpty(e.Type) {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}
	return nil
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
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

			} else {
				s.Type = tempObject.Type
				s.Name = tempObject.Name
				s.Message = tempObject.Message
			}

		}
	} else {
		s.Type = value
	}

	return nil
}
