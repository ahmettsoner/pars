package group

import (
	"fmt"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/structs"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type GroupBaseStruct struct {
	structs.Header
	Specifications GroupSpecification
}

func (e GroupBaseStruct) GetHeader() structs.SchemaHeader {
	return structs.SchemaHeader{
		Type: e.Header.Type,
		Name: e.Header.Name,
	}
}
func NewGroupBaseStruct(header structs.Header, specifications GroupSpecification) GroupBaseStruct {
	return GroupBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e GroupBaseStruct) Validate() error {
	if utils.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	if utils.IsEmpty(e.Specifications.Name) {
		return &errors.ErrFieldRequired{FieldName: "Specifications.Name"}
	}
	return nil
}

func (s *GroupBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		structs.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return fmt.Errorf("xxx: Group Header Çözümlenemedi\n%w", err)
	} else {
		s.Header = tempHeaderObject.Header
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
