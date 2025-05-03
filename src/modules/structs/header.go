package structs

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

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

	if utils.IsEmpty(string(s.Type)) {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}
	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
