package resource

import (
	"fmt"
	"strings"
)

type StructKind string

var StructKinds = struct {
	Object      StructKind
	Enumeration StructKind
	Data        StructKind
	// Spread        StructKind
}{
	Object:      "Object",
	Enumeration: "Enumeration",
	Data:        "Data",
	// Spread:        "Spread",
}

func (c StructKind) String() string {
	switch c {
	case StructKinds.Object:
		return "Object"
	case StructKinds.Enumeration:
		return "Enumeration"
	case StructKinds.Data:
		return "Data"
	default:
		return "Unknown"
	}
}
func StructKindEnumFromString(enum string) (StructKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(StructKinds.Object.String()):
		return StructKinds.Object, nil
	case strings.ToLower(StructKinds.Enumeration.String()):
		return StructKinds.Enumeration, nil
	case strings.ToLower(StructKinds.Data.String()):
		return StructKinds.Data, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *StructKind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := StructKindEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
