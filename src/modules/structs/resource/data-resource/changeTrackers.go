package dataresource

import (
	"fmt"
	"strings"
)

type ChangeTracker string

var ChangeTrackers = struct {
	Never    ChangeTracker
	OnCreate ChangeTracker
	OnChange ChangeTracker
	Always   ChangeTracker
	// OnDeleted ChangeTracker
}{
	Never:    "Never",
	OnCreate: "OnCreate",
	OnChange: "OnChange",
	Always:   "Always",
}

func (c ChangeTracker) String() string {
	switch c {

	case ChangeTrackers.Never:
		return "Never"
	case ChangeTrackers.OnCreate:
		return "OnCreate"
	case ChangeTrackers.OnChange:
		return "OnChange"
	case ChangeTrackers.Always:
		return "Always"
	default:
		return "Unknown"
	}
}

func ChangeTrackerEnumFromString(enum string) (ChangeTracker, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Never"):
		return ChangeTrackers.Never, nil
	case strings.ToLower("OnCreate"):
		return ChangeTrackers.OnCreate, nil
	case strings.ToLower("OnChange"):
		return ChangeTrackers.OnChange, nil
	case strings.ToLower("Always"):
		return ChangeTrackers.Always, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ChangeTracker) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ChangeTrackerEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

func ChangeTrackerToArray() []ChangeTracker {
	return []ChangeTracker{
		ChangeTrackers.Never,
		ChangeTrackers.OnCreate,
		ChangeTrackers.OnChange,
		ChangeTrackers.Always,
	}
}

type ChangeTrackerEnumFlag struct {
	Value ChangeTracker
}

func (e *ChangeTrackerEnumFlag) Type() string {
	return "ChangeTracker"
}

func (e *ChangeTrackerEnumFlag) String() string {
	return e.Value.String()
}

func (e *ChangeTrackerEnumFlag) Set(value string) error {
	validEnumValues := ChangeTrackerToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
