package models

import (
	"fmt"
	"strings"
)

type GoPlatformVersion string

var GoPlatformVersions = struct {
	Go121 GoPlatformVersion
}{
	Go121: "Go121",
}

func (c GoPlatformVersion) String() string {
	switch c {
	case GoPlatformVersions.Go121:
		return "Go121"
	default:
		return "Unknown"
	}
}

func GoPlatformVersionEnumFromString(enum string) (GoPlatformVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(GoPlatformVersions.Go121.String()):
		return GoPlatformVersions.Go121, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *GoPlatformVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := GoPlatformVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func GoPlatformVersionToArray() []GoPlatformVersion {
	return []GoPlatformVersion{
		GoPlatformVersions.Go121,
	}
}

type GoPlatformVersionEnumFlag struct {
	Value GoPlatformVersion
}

func (e *GoPlatformVersionEnumFlag) Type() string {
	return "GoPlatformVersion"
}

func (e *GoPlatformVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *GoPlatformVersionEnumFlag) Set(value string) error {
	validEnumValues := GoPlatformVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
