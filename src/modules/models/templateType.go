package models

import (
	"fmt"
	"strings"
)

type TemplateType string

var TemplateTypes = struct {
	Empty  TemplateType
	Simple TemplateType
	Sample TemplateType
}{
	Empty:  "empty",
	Simple: "simple",
	Sample: "sample",
}

func (c TemplateType) String() string {
	switch c {
	case "empty":
		return "Empty"
	case "simple":
		return "Simple"
	case "sample":
		return "Sample"
	default:
		return "Unknown"
	}
}

// Type to Flag configuration

func TemplateTypeToArray() []TemplateType {
	return []TemplateType{
		TemplateTypes.Empty,
		TemplateTypes.Simple,
		TemplateTypes.Sample,
	}
}

type TemplateTypeEnumFlag struct {
	Value TemplateType
}

func (e *TemplateTypeEnumFlag) Type() string {
	return "TemplateType"
}

func (e *TemplateTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *TemplateTypeEnumFlag) Set(value string) error {
	validEnumValues := TemplateTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}

func TemplateTypeEnumFromString(enum string) (TemplateType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Empty"):
		return TemplateTypes.Empty, nil
	case strings.ToLower("Simple"):
		return TemplateTypes.Simple, nil
	case strings.ToLower("Sample"):
		return TemplateTypes.Sample, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *TemplateType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := TemplateTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
