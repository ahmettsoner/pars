package structs

import (
	"parsdevkit.net/application/schemas"

	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type EnvironmentBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications EnvironmentSpecification
}

func (e EnvironmentBaseStruct) GetHeader() schemas.SchemaHeader {
	return schemas.SchemaHeader{
		Type: e.Header.Type,
		Name: e.Header.Name,
	}
}
func (s EnvironmentBaseStruct) GetKey() string {
	return "Environment"
}

func NewEnvironmentBaseStruct(header schemas.SchemaHeader, specifications EnvironmentSpecification) EnvironmentBaseStruct {
	return EnvironmentBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e EnvironmentBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *EnvironmentBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications EnvironmentSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	return nil
}
