package codetemplate

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs/template"

	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
)

type TemplateBaseStruct struct {
	template.Header
	Specifications TemplateSpecification
	Configurations TemplateConfiguration
}

func (e TemplateBaseStruct) GetHeader() schemas.SchemaHeader {
	return schemas.SchemaHeader{
		Type: e.Header.Type,
		Kind: string(e.Header.Kind),
		Name: e.Header.Name,
	}
}
func NewTemplateBaseStruct(header template.Header, specifications TemplateSpecification, configurations TemplateConfiguration) TemplateBaseStruct {
	return TemplateBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}
func (e TemplateBaseStruct) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Name, v.Required),
	)
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
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Specifications = tempSpecificationObject.Specifications
		s.Configurations = tempSpecificationObject.Configurations
	}

	if utils.IsEmpty(string(s.Configurations.Generate)) {
		s.Configurations.Generate = ChangeTrackers.OnChange
	}

	return nil
}
func (s *TemplateBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Name, s.Specifications.Set)
}
