package structs

import (
	v "github.com/go-ozzo/ozzo-validation/v4"

	"gopkg.in/yaml.v3"
)

type Schema interface {
	Validate() error
	// PrintInfo()
	GetHeader() SchemaHeader
}

type SchemaHeader struct {
	Type StructType `yaml:"Type"`
	Kind string     `yaml:"Kind"`
	Name string     `yaml:"Name"`
}

type Header struct {
	Type     StructType
	Name     string
	Metadata Metadata
}

func NewHeader(_type StructType, name string, metadata Metadata) Header {
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
				Type     StructType `yaml:"Type"`
				Name     string     `yaml:"Name"`
				Metadata Metadata   `yaml:"Metadata"`
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
