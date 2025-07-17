package project

import (
	"fmt"
	"reflect"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type Reference struct {
	Header         schemas.SchemaHeader
	Specifications ProjectSpecification
}

func (l Reference) Key() string {
	return l.Header.Name
}
func (s Reference) GetObjectKey() string {
	return fmt.Sprintf("%v-%v-%v", s.Specifications.Group, s.Header.Name, s.Specifications.Workspace)
}
func (s *Reference) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v-%v", s.Specifications.Group, s.Header.Name, s.Specifications.Workspace)
}
func (l Reference) IsEqual(other Reference) bool {
	return reflect.DeepEqual(l, other)
}

func NewReference(header schemas.SchemaHeader, specifications ProjectSpecification) Reference {
	return Reference{
		Header:         header,
		Specifications: specifications,
	}
}
func (e Reference) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *Reference) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	var tempSpecificationObject struct {
		Specifications ProjectSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {

		return err
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	return nil
}

func (s *Reference) GetInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}

func (s *Reference) GetFullName() string {
	if !_string.IsEmpty(s.Specifications.Group) {
		return fmt.Sprintf("%v/%v", s.Specifications.Group, s.Header.Name)
	} else {
		return fmt.Sprintf("%v", s.Header.Name)
	}
}

func (s *Reference) GetFullInformation() string {
	if !_string.IsEmpty(s.Specifications.Set) {
		return fmt.Sprintf("%v (%v)", s.GetFullName(), s.Specifications.Set)
	}

	return fmt.Sprintf("%v (none-set)", s.GetFullName())

}
