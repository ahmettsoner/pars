package structs

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type MethodParameter struct {
	Variable
}

func NewMethodParameter(name string, _type structs.DataType, order int, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation) MethodParameter {
	return MethodParameter{
		Variable: NewVariable(name, _type, order, hint, description, options, labels, validation, annotations),
	}
}

func (e MethodParameter) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
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
	return nil
}
