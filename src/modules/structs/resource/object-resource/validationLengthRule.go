package objectresource

import (
	"strconv"
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"
)

type ValidationLengthRule struct {
	ValidationRule
	Min int
	Max int
}

func NewValidationLengthRule(name string, min, max int, message Message) ValidationLengthRule {
	return ValidationLengthRule{
		ValidationRule: NewValidationRule("Length", name, message),
		Min:            min,
		Max:            max,
	}
}

func (s *ValidationLengthRule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawData = make(map[string]interface{}, 0)

	if err := unmarshal(&rawData); err != nil {
		return err
	} else {
		if len(rawData) == 1 && rawData["Length"] != nil {
			var tempObject struct {
				Length string `yaml:"Length"`
			}

			if err := unmarshal(&tempObject); err == nil {
				var parts []string = strings.Split(tempObject.Length, ":")
				if len(parts) == 1 {
					intVal1, err := strconv.Atoi(parts[0])
					if err != nil {
						return err
					}
					s.Min = intVal1
				} else if len(parts) == 2 {
					if !utils.IsEmpty(parts[0]) {
						intVal1, err := strconv.Atoi(parts[0])
						if err != nil {
							return err
						}
						s.Min = intVal1
					}

					if !utils.IsEmpty(parts[1]) {
						intVal2, err := strconv.Atoi(parts[1])
						if err != nil {
							return err
						}
						s.Max = intVal2
					}
				} else {
					return &errors.InvalidFormatForPackageError{Value: tempObject.Length}
				}

				s.ValidationRule = NewValidationRule("Length", "", Message{})
			}
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
				Min int `yaml:"Min"`
				Max int `yaml:"Max"`
			}

			if err := unmarshal(&tempObject); err != nil {
				return err
			} else {
				s.Min = tempObject.Min
				s.Max = tempObject.Max
			}
		}

	}
	return nil
}
