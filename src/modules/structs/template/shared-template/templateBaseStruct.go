package sharedtemplate

import (
	"fmt"

	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
)

type TemplateBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications TemplateSpecification
	Configurations TemplateConfiguration
}

func (e TemplateBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}

func NewTemplateBaseStruct(header schemas.SchemaHeader, specifications TemplateSpecification, configurations TemplateConfiguration) TemplateBaseStruct {
	return TemplateBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}
func (e TemplateBaseStruct) Validate() error {
	if utils.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *TemplateBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications TemplateSpecification `yaml:"Specifications"`
		Configurations TemplateConfiguration `yaml:"Configurations"`
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

	return nil
}

func (s *TemplateBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v", s.Header.Name)
}
