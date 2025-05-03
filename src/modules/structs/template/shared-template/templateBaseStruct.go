package sharedtemplate

import (
	"fmt"

	"parsdevkit.net/structs/template"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type TemplateBaseStruct struct {
	template.Header
	Specifications TemplateSpecification
	Configurations TemplateConfiguration
}

func NewTemplateBaseStruct(header template.Header, specifications TemplateSpecification, configurations TemplateConfiguration) TemplateBaseStruct {
	return TemplateBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}

func (s *TemplateBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		template.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications TemplateSpecification `yaml:"Specifications"`
		Configurations TemplateConfiguration `yaml:"Configurations"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Specifications = tempSpecificationObject.Specifications
		s.Configurations = tempSpecificationObject.Configurations
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}

func (s *TemplateBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v", s.Name)
}
