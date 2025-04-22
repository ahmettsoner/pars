package workspace

import (
	"path/filepath"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

const (
	CodeBasePath  = "codebase"
	TemplatesPath = "templates"
	ResourcesPath = "resources"
)

type WorkspaceSpecification struct {
	WorkspaceIdentifier
	Path string
}

func NewWorkspaceSpecification(id int, name, path string) WorkspaceSpecification {
	return WorkspaceSpecification{
		WorkspaceIdentifier: NewWorkspaceIdentifier(id, name),
		Path:                path,
	}
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
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Path = tempObject.Path
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	if utils.IsEmpty(s.Path) {
		return &errors.ErrFieldRequired{FieldName: "Path"}
	}

	return nil
}
