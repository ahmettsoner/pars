package project

import (
	"fmt"
	"strings"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
	_string "parsdevkit.net/pkg/utilities/string"
)

type Dependency struct {
	Name    string
	Version string
}

func NewDependency(name string, version string) Dependency {
	return Dependency{
		Name:    name,
		Version: version,
	}
}

func NewDependency_Basic(name string) Dependency {
	return Dependency{
		Name: name,
	}
}
func (e Dependency) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *Dependency) GetFullName() string {
	fullName := s.Name
	if !_string.IsEmpty(s.Version) {
		fullName = fmt.Sprintf("%v@%v", s.Name, s.Version)
	}

	return fullName
}

func (s *Dependency) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name    string `yaml:"Name"`
				Version string `yaml:"Version"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err
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

			if _string.IsEmpty(s.Name) {
				return &errors.InvalidDependencyError{Value: packageName}
			}
			s.Version = packageVersion
		} else {
			return &errors.InvalidFormatForDependencyError{Value: value}
		}
	}

	if _string.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
