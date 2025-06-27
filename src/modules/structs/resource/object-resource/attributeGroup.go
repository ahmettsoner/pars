package objectresource

import (
	"strconv"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs/option"

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
func (e AttributeGroup) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Group.Name, v.Required),
	)
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
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

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

	return nil
}
