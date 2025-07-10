package structs

import (
	"fmt"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	_string "parsdevkit.net/pkg/utilities/string"
)

type GroupBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications GroupSpecification
}

func (e GroupBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (s GroupBaseStruct) GetKey() string {
	return MODULE_KEY
}
func NewGroupBaseStruct(header schemas.SchemaHeader, specifications GroupSpecification) GroupBaseStruct {
	return GroupBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e GroupBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *GroupBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return fmt.Errorf("xxx: Group Header Çözümlenemedi\n%w", err)
	} else {
		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications GroupSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return fmt.Errorf("xxx: Group Specification dönüştürme hatası oluştu\n%w", err)
		}
		return fmt.Errorf("xxx: Group Specification Çözümlenemedi\n%w", err)
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}
	return nil
}
