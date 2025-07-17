package workspace

import (
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type WorkspaceSpecification struct {
	WorkspaceIdentifier
	schemas.SchemaSpecification
}

func NewWorkspaceSpecification(id int, name, path string) WorkspaceSpecification {
	return WorkspaceSpecification{
		WorkspaceIdentifier: NewWorkspaceIdentifier(id, name, path),
	}
}
func (e WorkspaceSpecification) Validate() error {
	if _string.IsEmpty(e.WorkspaceIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "WorkspaceIdentifier.Name"}
	}
	if _string.IsEmpty(e.Path) {
		return &errors.ErrFieldRequired{FieldName: "Path"}
	}
	return nil
}

func (s *WorkspaceSpecification) IsPathExists() bool {
	return !_string.IsEmpty(s.Path)
}

func (s *WorkspaceSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		WorkspaceIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.WorkspaceIdentifier = tempIdentifierObject.WorkspaceIdentifier
	}

	var tempObject struct {
		Path string `yaml:"Path"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Path = tempObject.Path
	}

	return nil
}
