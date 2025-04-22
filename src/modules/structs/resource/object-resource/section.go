package objectresource

import (
	"parsdevkit.net/structs/class"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"
	"parsdevkit.net/structs/section"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Section struct {
	section.Section
	Attributes []string
	Methods    []string
}

func NewSection(name string, attributes []string, methods []string, labels []label.Label, options []option.Option, classes []class.Class) Section {
	return Section{
		Section:    section.NewSection(name, labels, options, classes),
		Attributes: attributes,
		Methods:    methods,
	}
}

func (s *Section) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempSectionObject struct {
		section.Section
	}

	if err := unmarshal(&tempSectionObject); err != nil {
		return err
	} else {

		s.Section = tempSectionObject.Section
	}

	var tempObject struct {
		Attributes []string `yaml:"Attributes"`
		Methods    []string `yaml:"Methods"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Attributes = tempObject.Attributes
		s.Methods = tempObject.Methods
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Section.Name"}
	}

	return nil
}
