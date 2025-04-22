package objectresource

import (
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Attribute struct {
	Variable
	Visibility    VisibilityType
	Group         AttributeGroup
	Encapsulation Encapsulation
	Properties    AttributeProperties
	Common        bool
}

func NewAttribute(name string, visibility VisibilityType, _type DataType, order int, group AttributeGroup, encapsulation Encapsulation, properties AttributeProperties, hint Message, description Message, options []option.Option, labels []label.Label, validation Validation, annotations []Annotation, common bool) Attribute {
	return Attribute{
		Variable:      NewVariable(name, _type, order, hint, description, options, labels, validation, annotations),
		Visibility:    visibility,
		Group:         group,
		Encapsulation: encapsulation,
		Properties:    properties,
		Common:        common,
	}
}

func (s *Attribute) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempVariableObject struct {
		Variable
	}

	if err := unmarshal(&tempVariableObject); err != nil {
		return err
	} else {
		s.Variable = tempVariableObject.Variable
	}

	var tempObject = struct {
		Visibility    VisibilityType      `yaml:"Visibility"`
		Group         interface{}         `yaml:"Group"`
		Encapsulation Encapsulation       `yaml:"Encapsulation"`
		Properties    AttributeProperties `yaml:"Properties"`
		Common        bool                `yaml:"Common"`
	}{
		Common: true,
	}

	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Visibility = tempObject.Visibility
				s.Encapsulation = tempObject.Encapsulation
				s.Properties = tempObject.Properties
				s.Common = tempObject.Common

				switch arguments := tempObject.Group.(type) {
				case string:
					s.Group = NewAttributeGroup(NewGroupIdentifier(arguments), 0, []option.Option(nil))
				case interface{}:
					var tempGroupObject struct {
						Group AttributeGroup `yaml:"Group"`
					}

					if err := unmarshal(&tempGroupObject); err != nil {
						if _, ok := err.(*yaml.TypeError); !ok {
							return err
						}
					}
					s.Group = tempGroupObject.Group
				}
			}

		}
	} else {
		s.Common = true
	}
	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Attribute.Name"}
	}

	if utils.IsEmpty(string(s.Visibility)) {
		s.Visibility = VisibilityTypeTypes.Public
	}

	return nil
}
