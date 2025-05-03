package dataresource

import (
	"parsdevkit.net/structs/class"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"parsdevkit.net/structs/section"

	"gopkg.in/yaml.v3"
)

type Section struct {
	section.Section
}

func NewSection(name string, labels []label.Label, options []option.Option, classes []class.Class) Section {
	return Section{
		Section: section.NewSection(name, labels, options, classes),
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
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Section.Name"}
	}

	return nil
}
