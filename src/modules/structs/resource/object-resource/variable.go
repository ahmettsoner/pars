package objectresource

import (
	"strings"

	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Variable struct {
	Name        string
	Type        DataType
	Order       int
	Hint        Message
	Description Message
	Options     []option.Option
	Labels      []label.Label
	Validation  Validation
	Annotations []Annotation
}

func NewVariable(name string, _type DataType, order int, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation) Variable {
	return Variable{
		Name:        name,
		Type:        _type,
		Order:       order,
		Hint:        hint,
		Description: description,
		Options:     options,
		Labels:      labels,
		Validation:  validation,
		Annotations: annotations,
	}
}
func (s *Variable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name        string          `yaml:"Name"`
				Type        DataType        `yaml:"Type"`
				Order       int             `yaml:"Order"`
				Hint        Message         `yaml:"Hint"`
				Description Message         `yaml:"Description"`
				Options     []option.Option `yaml:"Options"`
				Labels      []label.Label   `yaml:"Labels"`
				Validation  Validation      `yaml:"Validation"`
				Annotations []Annotation    `yaml:"Annotations"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
				s.Type = tempObject.Type
				s.Order = tempObject.Order
				s.Hint = tempObject.Hint
				s.Description = tempObject.Description
				s.Options = tempObject.Options
				s.Labels = tempObject.Labels
				s.Validation = tempObject.Validation
				s.Annotations = tempObject.Annotations
			}
		} else {
			return err
		}
	} else {
		var parts []string = strings.Split(value, " ")
		if len(parts) == 1 {
			s.Name = value
		} else if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			_type := strings.TrimSpace(parts[1])

			_type, modifier := DetectDataTypeModifier(_type)
			category := DetectDataTypeCategory(_type)

			s.Name = name
			s.Type = NewDataType(_type, TypePackage{}, category, modifier, []DataType(nil))
		} else {
			return &errors.InvalidFormatForPackageError{Value: value}
		}
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Variable.Name"}
	}

	if utils.IsEmpty(string(s.Type.Name)) {
		s.Type = NewDataType(string(ValueTypes.String), TypePackage{}, DataTypeCategories.Value, ModifierTypes.Object, []DataType(nil))
	}

	return nil
}
