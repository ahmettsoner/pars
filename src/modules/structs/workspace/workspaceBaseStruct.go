package workspace

import (
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/structs"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type WorkspaceBaseStruct struct {
	structs.Header
	Specifications WorkspaceSpecification
}

func (e WorkspaceBaseStruct) GetHeader() schemas.SchemaHeader {
	return schemas.SchemaHeader{
		Type: e.Header.Type,
		Name: e.Header.Name,
	}
}

func NewWorkspaceBaseStruct(header structs.Header, specifications WorkspaceSpecification) WorkspaceBaseStruct {
	return WorkspaceBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e WorkspaceBaseStruct) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Name, v.Required),
	)
}

func (s *WorkspaceBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		structs.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications WorkspaceSpecification `yaml:"Specifications"`
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
