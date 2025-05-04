package template

import (
	"fmt"
	"strings"
)

type StructKind string

var StructKinds = struct {
	Code   StructKind
	File   StructKind
	Shared StructKind
}{
	Code:   "Code",
	File:   "File",
	Shared: "Shared",
}

func (c StructKind) String() string {
	switch c {
	case StructKinds.Code:
		return "Code"
	case StructKinds.File:
		return "File"
	case StructKinds.Shared:
		return "Shared"
	default:
		return "Unknown"
	}
}
func StructKindEnumFromString(enum string) (StructKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(StructKinds.Code.String()):
		return StructKinds.Code, nil
	case strings.ToLower(StructKinds.File.String()):
		return StructKinds.File, nil
	case strings.ToLower(StructKinds.Shared.String()):
		return StructKinds.Shared, nil
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
