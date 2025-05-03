package dataresource

import (
	"fmt"
	"strings"
)

type ResourceSourceType string

var ResourceSourceTypes = struct {
	File ResourceSourceType
	Data ResourceSourceType
}{
	File: "File",
	Data: "Data",
}

func (c ResourceSourceType) String() string {
	switch c {
	case ResourceSourceTypes.File:
		return "File"
	case ResourceSourceTypes.Data:
		return "Data"
	default:
		return "Unknown"
	}
}

func ResourceSourceTypeEnumFromString(enum string) (ResourceSourceType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Data"):
		return ResourceSourceTypes.Data, nil
	case strings.ToLower("File"):
		return ResourceSourceTypes.File, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ResourceSourceType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ResourceSourceTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func ResourceSourceTypeToArray() []ResourceSourceType {
	return []ResourceSourceType{
		ResourceSourceTypes.File,
		ResourceSourceTypes.Data,
	}
}

type ResourceSourceTypeEnumFlag struct {
	Value ResourceSourceType
}

func (e *ResourceSourceTypeEnumFlag) Type() string {
	return "ResourceSourceType"
}

func (e *ResourceSourceTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *ResourceSourceTypeEnumFlag) Set(value string) error {
	validEnumValues := ResourceSourceTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
