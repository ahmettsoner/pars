package models

import (
	"fmt"
	"strings"
)

type ParsPlatformVersion string

var ParsPlatformVersions = struct {
	BetaV1 ParsPlatformVersion
}{
	BetaV1: "BetaV1",
}

func (c ParsPlatformVersion) String() string {
	switch c {
	case "BetaV1":
		return "BetaV1"
	default:
		return "Unknown"
	}
}

func ParsPlatformVersionEnumFromString(enum string) (ParsPlatformVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("BetaV1"):
		return ParsPlatformVersions.BetaV1, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ParsPlatformVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ParsPlatformVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func ParsPlatformVersionToArray() []ParsPlatformVersion {
	return []ParsPlatformVersion{
		ParsPlatformVersions.BetaV1,
	}
}

type ParsPlatformVersionEnumFlag struct {
	Value ParsPlatformVersion
}

func (e *ParsPlatformVersionEnumFlag) Type() string {
	return "ParsPlatformVersion"
}

func (e *ParsPlatformVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *ParsPlatformVersionEnumFlag) Set(value string) error {
	validEnumValues := ParsPlatformVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
