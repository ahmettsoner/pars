package filetemplate

import (
	"fmt"
	"strings"
)

type TemplateSourceType string

var TemplateSourceTypes = struct {
	File    TemplateSourceType
	Content TemplateSourceType
}{
	File:    "File",
	Content: "Content",
}

func (c TemplateSourceType) String() string {
	switch c {
	case TemplateSourceTypes.File:
		return "File"
	case TemplateSourceTypes.Content:
		return "Content"
	default:
		return "Unknown"
	}
}

func TemplateSourceTypeEnumFromString(enum string) (TemplateSourceType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Content"):
		return TemplateSourceTypes.Content, nil
	case strings.ToLower("File"):
		return TemplateSourceTypes.File, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *TemplateSourceType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := TemplateSourceTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}

// Type to Flag configuration

func TemplateSourceTypeToArray() []TemplateSourceType {
	return []TemplateSourceType{
		TemplateSourceTypes.File,
		TemplateSourceTypes.Content,
	}
}

type TemplateSourceTypeEnumFlag struct {
	Value TemplateSourceType
}

func (e *TemplateSourceTypeEnumFlag) Type() string {
	return "TemplateSourceType"
}

func (e *TemplateSourceTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *TemplateSourceTypeEnumFlag) Set(value string) error {
	validEnumValues := TemplateSourceTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
