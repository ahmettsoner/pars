package models

import (
	"fmt"
	"strings"
)

type NodeJSPlatformVersion string

var NodeJSPlatformVersions = struct {
	V15 NodeJSPlatformVersion
	V16 NodeJSPlatformVersion
	V17 NodeJSPlatformVersion
	V18 NodeJSPlatformVersion
	V19 NodeJSPlatformVersion
	V20 NodeJSPlatformVersion
	V21 NodeJSPlatformVersion
}{
	V15: "V15",
	V16: "V16",
	V17: "V17",
	V18: "V18",
	V19: "V19",
	V20: "V20",
	V21: "V21",
}

func (c NodeJSPlatformVersion) String() string {
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

func NodeJSPlatformVersionEnumFromString(enum string) (NodeJSPlatformVersion, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("V15"):
		return NodeJSPlatformVersions.V15, nil
	case strings.ToLower("V16"):
		return NodeJSPlatformVersions.V16, nil
	case strings.ToLower("V17"):
		return NodeJSPlatformVersions.V17, nil
	case strings.ToLower("V18"):
		return NodeJSPlatformVersions.V18, nil
	case strings.ToLower("V19"):
		return NodeJSPlatformVersions.V19, nil
	case strings.ToLower("V20"):
		return NodeJSPlatformVersions.V20, nil
	case strings.ToLower("V21"):
		return NodeJSPlatformVersions.V21, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *NodeJSPlatformVersion) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := NodeJSPlatformVersionEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func NodeJSPlatformVersionToArray() []NodeJSPlatformVersion {
	return []NodeJSPlatformVersion{
		NodeJSPlatformVersions.V15,
		NodeJSPlatformVersions.V16,
		NodeJSPlatformVersions.V17,
		NodeJSPlatformVersions.V18,
		NodeJSPlatformVersions.V19,
		NodeJSPlatformVersions.V20,
		NodeJSPlatformVersions.V21,
	}
}

type NodeJSPlatformVersionEnumFlag struct {
	Value NodeJSPlatformVersion
}

func (e *NodeJSPlatformVersionEnumFlag) Type() string {
	return "NodeJSPlatformVersion"
}

func (e *NodeJSPlatformVersionEnumFlag) String() string {
	return e.Value.String()
}

func (e *NodeJSPlatformVersionEnumFlag) Set(value string) error {
	validEnumValues := NodeJSPlatformVersionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
