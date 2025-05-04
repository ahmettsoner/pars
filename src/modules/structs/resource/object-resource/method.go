package objectresource

import (
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Method struct {
	Name        string
	Visibility  VisibilityType
	Parameters  []MethodParameter
	ReturnTypes []DataType
	Hint        Message
	Description Message
	Options     []option.Option
	Labels      []label.Label
	Annotations []Annotation
	Code        string
	Common      bool
}

func NewMethod(name string, visibility VisibilityType, parameters []MethodParameter, returnTypes []DataType, hint Message, description Message, options []option.Option, labels []label.Label, annotations []Annotation, code string, common bool) Method {
	return Method{
		Name:        name,
		Visibility:  visibility,
		Parameters:  parameters,
		ReturnTypes: returnTypes,
		Hint:        hint,
		Description: description,
		Options:     options,
		Labels:      labels,
		Annotations: annotations,
		Code:        code,
		Common:      common,
	}
}

func (s *Method) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject = struct {
				Name        string            `yaml:"Name"`
				Visibility  VisibilityType    `yaml:"Visibility"`
				Parameters  []MethodParameter `yaml:"Parameters"`
				ReturnTypes []DataType        `yaml:"Returns"`
				Hint        Message           `yaml:"Hint"`
				Description Message           `yaml:"Description"`
				Options     []option.Option   `yaml:"Options"`
				Labels      []label.Label     `yaml:"Labels"`
				Annotations []Annotation      `yaml:"Annotations"`
				Code        string            `yaml:"Code"`
				Common      bool              `yaml:"Common"`
			}{
				Common: true,
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Visibility = tempObject.Visibility
				s.Parameters = tempObject.Parameters
				s.ReturnTypes = tempObject.ReturnTypes
				s.Hint = tempObject.Hint
				s.Description = tempObject.Description
				s.Options = tempObject.Options
				s.Labels = tempObject.Labels
				s.Annotations = tempObject.Annotations
				s.Code = tempObject.Code
				s.Common = tempObject.Common
			}
		} else {
			return err
		}

	} else {
		s.Name = value
		s.Common = true
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Method.Name"}
	}

	if utils.IsEmpty(string(s.Visibility)) {
		s.Visibility = VisibilityTypeTypes.Public
	}

	return nil
}
