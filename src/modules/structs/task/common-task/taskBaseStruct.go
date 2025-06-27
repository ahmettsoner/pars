package commontask

import (
	"fmt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/structs"
	"parsdevkit.net/structs/task"
)

type TaskBaseStruct struct {
	task.Header
	Specifications TaskSpecification
	Configurations TaskConfiguration
}

func (e TaskBaseStruct) GetHeader() structs.SchemaHeader {
	return structs.SchemaHeader{
		Type: e.Header.Type,
		Kind: string(e.Header.Kind),
		Name: e.Header.Name,
	}
}
func NewTaskBaseStruct(header task.Header, specifications TaskSpecification, configurations TaskConfiguration) TaskBaseStruct {
	return TaskBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
}

func (e TaskBaseStruct) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Header.Name, v.Required),
	)
}
func (s *TaskBaseStruct) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tempHeaderObject struct {
		task.Header
	}

	if err := unmarshal(&tempHeaderObject); err != nil {
		return err
	} else {

		s.Header = tempHeaderObject.Header
	}

	//TODO: Specification ve Header 2 işlemde alındı düzeltilmeli, aşağıda ki block Specification bölümünü yeniden almak için geçici olarak kullanıldı
	var tempSpecificationObject struct {
		Specifications TaskSpecification `yaml:"Specifications"`
		Configurations TaskConfiguration `yaml:"Configurations"`
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
	return fmt.Sprintf("%v", s.Name)
}
