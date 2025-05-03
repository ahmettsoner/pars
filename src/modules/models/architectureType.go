package models

import (
	"fmt"
	"strings"
)

type ArchitectureType string

var ArchitectureTypes = struct {
	None      ArchitectureType
	Clean     ArchitectureType
	Onion     ArchitectureType
	Hexagonal ArchitectureType
	// MVC                   ArchitectureType
	// MVP                   ArchitectureType
	// MVVM                  ArchitectureType
	// MVVMC                 ArchitectureType
	// VIPER                 ArchitectureType
}{
	None:      "none",
	Clean:     "clean",
	Onion:     "onion",
	Hexagonal: "hexagonal",
	//Patterns
	// MVC:   4,
	// MVP:   5,
	// MVVM:  6,
	// MVVMC: 7,
	// VIPER: 8,
}

func (c ArchitectureType) String() string {
	switch c {
	case "none":
		return "None"
	case "clean":
		return "Clean"
	case "onion":
		return "Onion"
	case "hexagonal":
		return "Hexagonal"
	// case 4:
	// 	return "MVC"
	// case 5:
	// 	return "MVP"
	// case 6:
	// 	return "MVVM"
	// case 7:
	// 	return "MVVMC"
	// case 8:
	// 	return "VIPER"
	default:
		return "Unknown"
	}
}

func ArchitectureTypeEnumFromString(enum string) (ArchitectureType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("None"):
		return ArchitectureTypes.None, nil
	case strings.ToLower("Clean"):
		return ArchitectureTypes.Clean, nil
	case strings.ToLower("Onion"):
		return ArchitectureTypes.Onion, nil
	case strings.ToLower("Hexagonal"):
		return ArchitectureTypes.Hexagonal, nil
	// case strings.ToLower("MVC"):
	// 	return ArchitectureTypes.MVC, nil
	// case strings.ToLower("MVP"):
	// 	return ArchitectureTypes.MVP, nil
	// case strings.ToLower("MVVM"):
	// 	return ArchitectureTypes.MVVM, nil
	// case strings.ToLower("MVVMC"):
	// 	return ArchitectureTypes.MVVMC, nil
	// case strings.ToLower("VIPER"):
	// 	return ArchitectureTypes.VIPER, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ArchitectureType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ArchitectureTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func ArchitectureTypeToArray() []ArchitectureType {
	return []ArchitectureType{
		ArchitectureTypes.None,
		ArchitectureTypes.Clean,
		ArchitectureTypes.Onion,
		ArchitectureTypes.Hexagonal,
		// ArchitectureTypes.MVC,
		// ArchitectureTypes.MVP,
		// ArchitectureTypes.MVVM,
		// ArchitectureTypes.MVVMC,
		// ArchitectureTypes.VIPER,
	}
}

type ArchitectureTypeEnumFlag struct {
	Value ArchitectureType
}

func (e *ArchitectureTypeEnumFlag) Type() string {
	return "ArchitectureType"
}

func (e *ArchitectureTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *ArchitectureTypeEnumFlag) Set(value string) error {
	validEnumValues := ArchitectureTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
