package task

import (
	"fmt"
	"strings"
)

type TaskKind string

var TaskKinds = struct {
	Common TaskKind
}{
	Common: "Common",
}

func (c TaskKind) String() string {
	switch c {
	case TaskKinds.Common:
		return "Common"
	default:
		return "Unknown"
	}
}
func TaskKindEnumFromString(enum string) (TaskKind, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(TaskKinds.Common.String()):
		return TaskKinds.Common, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *TaskKind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := TaskKindEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
