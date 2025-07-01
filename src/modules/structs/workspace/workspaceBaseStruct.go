package workspace

import (
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utilities"

	"parsdevkit.net/core/errors"
)

type WorkspaceBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications WorkspaceSpecification
}

func (e WorkspaceBaseStruct) GetHeader() schemas.SchemaHeader {
	return schemas.SchemaHeader{
		Type: e.Header.Type,
		Name: e.Header.Name,
	}
}

func NewWorkspaceBaseStruct(header schemas.SchemaHeader, specifications WorkspaceSpecification) WorkspaceBaseStruct {
	return WorkspaceBaseStruct{
		Header:         header,
		Specifications: specifications,
	}
}
func (e WorkspaceBaseStruct) Validate() error {
	if utilities.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}

func (s *WorkspaceBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
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
