package objectresource

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
)

type MethodParameter struct {
	Variable
}

func NewMethodParameter(name string, _type DataType, order int, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation) MethodParameter {
	return MethodParameter{
		Variable: NewVariable(name, _type, order, hint, description, options, labels, validation, annotations),
	}
}

func (e MethodParameter) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Name, v.Required),
	)
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
