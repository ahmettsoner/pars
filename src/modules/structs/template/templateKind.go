package template

import (
	"fmt"
	"strings"
)

type TemplateKind string

var TemplateKinds = struct {
	Code   TemplateKind
	File   TemplateKind
	Shared TemplateKind
}{
	Code:   "Code",
	File:   "File",
	Shared: "Shared",
}

func (c TemplateKind) String() string {
	switch c {
	case TemplateKinds.Code:
		return "Code"
	case TemplateKinds.File:
		return "File"
	case TemplateKinds.Shared:
		return "Shared"
	default:
		return "Unknown"
	}
}
func TemplateKindEnumFromString(enum string) (TemplateKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(TemplateKinds.Code.String()):
		return TemplateKinds.Code, nil
	case strings.ToLower(TemplateKinds.File.String()):
		return TemplateKinds.File, nil
	case strings.ToLower(TemplateKinds.Shared.String()):
		return TemplateKinds.Shared, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *TemplateKind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := TemplateKindEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
