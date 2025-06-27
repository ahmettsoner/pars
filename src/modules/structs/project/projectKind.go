package project

import (
	"fmt"
	"strings"
)

type ProjectKind string

var ProjectKinds = struct {
	Application ProjectKind
	// Release ProjectKind
	// Infrastructure ProjectKind
}{
	Application: "Application",
}

func (c ProjectKind) String() string {
	switch c {
	case ProjectKinds.Application:
		return "Application"
	default:
		return "Unknown"
	}
}
func ProjectKindEnumFromString(enum string) (ProjectKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(ProjectKinds.Application.String()):
		return ProjectKinds.Application, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ProjectKind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ProjectKindEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
