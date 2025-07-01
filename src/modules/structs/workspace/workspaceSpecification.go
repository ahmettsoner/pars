package workspace

import (
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/errors"
	_string "parsdevkit.net/core/utilities/string"
)

type WorkspaceSpecification struct {
	applicationWorkspace.WorkspaceIdentifier
}

func NewWorkspaceSpecification(id int, name, path string) WorkspaceSpecification {
	return WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.NewWorkspaceIdentifier(id, name, path),
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
		applicationWorkspace.WorkspaceIdentifier
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
