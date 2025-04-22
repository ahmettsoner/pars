package applicationproject

import (
	"encoding/json"
	"fmt"

	"parsdevkit.net/structs/project"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ProjectBaseStruct struct {
	project.Header
	Specifications ProjectSpecification
}

func NewProjectBaseStruct(header project.Header, specifications ProjectSpecification) ProjectBaseStruct {
	return ProjectBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}

func (s *ProjectBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		project.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications ProjectSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}

func (s *ProjectBaseStruct) UnmarshalJSON(data []byte) error {
	var tempHeaderObject struct {
		project.Header
	}

	if err := json.Unmarshal(data, &tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	var tempSpecificationObject struct {
		Specifications ProjectSpecification
	}

	// Unmarshal JSON into the temporary struct
	if err := json.Unmarshal(data, &tempSpecificationObject); err != nil {
		return err
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}

func (s *ProjectBaseStruct) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v-%v", s.Specifications.Group, s.Name, s.Specifications.Workspace)
}

func (s *ProjectBaseStruct) GetInformation() string {
	return fmt.Sprintf("%v (%v)", s.Name, s.Specifications.Set)
}

func (s *ProjectBaseStruct) GetFullName() string {
	if !utils.IsEmpty(s.Specifications.Group) {
		return fmt.Sprintf("%v/%v", s.Specifications.Group, s.Name)
	} else {
		return fmt.Sprintf("%v", s.Name)
	}
}

func (s *ProjectBaseStruct) GetFullInformation() string {
	if !utils.IsEmpty(s.Specifications.Set) {
		return fmt.Sprintf("%v (%v)", s.GetFullName(), s.Specifications.Set)
	}

	return fmt.Sprintf("%v (none-set)", s.GetFullName())

}
