package models

import (
	"fmt"
	"strings"
)

type NodeJSRuntimeVersion string

var NodeJSRuntimeVersions = struct {
	V15 NodeJSRuntimeVersion
	V16 NodeJSRuntimeVersion
	V17 NodeJSRuntimeVersion
	V18 NodeJSRuntimeVersion
	V19 NodeJSRuntimeVersion
	V20 NodeJSRuntimeVersion
	V21 NodeJSRuntimeVersion
}{
	V15: "V15",
	V16: "V16",
	V17: "V17",
	V18: "V18",
	V19: "V19",
	V20: "V20",
	V21: "V21",
}

func (c NodeJSRuntimeVersion) String() string {
	switch c {
	case "V15":
		return "V15"
	case "V16":
		return "V16"
	case "V17":
		return "V17"
	case "V18":
		return "V18"
	case "V19":
		return "V19"
	case "V20":
		return "V20"
	case "V21":
		return "V21"
	default:
		return "Unknown"
	}
}

func NodeJSRuntimeVersionEnumFromString(enum string) (NodeJSRuntimeVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("V15"):
		return NodeJSRuntimeVersions.V15, nil
	case strings.ToLower("V16"):
		return NodeJSRuntimeVersions.V16, nil
	case strings.ToLower("V17"):
		return NodeJSRuntimeVersions.V17, nil
	case strings.ToLower("V18"):
		return NodeJSRuntimeVersions.V18, nil
	case strings.ToLower("V19"):
		return NodeJSRuntimeVersions.V19, nil
	case strings.ToLower("V20"):
		return NodeJSRuntimeVersions.V20, nil
	case strings.ToLower("V21"):
		return NodeJSRuntimeVersions.V21, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *NodeJSRuntimeVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := NodeJSRuntimeVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func NodeJSRuntimeVersionToArray() []NodeJSRuntimeVersion {
	return []NodeJSRuntimeVersion{
		NodeJSRuntimeVersions.V15,
		NodeJSRuntimeVersions.V16,
		NodeJSRuntimeVersions.V17,
		NodeJSRuntimeVersions.V18,
		NodeJSRuntimeVersions.V19,
		NodeJSRuntimeVersions.V20,
		NodeJSRuntimeVersions.V21,
	}
}

type NodeJSRuntimeVersionEnumFlag struct {
	Value NodeJSRuntimeVersion
}

func (e *NodeJSRuntimeVersionEnumFlag) Type() string {
	return "NodeJSRuntimeVersion"
}

func (e *NodeJSRuntimeVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *NodeJSRuntimeVersionEnumFlag) Set(value string) error {
	validEnumValues := NodeJSRuntimeVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
