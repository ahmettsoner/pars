package models

import (
	"fmt"
	"strings"
)

type GoRuntimeVersion string

var GoRuntimeVersions = struct {
	Go121 GoRuntimeVersion
}{
	Go121: "Go121",
}

func (c GoRuntimeVersion) String() string {
	switch c {
	case "Go121":
		return "Go121"
	default:
		return "Unknown"
	}
}

func GoRuntimeVersionEnumFromString(enum string) (GoRuntimeVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Go121"):
		return GoRuntimeVersions.Go121, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *GoRuntimeVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := GoRuntimeVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func GoRuntimeVersionToArray() []GoRuntimeVersion {
	return []GoRuntimeVersion{
		GoRuntimeVersions.Go121,
	}
}

type GoRuntimeVersionEnumFlag struct {
	Value GoRuntimeVersion
}

func (e *GoRuntimeVersionEnumFlag) Type() string {
	return "GoRuntimeVersion"
}

func (e *GoRuntimeVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *GoRuntimeVersionEnumFlag) Set(value string) error {
	validEnumValues := GoRuntimeVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
