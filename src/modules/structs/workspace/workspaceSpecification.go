package workspace

import (
	"path/filepath"

	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/core/utils"
)

const (
	CodeBasePath  = "codebase"
	TemplatesPath = "templates"
	ResourcesPath = "resources"
)

type WorkspaceSpecification struct {
	applicationWorkspace.WorkspaceIdentifier
	Path string
}

func NewWorkspaceSpecification(id int, name, path string) WorkspaceSpecification {
	return WorkspaceSpecification{
		WorkspaceIdentifier: applicationWorkspace.NewWorkspaceIdentifier(id, name),
		Path:                path,
	}
}
func (e WorkspaceSpecification) Validate() error {
	if utils.IsEmpty(e.WorkspaceIdentifier.Name) {
		return &errors.ErrFieldRequired{FieldName: "WorkspaceIdentifier.Name"}
	}
	if utils.IsEmpty(e.Path) {
		return &errors.ErrFieldRequired{FieldName: "Path"}
	}
	return nil
}

func (s *WorkspaceSpecification) IsPathExists() bool {
	return !utils.IsEmpty(s.Path)
}

func (s WorkspaceSpecification) GetAbsolutePath() string {
	return filepath.Join(s.Path)
}

func (s WorkspaceSpecification) GetTemplatesFolder() string {
	return filepath.Join(s.GetAbsolutePath(), TemplatesPath)
}

func (s WorkspaceSpecification) GetCodeBaseFolder() string {
	return filepath.Join(s.GetAbsolutePath(), CodeBasePath)
}

func (s WorkspaceSpecification) GetResourcesFolder() string {
	return filepath.Join(s.GetAbsolutePath(), ResourcesPath)
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
