package objectresource

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"
)

type ValidationRegexRule struct {
	ValidationRule
	Pattern string
}

func NewValidationRegexRule(name, pattern string, message Message) ValidationRegexRule {
	return ValidationRegexRule{
		ValidationRule: NewValidationRule("Regex", name, message),
		Pattern:        pattern,
	}
}

func (s *ValidationRegexRule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawData = make(map[string]interface{}, 0)

	if err := unmarshal(&rawData); err != nil {
		return err
	} else {
		if len(rawData) == 1 && rawData["Regex"] != nil {
			var tempObject struct {
				Regex string `yaml:"Regex"`
			}

			if err := unmarshal(&tempObject); err == nil {
				s.Pattern = tempObject.Regex
			} else {
				return &errors.InvalidFormatForPackageError{Value: tempObject.Regex}
			}

			s.ValidationRule = NewValidationRule("Regex", "", Message{})
		} else {
			var tempHeaderObject struct {
				ValidationRule
			}

			if err := unmarshal(&tempHeaderObject); err != nil {
				return err
			} else {

				s.ValidationRule = tempHeaderObject.ValidationRule
			}

			var tempObject struct {
				Pattern string `yaml:"Pattern"`
			}

			if err := unmarshal(&tempObject); err != nil {
				return err
			} else {
				s.Pattern = tempObject.Pattern
			}
		}

	}

	if utils.IsEmpty(s.Pattern) {
		return &errors.ErrFieldRequired{FieldName: "ValidationRegexRule.Pattern"}
	}

	return nil
}
