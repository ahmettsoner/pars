package group

import (
	"fmt"

	"parsdevkit.net/structs"

	"parsdevkit.net/core/utils"

	"gopkg.in/yaml.v3"
)

type GroupBaseStruct struct {
	structs.Header
	Specifications GroupSpecification
}

func NewGroupBaseStruct(header structs.Header, specifications GroupSpecification) GroupBaseStruct {
	return GroupBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}

func (s *GroupBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		structs.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return fmt.Errorf("xxx: Group Header Çözümlenemedi %w", err)
	} else {
		s.Header = tempHeaderObject.Header
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications GroupSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return fmt.Errorf("xxx: Group Specification dönüştürme hatası oluştu %w", err)
		}
		return fmt.Errorf("xxx: Group Specification Çözümlenemedi %w", err)
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	if utils.IsEmpty(s.Name) {
		return fmt.Errorf("xxx: Group Name alanı tanımlı değil")
	}

	return nil
}
