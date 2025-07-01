package dataresource

import (
	"parsdevkit.net/application/models/option"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utilities"
)

type Group struct {
	GroupIdentifier
	Title   Message
	Order   int
	Options []option.Option
}

func NewGroup(name string, title Message, order int, options []option.Option) Group {
	return Group{
		GroupIdentifier: NewGroupIdentifier(name),
		Title:           title,
		Order:           order,
		Options:         options,
	}
}

func (e Group) Validate() error {
	if utilities.IsEmpty(e.GroupIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "GroupIdentifier.Name"}
	}
	return nil
}
func (s *Group) IsNameExists() bool {
	return !utilities.IsEmpty(s.Name)
}

func (s *Group) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempIdentifierObject struct {
		GroupIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.GroupIdentifier = tempIdentifierObject.GroupIdentifier
	}

	var tempObject struct {
		Title   Message         `yaml:"Title"`
		Order   int             `yaml:"Order"`
		Options []option.Option `yaml:"Options"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Title = tempObject.Title
		s.Order = tempObject.Order
		s.Options = tempObject.Options

	}

	return nil
}
