package core

import (
	"fmt"
	"strings"
)

type WorkspaceDescribeViewType string

var WorkspaceDescribeViewTypes = struct {
	Hierarchical WorkspaceDescribeViewType
	Flat         WorkspaceDescribeViewType
}{
	Hierarchical: "hierarchical",
	Flat:         "flat",
}

func (c WorkspaceDescribeViewType) String() string {
	switch c {
	case "hierarchical":
		return "Hierarchical"
	case "flat":
		return "Flat"
	default:
		return "Unknown"
	}
}

func WorkspaceDescribeViewTypeEnumFromString(enum string) (WorkspaceDescribeViewType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Hierarchical"):
		return WorkspaceDescribeViewTypes.Hierarchical, nil
	case strings.ToLower("H"):
		return WorkspaceDescribeViewTypes.Hierarchical, nil
	case strings.ToLower("Flat"):
		return WorkspaceDescribeViewTypes.Flat, nil
	case strings.ToLower("F"):
		return WorkspaceDescribeViewTypes.Flat, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *WorkspaceDescribeViewType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := WorkspaceDescribeViewTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func WorkspaceDescribeViewTypeToArray() []WorkspaceDescribeViewType {
	return []WorkspaceDescribeViewType{
		WorkspaceDescribeViewTypes.Hierarchical,
		WorkspaceDescribeViewTypes.Flat,
	}
}

type WorkspaceDescribeViewTypeEnumFlag struct {
	Value WorkspaceDescribeViewType
}

func (e *WorkspaceDescribeViewTypeEnumFlag) Type() string {
	return "WorkspaceDescribeViewType"
}

func (e *WorkspaceDescribeViewTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *WorkspaceDescribeViewTypeEnumFlag) Set(value string) error {
	validEnumValues := WorkspaceDescribeViewTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
