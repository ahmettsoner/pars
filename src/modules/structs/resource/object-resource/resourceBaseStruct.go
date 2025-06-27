package objectresource

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs/resource"

	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
)

type ResourceBaseStruct struct {
	resource.Header
	Specifications ResourceSpecification
	Configurations ResourceConfiguration
}

func (e ResourceBaseStruct) GetHeader() schemas.SchemaHeader {
	return schemas.SchemaHeader{
		Type: e.Header.Type,
		Kind: string(e.Header.Kind),
		Name: e.Header.Name,
	}
}

func NewResourceBaseStruct(header resource.Header, specifications ResourceSpecification, configurations ResourceConfiguration) ResourceBaseStruct {
	return ResourceBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}
func (e ResourceBaseStruct) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Name, v.Required),
	)
}

func (s *ResourceBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		resource.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
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

	if utils.IsEmpty(string(s.Configurations.Generate)) {
		s.Configurations.Generate = ChangeTrackers.OnChange
	}

	return nil
}

func (s *ResourceBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Name, s.Specifications.Set)
}
