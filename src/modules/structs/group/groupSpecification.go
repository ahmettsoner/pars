package group

import (
	"fmt"
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type GroupSpecification struct {
	GroupIdentifier
	Path    string
	Package []string
}

func NewGroupSpecification(id int, name, path string, _package []string) GroupSpecification {
	return GroupSpecification{
		GroupIdentifier: NewGroupIdentifier(id, name),
		Path:            path,
		Package:         _package,
	}
}

func NewGroupSpecification_Empty(name string) GroupSpecification {
	return GroupSpecification{
		GroupIdentifier: NewGroupIdentifier(0, name),
	}
}

func (s *GroupSpecification) GetPackageString() string {
	return strings.Join(s.Package, "/")
}
func (s *GroupSpecification) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *GroupSpecification) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}
func (s *GroupSpecification) IsPackageExists() bool {
	return len(s.Package) > 0
}
func (s *GroupSpecification) IsPathExists() bool {
	return !utils.IsEmpty(s.Path)
}
func (s *GroupSpecification) GetRelativeGroupPath() string {
	return s.Path
}

func (s *GroupSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		GroupIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.GroupIdentifier = tempIdentifierObject.GroupIdentifier
	}

	var tempObject struct {
		Path    string      `yaml:"Path"`
		Package interface{} `yaml:"Package"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Path = tempObject.Path

		switch packages := tempObject.Package.(type) {
		case string:
			s.SetPackageFromString(packages)
		case []interface{}:
			for _, _package := range packages {
				s.AppendPackage(fmt.Sprint(_package))
			}
		}
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	if len(s.Package) == 0 {
		s.Package = []string{s.Name}
	}
	if utils.IsEmpty(s.Path) {
		s.Path = s.Name
	}

	return nil
}
