package task

import (
	"parsdevkit.net/structs"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
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
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Kind = tempProjectHeaderObject.Kind
	}

	if utils.IsEmpty(string(s.Type)) {
		return &errors.ErrFieldRequired{FieldName: "Type"}
	}
	if utils.IsEmpty(string(s.Kind)) {
		return &errors.ErrFieldRequired{FieldName: "Kind"}
	}
	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
