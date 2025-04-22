package models

import (
	"fmt"
	"strings"
)

type MethodologyType string

var MethodologyTypes = struct {
	Basic   MethodologyType
	Layered MethodologyType
	NTier   MethodologyType
}{
	Basic:   "basic",
	Layered: "layered",
	NTier:   "ntier",
}

func (c MethodologyType) String() string {
	switch c {
	case "basic":
		return "Basic"
	case "layered":
		return "Layered"
	case "ntier":
		return "NTier"
	default:
		return "Unknown"
	}
}

func MethodologyTypeEnumFromString(enum string) (MethodologyType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Basic"):
		return MethodologyTypes.Basic, nil
	case strings.ToLower("Layered"):
		return MethodologyTypes.Layered, nil
	case strings.ToLower("NTier"):
		return MethodologyTypes.NTier, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *MethodologyType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := MethodologyTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func MethodologyTypeToArray() []MethodologyType {
	return []MethodologyType{
		MethodologyTypes.Basic,
		MethodologyTypes.Layered,
		MethodologyTypes.NTier,
	}
}

type MethodologyTypeEnumFlag struct {
	Value MethodologyType
}

func (e *MethodologyTypeEnumFlag) Type() string {
	return "MethodologyType"
}

func (e *MethodologyTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *MethodologyTypeEnumFlag) Set(value string) error {
	validEnumValues := MethodologyTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
