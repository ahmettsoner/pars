package objectresource

import (
	"fmt"

	"parsdevkit.net/structs/resource"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ResourceBaseStruct struct {
	resource.Header
	Specifications ResourceSpecification
	Configurations ResourceConfiguration
}

func NewResourceBaseStruct(header resource.Header, specifications ResourceSpecification, configurations ResourceConfiguration) ResourceBaseStruct {
	return ResourceBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
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

	if utils.IsEmpty(string(s.Configurations.Generate)) {
		s.Configurations.Generate = ChangeTrackers.OnChange
	}

	return nil
}

func (s *ResourceBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v (%v)", s.Name, s.Specifications.Set)
}
