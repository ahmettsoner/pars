package objectresource

import (
	"parsdevkit.net/application/models/class"
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/core/errors"
	_string "parsdevkit.net/core/utilities/string"
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
func (e Section) Validate() error {
	if _string.IsEmpty(e.Section.Name) {
		return &errors.ErrFieldRequired{FieldName: "Section.Name"}
	}
	return nil
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
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Attributes = tempObject.Attributes
		s.Methods = tempObject.Methods
	}
	return nil
}
