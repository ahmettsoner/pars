package tool

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ToolIdentifier struct {
	ID   int
	Name string
}

func NewToolIdentifier(id int, name string) ToolIdentifier {
	return ToolIdentifier{
		ID:   id,
		Name: name,
	}
}
func NewToolIdentifier_Empty(name string) ToolIdentifier {
	return ToolIdentifier{
		ID:   0,
		Name: name,
	}
}
func (e ToolIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}
func (s *ToolIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return fmt.Errorf("xxx: Tool Identifier Çözümlenemedi\n%w", err)
			}

			s.Name = tempObject.Name
		} else {
			return fmt.Errorf("xxx: Tool Identifier dönüştürme hatası oluştu\n%w", err)
		}
	} else {
		s.Name = value
	}

	return nil
}
