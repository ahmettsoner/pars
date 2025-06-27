package project

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/structs"
)

type Header struct {
	structs.Header
	Kind ProjectKind
}

func NewHeader(_type schemas.StructType, kind ProjectKind, name string, metadata structs.Metadata) Header {
	return Header{
		Header: structs.NewHeader(_type, name, metadata),
		Kind:   kind,
	}
}
func (e Header) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Type, v.Required),
		v.Field(&e.Kind, v.Required),
		v.Field(&e.Header.Name, v.Required),
	)
}

func (s *Header) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		structs.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	var tempProjectHeaderObject struct {
		Kind ProjectKind `yaml:"Kind"`
	}

	if err := unmarshal(&tempProjectHeaderObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Kind = tempProjectHeaderObject.Kind
	}

	return nil
}
