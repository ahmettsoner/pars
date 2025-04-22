package project

import (
	"fmt"
	"strings"
)

type StructKind string

var StructKinds = struct {
	Application StructKind
	// Release StructKind
	// Infrastructure StructKind
}{
	Application: "Application",
}

func (c StructKind) String() string {
	switch c {
	case StructKinds.Application:
		return "Application"
	default:
		return "Unknown"
	}
}
func StructKindEnumFromString(enum string) (StructKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(StructKinds.Application.String()):
		return StructKinds.Application, nil
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
