package objectresource

import (
	"fmt"
	"strings"
)

type ModifierType string

var ModifierTypes = struct {
	Object ModifierType
	Array  ModifierType
}{
	Object: "Object",
	Array:  "Array",
}

func (c ModifierType) String() string {
	switch c {

	case ModifierTypes.Object:
		return "Object"
	case ModifierTypes.Array:
		return "Array"
	default:
		return "Unknown"
	}
}

func ModifierTypeEnumFromString(enum string) (ModifierType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Object"):
		return ModifierTypes.Object, nil
	case strings.ToLower("Array"):
		return ModifierTypes.Array, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ModifierType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ModifierTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

func ModifierTypeToArray() []ModifierType {
	return []ModifierType{
		ModifierTypes.Object,
		ModifierTypes.Array,
	}
}

type ModifierTypeEnumFlag struct {
	Value ModifierType
}

func (e *ModifierTypeEnumFlag) Type() string {
	return "ModifierType"
}

func (e *ModifierTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *ModifierTypeEnumFlag) Set(value string) error {
	validEnumValues := ModifierTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
