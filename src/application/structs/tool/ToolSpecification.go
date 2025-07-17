package tool

import (
	"fmt"

	"parsdevkit.net/application/schemas"
)

type ToolSpecification struct {
	ToolIdentifier
	schemas.SchemaSpecification
}

func NewToolSpecification(id int, name string) ToolSpecification {
	return ToolSpecification{
		ToolIdentifier: NewToolIdentifier(id, name),
	}
}

func (s *ToolSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		ToolIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return fmt.Errorf("xxx: Tool Specification Çözümlenemedi\n%w", err)
	} else {
		s.ToolIdentifier = tempIdentifierObject.ToolIdentifier
	}

	return nil
}
