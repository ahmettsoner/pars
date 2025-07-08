package project

import (
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	"parsdevkit.net/pkg/utilities/file"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ProjectIdentifier struct {
	ID        int
	Name      string
	Group     string
	Workspace string
	Path      []string
}

func (l ProjectIdentifier) Key() string {
	return l.Name
}

func (l ProjectIdentifier) IsEqual(other ProjectIdentifier) bool {
	return reflect.DeepEqual(l, other)
}

func NewProjectIdentifier(id int, name string, path []string, group string, workspace string) ProjectIdentifier {
	return ProjectIdentifier{
		ID:        id,
		Name:      name,
		Path:      path,
		Group:     group,
		Workspace: workspace,
	}
}
func (e ProjectIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *ProjectIdentifier) IsPathExists() bool {
	return len(s.Path) > 0
}
func (s *ProjectIdentifier) GetProjectPath() string {

	folders := file.CombinePaths(s.Path)

	relativeFullPath := filepath.Join(folders...)

	return relativeFullPath
}
func (s *ProjectIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name      string `yaml:"Name"`
				Group     string `yaml:"Group"`
				Workspace string `yaml:"Workspace"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
			s.Group = tempObject.Group
			s.Workspace = tempObject.Workspace
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
