package models

import (
	"fmt"
	"strings"
)

type RuntimeType string

var RuntimeTypes = struct {
	Pars   RuntimeType
	Dotnet RuntimeType
	Java   RuntimeType
	GO     RuntimeType
	NodeJS RuntimeType
}{
	Pars:   "pars",
	Dotnet: "dotnet",
	Java:   "java",
	GO:     "go",
	NodeJS: "nodejs",
}

func (c RuntimeType) String() string {
	switch c {
	case "pars":
		return "Pars"
	case "dotnet":
		return "Dotnet"
	case "java":
		return "Java"
	case "go":
		return "GO"
	case "nodejs":
		return "NodeJS"
	default:
		return "Unknown"
	}
}

func RuntimeTypeEnumFromString(enum string) (RuntimeType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Pars"):
		return RuntimeTypes.Pars, nil
	case strings.ToLower("Dotnet"):
		return RuntimeTypes.Dotnet, nil
	case strings.ToLower("Java"):
		return RuntimeTypes.Java, nil
	case strings.ToLower("GO"):
		return RuntimeTypes.GO, nil
	case strings.ToLower("NodeJS"):
		return RuntimeTypes.NodeJS, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *RuntimeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := RuntimeTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
