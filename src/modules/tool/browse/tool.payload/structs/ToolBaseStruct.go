package structs

import (
	"fmt"

	"parsdevkit.net/pkg/errors"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ToolBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications ToolSpecification
}

func (e ToolBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (s ToolBaseStruct) GetKey() string {
	return MODULE_KEY
}
func NewToolBaseStruct(header schemas.SchemaHeader, specifications ToolSpecification) ToolBaseStruct {
	return ToolBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e ToolBaseStruct) Validate() error {
	if _string.IsEmpty(e.Specifications.Url) {
		return &errors.ErrFieldRequired{FieldName: "Specifications.Url"}
	}
	return nil
}

func (s *ToolBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return fmt.Errorf("xxx: Tool Header Çözümlenemedi\n%w", err)
	} else {
		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications ToolSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return fmt.Errorf("xxx: Tool Specification dönüştürme hatası oluştu\n%w", err)
		}
		return fmt.Errorf("xxx: Tool Specification Çözümlenemedi\n%w", err)
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}
	return nil
}
