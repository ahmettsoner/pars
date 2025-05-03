package section

import (
	"parsdevkit.net/structs/class"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Section struct {
	SectionIdentifier
	Labels  []label.Label
	Options []option.Option
	Classes []class.Class
}

func NewSection(name string, labels []label.Label, options []option.Option, classes []class.Class) Section {
	return Section{
		SectionIdentifier: NewSectionIdentifier(name),
		Labels:            labels,
		Options:           options,
		Classes:           classes,
	}
}

func (s *Section) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {

			var tempIdentifierObject struct {
				SectionIdentifier
			}

			if err := unmarshal(&tempIdentifierObject); err != nil {
				return err
			} else {

				s.SectionIdentifier = tempIdentifierObject.SectionIdentifier
			}

			var tempObject struct {
				Labels  []label.Label   `yaml:"Labels"`
				Options []option.Option `yaml:"Options"`
				Classes []class.Class   `yaml:"Classes"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Labels = tempObject.Labels
				s.Options = tempObject.Options
				s.Classes = tempObject.Classes
			}
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	if utils.IsEmpty(string(s.Name)) {
		return &errors.ErrFieldRequired{FieldName: "Section.Name"}
	}

	return nil
}
