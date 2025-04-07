package models

import (
	"fmt"
	"strings"
)

type ClassicLayerType int

var ClassicLayerTypes = struct {
	Service            ClassicLayerType
	ServiceAbstracts   ClassicLayerType
	ServiceConcretes   ClassicLayerType
	DataTransferObject ClassicLayerType
}{
	DataTransferObject: 100,
	Service:            200,
	ServiceAbstracts:   210,
	ServiceConcretes:   220,
}

func (c ClassicLayerType) String() string {
	switch c {
	case 100:
		return "Data Transfer Object"
	case 200:
		return "Service"
	case 210:
		return "Service / Abstracts"
	case 220:
		return "Service / Concretes"
	default:
		return "Unknown"
	}
}

// Type to Flag configuration

func ClassicLayerTypeToArray() []ClassicLayerType {
	return []ClassicLayerType{
		ClassicLayerTypes.Service,
		ClassicLayerTypes.ServiceAbstracts,
		ClassicLayerTypes.ServiceConcretes,
		ClassicLayerTypes.DataTransferObject,
	}
}

type ClassicLayerTypeEnumFlag struct {
	Value ClassicLayerType
}

func (e *ClassicLayerTypeEnumFlag) Type() string {
	return "ClassicLayerType"
}

func (e *ClassicLayerTypeEnumFlag) String() string {
	return e.Value.String()
}

func (e *ClassicLayerTypeEnumFlag) Set(value string) error {
	validEnumValues := ClassicLayerTypeToArray()
	for _, validValue := range validEnumValues {
		if strings.EqualFold(value, validValue.String()) {
			e.Value = validValue
			return nil
		}
	}
	return fmt.Errorf("invalid value: %s, valid values are %v", value, validEnumValues)
}
