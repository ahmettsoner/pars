package applicationproject

import (
	"encoding/json"
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"

	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
)

type ProjectBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications ProjectSpecification
}

func (e ProjectBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}

func NewProjectBaseStruct(header schemas.SchemaHeader, specifications ProjectSpecification) ProjectBaseStruct {
	return ProjectBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e ProjectBaseStruct) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Name, v.Required),
	)
}

func (s *ProjectBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications ProjectSpecification `yaml:"Specifications"`
	}

	if err := unmarshal(&tempSpecificationObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err
	} else {
		s.Specifications = tempSpecificationObject.Specifications
	}

	return nil
}

func (s *ProjectBaseStruct) UnmarshalJSON(data []byte) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := json.Unmarshal(data, &tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
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

	return nil
}

func (s *ProjectBaseStruct) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v-%v", s.Specifications.Group, s.Header.Name, s.Specifications.Workspace)
}

func (s *ProjectBaseStruct) GetInformation() string {
	return fmt.Sprintf("%v (%v)", s.Header.Name, s.Specifications.Set)
}

func (s *ProjectBaseStruct) GetFullName() string {
	if !utils.IsEmpty(s.Specifications.Group) {
		return fmt.Sprintf("%v/%v", s.Specifications.Group, s.Header.Name)
	} else {
		return fmt.Sprintf("%v", s.Header.Name)
	}
}

func (s *ProjectBaseStruct) GetFullInformation() string {
	if !utils.IsEmpty(s.Specifications.Set) {
		return fmt.Sprintf("%v (%v)", s.GetFullName(), s.Specifications.Set)
	}

	return fmt.Sprintf("%v (none-set)", s.GetFullName())

}
