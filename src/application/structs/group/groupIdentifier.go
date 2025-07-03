package group

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type GroupIdentifier struct {
	ID      int
	Name    string
	Path    string
	Package []string
}

func NewGroupIdentifier(id int, name, path string, _package []string) GroupIdentifier {
	return GroupIdentifier{
		ID:      id,
		Name:    name,
		Path:    path,
		Package: _package,
	}
}
func NewGroupIdentifier_Empty(name string) GroupIdentifier {
	return GroupIdentifier{
		ID:      0,
		Name:    name,
		Path:    name,
		Package: []string{},
	}
}
func (e GroupIdentifier) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	return nil
}

func (s *GroupIdentifier) IsPathExists() bool {
	return !_string.IsEmpty(s.Path)
}
func (s *GroupIdentifier) GetRelativeGroupPath() string {
	return s.Path
}

func (s *GroupIdentifier) GetPackageString() string {
	return strings.Join(s.Package, "/")
}
func (s *GroupIdentifier) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *GroupIdentifier) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}
func (s *GroupIdentifier) IsPackageExists() bool {
	return len(s.Package) > 0
}
func (s *GroupIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return fmt.Errorf("xxx: Group Identifier Çözümlenemedi\n%w", err)
			}

			s.Name = tempObject.Name
		} else {
			return fmt.Errorf("xxx: Group Identifier dönüştürme hatası oluştu\n%w", err)
		}
	} else {
		s.Name = value
	}

	return nil
}
