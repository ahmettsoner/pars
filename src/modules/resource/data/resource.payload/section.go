package data_resource_payload

import (
	"parsdevkit.net/application/models/class"
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/models/section"
)

type Section struct {
	section.Section
}

func NewSection(name string, labels []label.Label, options []option.Option, classes []class.Class) Section {
	return Section{
		Section: section.NewSection(name, labels, options, classes),
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
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
	}

	return nil
}
