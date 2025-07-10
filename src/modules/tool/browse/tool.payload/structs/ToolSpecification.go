package structs

import (
	"fmt"

	applicationTool "parsdevkit.net/application/structs/tool"

	"gopkg.in/yaml.v3"
)

type ToolSpecification struct {
	applicationTool.ToolIdentifier
	Url string
}

func NewToolSpecification(id int, name, url string) ToolSpecification {
	return ToolSpecification{
		ToolIdentifier: applicationTool.NewToolIdentifier(id, name),
		Url:            url,
	}
}

func (s *ToolSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		applicationTool.ToolIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return fmt.Errorf("xxx: Tool Specification Çözümlenemedi\n%w", err)
	} else {
		s.ToolIdentifier = tempIdentifierObject.ToolIdentifier
	}

	var tempObject struct {
		Url string `yaml:"Url"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return fmt.Errorf("xxx: Tool Specification Url dönüştürme hatası oluştu\n%w", err)
		}
	} else {
		s.Url = tempObject.Url
	}

	return nil
}
