package structs

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/core/schemas"

	"gopkg.in/yaml.v3"
)

type Header struct {
	Type     schemas.StructType
	Name     string
	Metadata Metadata
}

func NewHeader(_type schemas.StructType, name string, metadata Metadata) Header {
	return Header{
		Type:     _type,
		Name:     name,
		Metadata: metadata,
	}
}

func (e Header) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Type, v.Required),
		v.Field(&e.Name, v.Required),
	)
}

func (s *Header) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type     schemas.StructType `yaml:"Type"`
				Name     string             `yaml:"Name"`
				Metadata Metadata           `yaml:"Metadata"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Type = tempObject.Type
			s.Name = tempObject.Name
			s.Metadata = tempObject.Metadata
		} else {
			return err
		}
	}

	return nil
}
