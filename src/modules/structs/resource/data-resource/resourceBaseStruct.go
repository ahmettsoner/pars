package dataresource

import (
	"fmt"

	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utilities"

	"parsdevkit.net/core/errors"
)

type ResourceBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications ResourceSpecification
	Configurations ResourceConfiguration
}

func (e ResourceBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}

func NewResourceBaseStruct(header schemas.SchemaHeader, specifications ResourceSpecification, configurations ResourceConfiguration) ResourceBaseStruct {
	return ResourceBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}
func (e ResourceBaseStruct) Validate() error {
	if utilities.IsEmpty(e.Header.Name) {
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

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications ResourceSpecification `yaml:"Specifications"`
		Configurations ResourceConfiguration `yaml:"Configurations"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Specifications = tempSpecificationObject.Specifications
		s.Configurations = tempSpecificationObject.Configurations
	}

	if utilities.IsEmpty(string(s.Configurations.Generate)) {
		s.Configurations.Generate = ChangeTrackers.OnChange
	}

	return nil
}

func (s *ResourceBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}
