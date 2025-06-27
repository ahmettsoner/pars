package models

import (
	"fmt"
	"strings"
)

type CSharpVersion string

var CSharpVersions = struct {
	V5 CSharpVersion
	V6 CSharpVersion
	V7 CSharpVersion
	V8 CSharpVersion
}{
	V5: "V5",
	V6: "V6",
	V7: "V7",
	V8: "V8",
}

func (c CSharpVersion) String() string {
	switch c {
	case "V5":
		return "V5"
	case "V6":
		return "V6"
	case "V7":
		return "V7"
	case "V8":
		return "V8"
	default:
		return "Unknown"
	}
}

func CSharpVersionEnumFromString(enum string) (CSharpVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("V5"):
		return CSharpVersions.V5, nil
	case strings.ToLower("V6"):
		return CSharpVersions.V6, nil
	case strings.ToLower("V7"):
		return CSharpVersions.V7, nil
	case strings.ToLower("V8"):
		return CSharpVersions.V8, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *CSharpVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := CSharpVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func CSharpVersionToArray() []CSharpVersion {
	return []CSharpVersion{
		CSharpVersions.V5,
		CSharpVersions.V6,
		CSharpVersions.V7,
		CSharpVersions.V8,
	}
}

type CSharpVersionEnumFlag struct {
	Value CSharpVersion
}

func (e *CSharpVersionEnumFlag) Type() string {
	return "CSharpVersion"
}

func (e *CSharpVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *CSharpVersionEnumFlag) Set(value string) error {
	validEnumValues := CSharpVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
