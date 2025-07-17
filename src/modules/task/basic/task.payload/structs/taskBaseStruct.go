package structs

import (
	"fmt"

	"parsdevkit.net/application/schemas"
	applicationTask "parsdevkit.net/application/structs/task"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type TaskBaseStruct struct {
	Header         schemas.SchemaHeader
	Specifications applicationTask.TaskSpecification
	Configurations TaskConfiguration
}

func (e TaskBaseStruct) GetHeader() schemas.SchemaHeader {
	return e.Header
}
func (e TaskBaseStruct) GetSpecification() any {
	return e.Specifications
}
func (s TaskBaseStruct) GetKey() string {
	return MODULE_KEY
}
func NewTaskBaseStruct(header schemas.SchemaHeader, specifications applicationTask.TaskSpecification, configurations TaskConfiguration) TaskBaseStruct {
	return TaskBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}

func (e TaskBaseStruct) Validate() error {
	if _string.IsEmpty(e.Header.Name) {
		return &errors.ErrFieldRequired{FieldName: "Header.Name"}
	}
	return nil
}
func (s *TaskBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject schemas.SchemaHeader

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications applicationTask.TaskSpecification `yaml:"Specifications"`
		Configurations TaskConfiguration                 `yaml:"Configurations"`
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

func (s *TaskBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v", s.Header.Name)
}
