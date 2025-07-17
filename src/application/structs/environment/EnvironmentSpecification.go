package environment

import (
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type EnvironmentSpecification struct {
	EnvironmentIdentifier
}

func NewEnvironmentSpecification(id int, name string) EnvironmentSpecification {
	return EnvironmentSpecification{
		EnvironmentIdentifier: NewEnvironmentIdentifier(id, name),
	}
}
func (e EnvironmentSpecification) Validate() error {
	if _string.IsEmpty(e.EnvironmentIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "EnvironmentIdentifier.Name"}
	}
	return nil
}

func (s *EnvironmentSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		EnvironmentIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.EnvironmentIdentifier = tempIdentifierObject.EnvironmentIdentifier
	}

	return nil
}
