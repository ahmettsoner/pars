package structs

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ToolBrowse struct {
	Url string
}

func NewToolBrowse(url string) ToolBrowse {
	return ToolBrowse{
		Url: url,
	}
}

func (s *ToolBrowse) UnmarshalYAML(unmarshal func(interface{}) error) error {

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
