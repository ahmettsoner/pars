package resource

import (
	"fmt"
	"strings"
)

type ResourceKind string

var ResourceKinds = struct {
	Object      ResourceKind
	Enumeration ResourceKind
	Data        ResourceKind
	// Spread        ResourceKind
}{
	Object:      "Object",
	Enumeration: "Enumeration",
	Data:        "Data",
	// Spread:        "Spread",
}

func (c ResourceKind) String() string {
	switch c {
	case ResourceKinds.Object:
		return "Object"
	case ResourceKinds.Enumeration:
		return "Enumeration"
	case ResourceKinds.Data:
		return "Data"
	default:
		return "Unknown"
	}
}
func ResourceKindEnumFromString(enum string) (ResourceKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(ResourceKinds.Object.String()):
		return ResourceKinds.Object, nil
	case strings.ToLower(ResourceKinds.Enumeration.String()):
		return ResourceKinds.Enumeration, nil
	case strings.ToLower(ResourceKinds.Data.String()):
		return ResourceKinds.Data, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ResourceKind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ResourceKindEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
