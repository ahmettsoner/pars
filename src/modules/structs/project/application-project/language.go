package applicationproject

import (
	"fmt"
	"strings"

	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
)

type Language struct {
	Type    models.LanguageType
	Version string
}

func NewLanguage(_type models.LanguageType, version string) Language {
	return Language{
		Type:    _type,
		Version: version,
	}
}

func NewLanguage_LanguageOnly(_type models.LanguageType) Language {
	return Language{
		Type: _type,
	}
}
func (s Language) Validate() error {
	if _string.IsEmpty(string(s.Type)) || s.Type.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}
	return nil
}

func (s *Language) GetFullName() string {
	fullName := s.Type.String()
	if !_string.IsEmpty(s.Version) {
		fullName = fmt.Sprintf("%v@%v", s.Type.String(), s.Version)
	}

	return fullName
}

func (s *Language) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type    models.LanguageType `yaml:"Type"`
				Version string              `yaml:"Version"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err
			} else {
				s.Type = tempObject.Type
				s.Version = tempObject.Version
			}

		} else {
			return err
		}

	} else {
		var parts []string = strings.Split(value, "@")
		if len(parts) == 1 {
			// No specific version provided so assume latest

			enum, err := models.LanguageTypeEnumFromString(value)
			if err != nil {
				return err
			}
			s.Type = enum
		} else if len(parts) == 2 {
			languageName := strings.TrimSpace(strings.ToLower(parts[0]))
			languageVersion := strings.TrimSpace(parts[1])

			enum, err := models.LanguageTypeEnumFromString(languageName)
			if err != nil {
				return err
			}
			s.Type = enum

			if _string.IsEmpty(string(s.Type)) {
				return &errors.InvalidLanguageError{Value: languageName}
			}
			s.Version = languageVersion
		} else {
			return &errors.InvalidFormatForLanguageError{Value: value}
		}
	}

	return nil
}
