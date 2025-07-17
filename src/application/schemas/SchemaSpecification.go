package schemas

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

type SchemaSpecification struct {
	Name string `yaml:"Name"`
}

func (l SchemaSpecification) Key() string {
	return l.Name
}

func (l SchemaSpecification) IsEqual(other SchemaSpecification) bool {
	return reflect.DeepEqual(l, other)
}

func NewSchemaSpecification(name string) SchemaSpecification {
	return SchemaSpecification{
		Name: name,
	}
}

func (e SchemaSpecification) Validate() error {
	// if _string.IsEmpty(e.Type) {
	// 	return &errors.ErrFieldRequired{FieldName: "Type"}
	// }
	// if _string.IsEmpty(e.Name) {
	// 	return &errors.ErrFieldRequired{FieldName: "Name"}
	// }

	return nil
}

func (s *SchemaSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
		} else {
			return err
		}
	}

	return nil
}
