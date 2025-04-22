package objectresource

import (
	"reflect"
	"strconv"

	"parsdevkit.net/structs/option"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type AttributeGroup struct {
	Group   GroupIdentifier
	Order   int
	Options []option.Option
}

func NewAttributeGroup(group GroupIdentifier, order int, options []option.Option) AttributeGroup {
	return AttributeGroup{
		Group:   group,
		Order:   order,
		Options: options,
	}
}

func (s *AttributeGroup) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Group   GroupIdentifier `yaml:"RefGroup"`
				Order   int             `yaml:"Order"`
				Options []option.Option `yaml:"Options"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Group = tempObject.Group
				s.Order = tempObject.Order
				s.Options = tempObject.Options
			}
		} else {
			if intValue, err := strconv.Atoi(value); err != nil {
				return err
			} else {
				s.Order = intValue
			}
		}
	}

	if reflect.DeepEqual(s.Group, GroupIdentifier{}) {
		return &errors.ErrFieldRequired{FieldName: "Group"}
	}

	return nil
}
