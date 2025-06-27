package models

import (
	"fmt"
	"strings"
)

type AngularPlatformVersion string

var AngularPlatformVersions = struct {
	V15 AngularPlatformVersion
	V16 AngularPlatformVersion
	V17 AngularPlatformVersion
}{
	V15: "V15",
	V16: "V16",
	V17: "V17",
}

func (c AngularPlatformVersion) String() string {
	switch c {
	case "V15":
		return "V15"
	case "V16":
		return "V16"
	case "V17":
		return "V17"
	default:
		return "Unknown"
	}
}

func AngularPlatformVersionEnumFromString(enum string) (AngularPlatformVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("V15"):
		return AngularPlatformVersions.V15, nil
	case strings.ToLower("V16"):
		return AngularPlatformVersions.V16, nil
	case strings.ToLower("V17"):
		return AngularPlatformVersions.V17, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *AngularPlatformVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := AngularPlatformVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func AngularPlatformVersionToArray() []AngularPlatformVersion {
	return []AngularPlatformVersion{
		AngularPlatformVersions.V15,
		AngularPlatformVersions.V16,
		AngularPlatformVersions.V17,
	}
}

type AngularPlatformVersionEnumFlag struct {
	Value AngularPlatformVersion
}

func (e *AngularPlatformVersionEnumFlag) Type() string {
	return "AngularPlatformVersion"
}

func (e *AngularPlatformVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *AngularPlatformVersionEnumFlag) Set(value string) error {
	validEnumValues := AngularPlatformVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
