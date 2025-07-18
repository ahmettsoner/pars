package structs

import (
	"fmt"

	"parsdevkit.net/application/schemas"
	applicationResource "parsdevkit.net/application/structs/resource"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ResourceBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications applicationResource.ResourceSpecification
	Object         ResourceObject
	Configurations ResourceConfiguration
}

func (e ResourceBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (e ResourceBaseStruct) GetSpecification() any {
	return e.Specifications
}
func (s ResourceBaseStruct) GetKey() string {
	return MODULE_KEY
}

func NewResourceBaseStruct(header schemas.SchemaHeader, specifications applicationResource.ResourceSpecification, object ResourceObject, configurations ResourceConfiguration) ResourceBaseStruct {
	return ResourceBaseStruct{
		Header:         header,
		Specifications: specifications,
		Object:         object,
		Configurations: configurations,
	}
}
func (e ResourceBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *ResourceBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	var tempSpecificationObject struct {
		Specifications applicationResource.ResourceSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		return err

	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	var tempObjectObject struct {
		Object ResourceObject `yaml:"Object"`
	}

	if err := unmarshal(&tempObjectObject); err != nil {
		return err

	} else {
		s.Object = tempObjectObject.Object
	}

	var tempConfigurationObject struct {
		Configurations ResourceConfiguration `yaml:"Configurations"`
	}

	if err := unmarshal(&tempConfigurationObject); err != nil {
		return err

	} else {
		s.Configurations = tempConfigurationObject.Configurations
	}

	return nil
}

func (s *ResourceBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}
