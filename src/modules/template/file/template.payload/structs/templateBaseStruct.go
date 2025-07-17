package structs

import (
	"fmt"

	"parsdevkit.net/application/schemas"
	applicationTemplate "parsdevkit.net/application/structs/template"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type TemplateBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications applicationTemplate.TemplateSpecification
	Configurations TemplateConfiguration
}

func (e TemplateBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (e TemplateBaseStruct) GetSpecification() any {
	return e.Specifications
}
func (s TemplateBaseStruct) GetKey() string {
	return MODULE_KEY
}

func NewTemplateBaseStruct(header schemas.SchemaHeader, specifications applicationTemplate.TemplateSpecification, configurations TemplateConfiguration) TemplateBaseStruct {
	return TemplateBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}
func (e TemplateBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
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
		Specifications applicationTemplate.TemplateSpecification `yaml:"Specifications"`
		Configurations TemplateConfiguration                     `yaml:"Configurations"`
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

	if _string.IsEmpty(string(s.Configurations.Generate)) {
		s.Configurations.Generate = ChangeTrackers.OnChange
	}

	return nil
}

func (s *TemplateBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}
