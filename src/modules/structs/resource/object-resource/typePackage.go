package objectresource

import (
	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
)

type TypePackage struct {
	Name  string
	Alias string
}

func NewTypePackage(name string, alias string) TypePackage {
	return TypePackage{
		Name:  name,
		Alias: alias,
	}
}

func NewTypePackageOnly(name string) TypePackage {
	return TypePackage{
		Name: name,
	}
}

func (e TypePackage) Validate() error {
	if utils.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *TypePackage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err == nil {
		s.Name = value
	} else {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name  string `yaml:"Name"`
				Alias string `yaml:"Alias"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

			} else {
				s.Name = tempObject.Name
				s.Alias = tempObject.Alias
			}
		} else {
			return err
		}
	}

	return nil
}
