package objectresource

import (
	"gopkg.in/yaml.v3"
)

type AttributeProperties struct {
	Key      bool
	Required bool
	ReadOnly bool
	Unique   bool
	Default  string
	Format   string
}

func NewAttributeProperties(key, required, readOnly, unique bool, _default, format string) AttributeProperties {
	return AttributeProperties{
		Key:      key,
		ReadOnly: readOnly,
		Required: required,
		Unique:   unique,
		Default:  _default,
		Format:   format,
	}
}

func (s *AttributeProperties) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Key      bool   `yaml:"Key"`
				Required bool   `yaml:"Required"`
				ReadOnly bool   `yaml:"ReadOnly"`
				Unique   bool   `yaml:"Unique"`
				Default  string `yaml:"Default"`
				Format   string `yaml:"Format"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Key = tempObject.Key
				s.Required = tempObject.Required
				s.ReadOnly = tempObject.ReadOnly
				s.Unique = tempObject.Unique
				s.Default = tempObject.Default
				s.Format = tempObject.Format
			}

		}
	}

	return nil
}
