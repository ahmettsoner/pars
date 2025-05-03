package objectresource

import (
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"
)

type MethodParameter struct {
	Variable
}

func NewMethodParameter(name string, _type DataType, order int, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation) MethodParameter {
	return MethodParameter{
		Variable: NewVariable(name, _type, order, hint, description, options, labels, validation, annotations),
	}
}

func (s *MethodParameter) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempVariableObject struct {
		Variable
	}

	if err := unmarshal(&tempVariableObject); err != nil {
		return err
	} else {
		s.Variable = tempVariableObject.Variable
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "MethodParameter.Name"}
	}

	return nil
}
