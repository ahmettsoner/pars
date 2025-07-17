package structs

import (
	"fmt"
	"reflect"

	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ProjectBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications ProjectSpecification
	Application    Application
}

func (e ProjectBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (e ProjectBaseStruct) GetSpecification() any {
	return e.Specifications
}
func (s ProjectBaseStruct) GetKey() string {
	return MODULE_KEY
}

func (l ProjectBaseStruct) Key() string {
	return l.Header.Name
}

func (l ProjectBaseStruct) IsEqual(other ProjectBaseStruct) bool {
	return reflect.DeepEqual(l, other)
}
func NewProjectBaseStruct(header schemas.SchemaHeader, specifications ProjectSpecification, application Application) ProjectBaseStruct {
	return ProjectBaseStruct{
		Header:         header,
		Specifications: specifications,
		Application:    application,
	}
}
func (e ProjectBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *ProjectBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
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

	var tempApplicationObject struct {
		Application Application `yaml:"Application"`
	}

	if err := unmarshal(&tempApplicationObject); err != nil {
		return err
	} else {
		s.Application = tempApplicationObject.Application
	}

	return nil
}

func (s *ProjectBaseStruct) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v-%v", s.Specifications.Group, s.Header.Name, s.Specifications.Workspace)
}

func (s *ProjectBaseStruct) GetInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}

func (s *ProjectBaseStruct) GetFullName() string {
	if !_string.IsEmpty(s.Specifications.Group) {
		return fmt.Sprintf("%v/%v", s.Specifications.Group, s.Header.Name)
	} else {
		return fmt.Sprintf("%v", s.Header.Name)
	}
}

func (s *ProjectBaseStruct) GetFullInformation() string {
	if !_string.IsEmpty(s.Specifications.Set) {
		return fmt.Sprintf("%v (%v)", s.GetFullName(), s.Specifications.Set)
	}

	return fmt.Sprintf("%v (none-set)", s.GetFullName())

}
