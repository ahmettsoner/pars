package models

import (
	"fmt"
	"strings"
)

type PlatformType string

// create array from enum  values for iteration over them
var PlatformTypes = struct {
	Pars      PlatformType
	Dotnet    PlatformType
	Java      PlatformType
	GO        PlatformType
	NodeJS    PlatformType
	PHP       PlatformType
	Angular   PlatformType
	VueJS     PlatformType
	React     PlatformType
	Flutter   PlatformType
	Chrome    PlatformType
	VSCode    PlatformType
	Terraform PlatformType
}{
	Pars:      "Pars",
	Dotnet:    "Dotnet",
	Java:      "Java",
	GO:        "GO",
	NodeJS:    "NodeJS",
	PHP:       "PHP",
	Angular:   "Angular",
	VueJS:     "VueJS",
	React:     "React",
	Chrome:    "Chrome",
	VSCode:    "VSCode",
	Terraform: "Terraform",
}

func (c PlatformType) String() string {
	switch c {
	case PlatformTypes.Pars:
		return "Pars"
	case PlatformTypes.Dotnet:
		return "Dotnet"
	case PlatformTypes.Java:
		return "Java"
	case PlatformTypes.GO:
		return "GO"
	case PlatformTypes.NodeJS:
		return "NodeJS"
	case PlatformTypes.PHP:
		return "PHP"
	case PlatformTypes.Angular:
		return "Angular"
	case PlatformTypes.VueJS:
		return "VueJS"
	case PlatformTypes.React:
		return "React"
	case PlatformTypes.Flutter:
		return "Flutter"
	case PlatformTypes.Chrome:
		return "Chrome"
	case PlatformTypes.VSCode:
		return "VSCode"
	case PlatformTypes.Terraform:
		return "Terraform"
	default:
		return "Unknown"
	}
}

func PlatformTypeEnumFromString(enum string) (PlatformType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Pars"):
		return PlatformTypes.Pars, nil
	case strings.ToLower("Dotnet"):
		return PlatformTypes.Dotnet, nil
	case strings.ToLower("Java"):
		return PlatformTypes.Java, nil
	case strings.ToLower("GO"):
		return PlatformTypes.GO, nil
	case strings.ToLower("NodeJS"):
		return PlatformTypes.NodeJS, nil
	case strings.ToLower("PHP"):
		return PlatformTypes.PHP, nil
	case strings.ToLower("Angular"):
		return PlatformTypes.Angular, nil
	case strings.ToLower("VueJS"):
		return PlatformTypes.VueJS, nil
	case strings.ToLower("React"):
		return PlatformTypes.React, nil
	case strings.ToLower("Flutter"):
		return PlatformTypes.Flutter, nil
	case strings.ToLower("Chrome"):
		return PlatformTypes.Chrome, nil
	case strings.ToLower("VSCode"):
		return PlatformTypes.VSCode, nil
	case strings.ToLower("Terraform"):
		return PlatformTypes.Terraform, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *PlatformType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := PlatformTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

func GetPlatformTypeEnumArray() []PlatformType {
	return []PlatformType{
		PlatformTypes.Pars,
		PlatformTypes.Dotnet,
		// PlatformTypes.Java,
		PlatformTypes.GO,
		// PlatformTypes.NodeJS,
		// PlatformTypes.PHP,
		PlatformTypes.Angular,
		// PlatformTypes.VueJS,
		// PlatformTypes.React,
		// PlatformTypes.Flutter,
	}
}
