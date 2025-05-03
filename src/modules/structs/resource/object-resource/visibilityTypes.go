package objectresource

import (
	"fmt"
	"strings"
)

type VisibilityType string

var VisibilityTypeTypes = struct {
	Public    VisibilityType
	Protected VisibilityType
	Private   VisibilityType
}{
	Public:    "Public",
	Protected: "Protected",
	Private:   "Private",
}

func (c VisibilityType) String() string {
	switch c {

	case VisibilityTypeTypes.Public:
		return "Public"
	case VisibilityTypeTypes.Protected:
		return "Protected"
	case VisibilityTypeTypes.Private:
		return "Private"
	default:
		return "Unknown"
	}
}

func VisibilityTypeEnumFromString(enum string) (VisibilityType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Public"):
		return VisibilityTypeTypes.Public, nil
	case strings.ToLower("Protected"):
		return VisibilityTypeTypes.Protected, nil
	case strings.ToLower("Private"):
		return VisibilityTypeTypes.Private, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *VisibilityType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := VisibilityTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

func VisibilityTypeToArray() []VisibilityType {
	return []VisibilityType{
		VisibilityTypeTypes.Public,
		VisibilityTypeTypes.Protected,
		VisibilityTypeTypes.Private,
	}
}

type VisibilityTypeEnumFlag struct {
	Value VisibilityType
}

func (e *VisibilityTypeEnumFlag) Type() string {
	return "VisibilityType"
}

func (e *VisibilityTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *VisibilityTypeEnumFlag) Set(value string) error {
	validEnumValues := VisibilityTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
