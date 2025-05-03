package applicationproject

import (
	"strings"

	"parsdevkit.net/models"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Runtime struct {
	Type    models.RuntimeType
	Version string
}

func NewRuntime(_type models.RuntimeType, version string) Runtime {
	return Runtime{
		Type:    _type,
		Version: version,
	}
}

func NewRuntime_Basic(_type models.RuntimeType) Runtime {
	return Runtime{
		Type: _type,
	}
}

func (s *Runtime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type    models.RuntimeType `yaml:"Type"`
				Version string             `yaml:"Version"`
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

			enum, err := models.RuntimeTypeEnumFromString(value)
			if err != nil {
				return err
			}
			s.Type = enum
		} else if len(parts) == 2 {
			runtimeName := strings.TrimSpace(strings.ToLower(parts[0]))
			runtimeVersion := strings.TrimSpace(parts[1])

			enum, err := models.RuntimeTypeEnumFromString(runtimeName)
			if err != nil {
				return err
			}
			s.Type = enum

			if utils.IsEmpty(string(s.Type)) {
				return &errors.InvalidRuntimeError{Value: runtimeName}
			}
			s.Version = runtimeVersion
		} else {
			return &errors.InvalidFormatForRuntimeError{Value: value}
		}
	}

	if utils.IsEmpty(string(s.Type)) || s.Type.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}

	return nil
}
