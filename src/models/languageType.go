package models

import (
	"fmt"
	"strings"
)

type LanguageType string

var LanguageTypes = struct {
	CSharp     LanguageType
	Java       LanguageType
	GO         LanguageType
	TypeScript LanguageType
}{
	CSharp:     "csharp",
	Java:       "java",
	GO:         "go",
	TypeScript: "typescript",
}

func (c LanguageType) String() string {
	switch c {
	case "csharp":
		return "CSharp"
	case "java":
		return "Java"
	case "go":
		return "GO"
	case "typescript":
		return "TypeScript"
	default:
		return "Unknown"
	}
}

func LanguageTypeEnumFromString(enum string) (LanguageType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("CSharp"):
		return LanguageTypes.CSharp, nil
	case strings.ToLower("Java"):
		return LanguageTypes.Java, nil
	case strings.ToLower("GO"):
		return LanguageTypes.GO, nil
	case strings.ToLower("TypeScript"):
		return LanguageTypes.TypeScript, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *LanguageType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := LanguageTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
