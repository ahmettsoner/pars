package applicationproject

import (
	"fmt"
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Package struct {
	Name    string
	Version string
}

func NewPackage(name string, version string) Package {
	return Package{
		Name:    name,
		Version: version,
	}
}

func NewPackage_Basic(name string) Package {
	return Package{
		Name: name,
	}
}

func (s *Package) GetFullName() string {
	fullName := s.Name
	if !utils.IsEmpty(s.Version) {
		fullName = fmt.Sprintf("%v@%v", s.Name, s.Version)
	}

	return fullName
}

func (s *Package) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name    string `yaml:"Name"`
				Version string `yaml:"Version"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Version = tempObject.Version
			}

		} else {
			return err
		}

	} else {
		hasAtPrefix := strings.HasPrefix(value, "@")
		if hasAtPrefix {
			value = strings.TrimPrefix(value, "@")
		}

		var parts []string = strings.Split(value, "@")
		if len(parts) == 1 {
			// No specific version provided so assume latest

			if hasAtPrefix {
				s.Name = fmt.Sprintf("@%v", value)
			} else {
				s.Name = value
			}
		} else if len(parts) == 2 {
			packageName := strings.TrimSpace(parts[0])
			packageVersion := strings.TrimSpace(parts[1])

			if hasAtPrefix {
				s.Name = fmt.Sprintf("@%v", packageName)
			} else {
				s.Name = packageName
			}

			if utils.IsEmpty(s.Name) {
				return &errors.InvalidPackageError{Value: packageName}
			}
			s.Version = packageVersion
		} else {
			return &errors.InvalidFormatForPackageError{Value: value}
		}
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
