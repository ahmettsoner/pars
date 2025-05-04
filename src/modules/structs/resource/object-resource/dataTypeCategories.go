package objectresource

import (
	"fmt"
	"strings"
)

type DataTypeCategory string

var DataTypeCategories = struct {
	Value     DataTypeCategory
	Reference DataTypeCategory
	Resource  DataTypeCategory
}{
	Value:     "Value",
	Reference: "Reference",
	Resource:  "Resource",
}

func (c DataTypeCategory) String() string {
	switch c {

	case DataTypeCategories.Value:
		return "Value"
	case DataTypeCategories.Reference:
		return "Reference"
	case DataTypeCategories.Resource:
		return "Resource"
	default:
		return "Unknown"
	}
}

func DataTypeTypeEnumFromString(enum string) (DataTypeCategory, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Value"):
		return DataTypeCategories.Value, nil
	case strings.ToLower("Reference"):
		return DataTypeCategories.Reference, nil
	case strings.ToLower("Resource"):
		return DataTypeCategories.Resource, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *DataTypeCategory) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := DataTypeTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

func DataTypeTypeToArray() []DataTypeCategory {
	return []DataTypeCategory{
		DataTypeCategories.Value,
		DataTypeCategories.Reference,
		DataTypeCategories.Resource,
	}
}

type DataTypeTypeEnumFlag struct {
	Value DataTypeCategory
}

func (e *DataTypeTypeEnumFlag) Type() string {
	return "DataTypeCategory"
}

func (e *DataTypeTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *DataTypeTypeEnumFlag) Set(value string) error {
	validEnumValues := DataTypeTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
