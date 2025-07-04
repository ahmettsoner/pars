package object_resource_payload

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"

	"gopkg.in/yaml.v3"
)

type Method struct {
	Name        string
	Visibility  structs.VisibilityType
	Parameters  []MethodParameter
	ReturnTypes []structs.DataType
	Hint        Message
	Description Message
	Options     []option.Option
	Labels      []label.Label
	Annotations []Annotation
	Code        string
	Common      bool
}

func NewMethod(name string, visibility structs.VisibilityType, parameters []MethodParameter, returnTypes []structs.DataType, hint Message, description Message, options []option.Option, labels []label.Label, annotations []Annotation, code string, common bool) Method {
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

func (e Method) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *Method) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject = struct {
				Name        string                 `yaml:"Name"`
				Visibility  structs.VisibilityType `yaml:"Visibility"`
				Parameters  []MethodParameter      `yaml:"Parameters"`
				ReturnTypes []structs.DataType     `yaml:"Returns"`
				Hint        Message                `yaml:"Hint"`
				Description Message                `yaml:"Description"`
				Options     []option.Option        `yaml:"Options"`
				Labels      []label.Label          `yaml:"Labels"`
				Annotations []Annotation           `yaml:"Annotations"`
				Code        string                 `yaml:"Code"`
				Common      bool                   `yaml:"Common"`
			}{
				Common: true,
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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

	if _string.IsEmpty(string(s.Visibility)) {
		s.Visibility = structs.VisibilityTypeTypes.Public
	}

	return nil
}
