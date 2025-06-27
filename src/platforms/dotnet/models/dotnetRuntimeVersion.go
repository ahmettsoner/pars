package models

import (
	"fmt"
	"strings"
)

type DotnetRuntimeVersion string

var DotnetRuntimeVersions = struct {
	Net5 DotnetRuntimeVersion
	Net6 DotnetRuntimeVersion
	Net7 DotnetRuntimeVersion
	Net8 DotnetRuntimeVersion
}{
	Net5: "Net5",
	Net6: "Net6",
	Net7: "Net7",
	Net8: "Net8",
}

func (c DotnetRuntimeVersion) String() string {
	switch c {
	case "Net5":
		return "Net5"
	case "Net6":
		return "Net6"
	case "Net7":
		return "Net7"
	case "Net8":
		return "Net8"
	default:
		return "Unknown"
	}
}

func DotnetRuntimeVersionEnumFromString(enum string) (DotnetRuntimeVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Net5"):
		return DotnetRuntimeVersions.Net5, nil
	case strings.ToLower("Net6"):
		return DotnetRuntimeVersions.Net6, nil
	case strings.ToLower("Net7"):
		return DotnetRuntimeVersions.Net7, nil
	case strings.ToLower("Net8"):
		return DotnetRuntimeVersions.Net8, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *DotnetRuntimeVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := DotnetRuntimeVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func DotnetRuntimeVersionToArray() []DotnetRuntimeVersion {
	return []DotnetRuntimeVersion{
		DotnetRuntimeVersions.Net5,
		DotnetRuntimeVersions.Net6,
		DotnetRuntimeVersions.Net7,
		DotnetRuntimeVersions.Net8,
	}
}

type DotnetRuntimeVersionEnumFlag struct {
	Value DotnetRuntimeVersion
}

func (e *DotnetRuntimeVersionEnumFlag) Type() string {
	return "DotnetRuntimeVersion"
}

func (e *DotnetRuntimeVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *DotnetRuntimeVersionEnumFlag) Set(value string) error {
	validEnumValues := DotnetRuntimeVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
