package models

import (
	"fmt"
	"strings"
)

type DesignType string

var DesignTypes = struct {
	Classic DesignType
	DDD     DesignType
	Cqrs    DesignType
}{
	Classic: "classic",
	DDD:     "ddd",
	Cqrs:    "cqrs",
}

func (c DesignType) String() string {
	switch c {
	case "classic":
		return "Classic"
	case "ddd":
		return "DDD"
	case "cqrs":
		return "Cqrs"
	default:
		return "Unknown"
	}
}

func DesignTypeEnumFromString(enum string) (DesignType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Classic"):
		return DesignTypes.Classic, nil
	case strings.ToLower("DDD"):
		return DesignTypes.DDD, nil
	case strings.ToLower("Cqrs"):
		return DesignTypes.Cqrs, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *DesignType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := DesignTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func DesignTypeToArray() []DesignType {
	return []DesignType{
		DesignTypes.Classic,
		DesignTypes.DDD,
		DesignTypes.Cqrs,
	}
}

type DesignTypeEnumFlag struct {
	Value DesignType
}

func (e *DesignTypeEnumFlag) Type() string {
	return "DesignType"
}

func (e *DesignTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *DesignTypeEnumFlag) Set(value string) error {
	validEnumValues := DesignTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
