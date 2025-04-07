package models

import (
	"fmt"
	"strings"
)

// TODO: Default değer tanımlanabilmeli? çözüm planlanmalı
type DotnetPlatformVersion string

var DotnetPlatformVersions = struct {
	Net5 DotnetPlatformVersion
	Net6 DotnetPlatformVersion
	Net7 DotnetPlatformVersion
	Net8 DotnetPlatformVersion
}{
	Net5: "Net5",
	Net6: "Net6",
	Net7: "Net7",
	Net8: "Net8",
}

func (c DotnetPlatformVersion) String() string {
	switch c {
	case DotnetPlatformVersions.Net5:
		return "Net5"
	case DotnetPlatformVersions.Net6:
		return "Net6"
	case DotnetPlatformVersions.Net7:
		return "Net7"
	case DotnetPlatformVersions.Net8:
		return "Net8"
	default:
		return "Unknown"
	}
}

func DotnetPlatformVersionEnumFromString(enum string) (DotnetPlatformVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Net5"):
		return DotnetPlatformVersions.Net5, nil
	case strings.ToLower("Net6"):
		return DotnetPlatformVersions.Net6, nil
	case strings.ToLower("Net7"):
		return DotnetPlatformVersions.Net7, nil
	case strings.ToLower("Net8"):
		return DotnetPlatformVersions.Net8, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *DotnetPlatformVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := DotnetPlatformVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func DotnetPlatformVersionToArray() []DotnetPlatformVersion {
	return []DotnetPlatformVersion{
		DotnetPlatformVersions.Net5,
		DotnetPlatformVersions.Net6,
		DotnetPlatformVersions.Net7,
		DotnetPlatformVersions.Net8,
	}
}

type DotnetPlatformVersionEnumFlag struct {
	Value DotnetPlatformVersion
}

func (e *DotnetPlatformVersionEnumFlag) Type() string {
	return "DotnetPlatformVersion"
}

func (e *DotnetPlatformVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *DotnetPlatformVersionEnumFlag) Set(value string) error {
	validEnumValues := DotnetPlatformVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
