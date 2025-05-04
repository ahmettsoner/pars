package models

import (
	"fmt"
	"strings"
)

type DotnetWebAppOption string

var DotnetWebAppOptions = struct {
	MVC   DotnetWebAppOption
	Razor DotnetWebAppOption
}{
	MVC:   "mvc",
	Razor: "razor",
}

func (c DotnetWebAppOption) String() string {
	switch c {
	case "mvc":
		return "MVC"
	case "razor":
		return "Razor"
	default:
		return "Unknown"
	}
}

// Type to Flag configuration

func DotnetWebAppOptionToArray() []DotnetWebAppOption {
	return []DotnetWebAppOption{
		DotnetWebAppOptions.MVC,
		DotnetWebAppOptions.Razor,
	}
}

type DotnetWebAppOptionEnumFlag struct {
	Value DotnetWebAppOption
}

func (e *DotnetWebAppOptionEnumFlag) Type() string {
	return "DotnetWebAppOption"
}

func (e *DotnetWebAppOptionEnumFlag) String() string {
	return e.Value.String()
}

func (e *DotnetWebAppOptionEnumFlag) Set(value string) error {
	validEnumValues := DotnetWebAppOptionToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
