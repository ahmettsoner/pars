package objectresource

import (
	"strings"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/structs"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
)

type Variable struct {
	Name        string
	Type        structs.DataType
	Order       int
	Hint        Message
	Description Message
	Options     []option.Option
	Labels      []label.Label
	Validation  Validation
	Annotations []Annotation
}

func NewVariable(name string, _type structs.DataType, order int, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation) Variable {
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

func (e Variable) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *Variable) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name        string           `yaml:"Name"`
				Type        structs.DataType `yaml:"Type"`
				Order       int              `yaml:"Order"`
				Hint        Message          `yaml:"Hint"`
				Description Message          `yaml:"Description"`
				Options     []option.Option  `yaml:"Options"`
				Labels      []label.Label    `yaml:"Labels"`
				Validation  Validation       `yaml:"Validation"`
				Annotations []Annotation     `yaml:"Annotations"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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

			_type, modifier := structs.DetectDataTypeModifier(_type)
			category := structs.DetectDataTypeCategory(_type)

			s.Name = name
			s.Type = structs.NewDataType(_type, structs.TypePackage{}, category, modifier, []structs.DataType(nil))
		} else {
			return &errors.InvalidFormatForPackageError{Value: value}
		}
	}

	if _string.IsEmpty(string(s.Type.Name)) {
		s.Type = structs.NewDataType(string(structs.ValueTypes.String), structs.TypePackage{}, structs.DataTypeCategories.Value, structs.ModifierTypes.Object, []structs.DataType(nil))
	}

	return nil
}
