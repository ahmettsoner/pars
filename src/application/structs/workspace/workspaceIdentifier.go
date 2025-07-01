package workspace

import (
	"path/filepath"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/core/errors"
	_string "parsdevkit.net/core/utilities/string"
)

const (
	CodeBasePath  = "codebase"
	TemplatesPath = "templates"
	ResourcesPath = "resources"
)

type WorkspaceIdentifier struct {
	ID   int
	Name string
	Path string
}

func NewWorkspaceIdentifier(id int, name, path string) WorkspaceIdentifier {
	return WorkspaceIdentifier{
		ID:   id,
		Name: name,
		Path: path,
	}
}
func (e WorkspaceIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *WorkspaceIdentifier) IsPathExists() bool {
	return !_string.IsEmpty(s.Path)
}
func (s WorkspaceIdentifier) GetAbsolutePath() string {
	return filepath.Join(s.Path)
}
func (s WorkspaceIdentifier) GetTemplatesFolder() string {
	return filepath.Join(s.GetAbsolutePath(), TemplatesPath)
}

func (s WorkspaceIdentifier) GetCodeBaseFolder() string {
	return filepath.Join(s.GetAbsolutePath(), CodeBasePath)
}

func (s WorkspaceIdentifier) GetResourcesFolder() string {
	return filepath.Join(s.GetAbsolutePath(), ResourcesPath)
}

func (s *WorkspaceIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
