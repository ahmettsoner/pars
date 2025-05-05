package applicationproject

import (
	"fmt"
	"strings"

	"parsdevkit.net/models"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Platform struct {
	Type    models.PlatformType
	Version string
}

func NewPlatform(_type models.PlatformType, version string) Platform {
	return Platform{
		Type:    _type,
		Version: version,
	}
}

func NewPlatform_Basic(_type models.PlatformType) Platform {
	return Platform{
		Type: _type,
	}
}

func (s *Platform) GetFullName() string {
	fullName := s.Type.String()
	if !utils.IsEmpty(s.Version) {
		fullName = fmt.Sprintf("%v@%v", s.Type.String(), s.Version)
	}

	return fullName
}

func (s *Platform) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type    models.PlatformType `yaml:"Type"`
				Version string              `yaml:"Version"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
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

			enum, err := models.PlatformTypeEnumFromString(value)
			if err != nil {
				return err
			}
			s.Type = enum

		} else if len(parts) == 2 {
			platformName := strings.TrimSpace(strings.ToLower(parts[0]))
			platformVersion := strings.TrimSpace(parts[1])

			enum, err := models.PlatformTypeEnumFromString(platformName)
			if err != nil {
				return err
			}
			s.Type = enum

			if utils.IsEmpty(string(s.Type)) {
				return &errors.InvalidPlatformError{Value: platformName}
			}
			s.Version = platformVersion
		} else {
			return &errors.InvalidFormatForPlatformError{Value: value}
		}
	}

	if utils.IsEmpty(string(s.Type)) || s.Type.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}

	return nil
}
