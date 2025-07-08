package schemas

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

type SchemaHeader struct {
	Type     StructType `yaml:"Type"`
	Kind     string     `yaml:"Kind"`
	Name     string     `yaml:"Name"`
	Metadata Metadata   `yaml:"Metadata"`
}

func (l SchemaHeader) Key() string {
	return l.Name
}

func (l SchemaHeader) IsEqual(other SchemaHeader) bool {
	return reflect.DeepEqual(l, other)
}

func NewSchemaHeader(_type StructType, kind string, name string, metadata Metadata) SchemaHeader {
	return SchemaHeader{
		Type:     _type,
		Kind:     kind,
		Name:     name,
		Metadata: metadata,
	}
}

func (e SchemaHeader) Validate() error {
	// if _string.IsEmpty(e.Type) {
	// 	return &errors.ErrFieldRequired{FieldName: "Type"}
	// }
	// if _string.IsEmpty(e.Name) {
	// 	return &errors.ErrFieldRequired{FieldName: "Name"}
	// }

	return nil
}

func (s SchemaHeader) GetKey() string {
	var key string
	if s.Kind != "" {
		key = fmt.Sprintf("%s.%s", s.Type, s.Kind)
	} else {
		key = string(s.Type)
	}
	return key
}

func (s *SchemaHeader) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type     StructType `yaml:"Type"`
				Name     string     `yaml:"Name"`
				Kind     string     `yaml:"Kind"`
				Metadata Metadata   `yaml:"Metadata"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Type = tempObject.Type
			s.Kind = tempObject.Kind
			s.Name = tempObject.Name
			s.Metadata = tempObject.Metadata
		} else {
			return err
		}
	}

	return nil
}
