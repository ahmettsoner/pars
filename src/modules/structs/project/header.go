package project

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs"
)

type Header struct {
	structs.Header
	Kind StructKind
}

func NewHeader(_type structs.StructType, kind StructKind, name string, metadata structs.Metadata) Header {
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
		Kind StructKind `yaml:"Kind"`
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
