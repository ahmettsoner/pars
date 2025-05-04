package commontask

import (
	"fmt"

	"parsdevkit.net/structs/task"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type TaskBaseStruct struct {
	task.Header
	Specifications TaskSpecification
	Configurations TaskConfiguration
}

func NewTaskBaseStruct(header task.Header, specifications TaskSpecification, configurations TaskConfiguration) TaskBaseStruct {
	return TaskBaseStruct{
		Header:         header,
		Specifications: specifications,
		Configurations: configurations,
	}
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

	return nil
}

func (s *TaskBaseStruct) GetFullInformation() string {
	return fmt.Sprintf("%v", s.Name)
}
