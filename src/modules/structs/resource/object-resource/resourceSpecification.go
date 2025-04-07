package objectresource

import (
	"fmt"
	"strings"

	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ResourceSpecification struct {
	ResourceIdentifier
	Path            string
	Set             string
	Package         []string
	WorkspaceObject workspace.WorkspaceSpecification
	Labels          []label.Label
	Layers          []Layer
	Dictionary      []Dictionary
	Groups          []Group
	Attributes      []Attribute
	Methods         []Method
}

func NewResourceSpecification(id int, name, workspace, path, set string, _package []string, labels []label.Label, layers []Layer, attributes []Attribute, methods []Method, workspaceObject workspace.WorkspaceSpecification) ResourceSpecification {
	return ResourceSpecification{
		ResourceIdentifier: NewResourceIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Path:               path,
		Set:                set,
		Package:            _package,
		Labels:             labels,
		Layers:             layers,
		Attributes:         attributes,
		Methods:            methods,
	}
}

func (s *ResourceSpecification) GetPackageString() string {
	return strings.Join(s.Package, "/")
}
func (s *ResourceSpecification) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *ResourceSpecification) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}
func (s *ResourceSpecification) IsPackageExists() bool {
	return len(s.Package) > 0
}
func (s *ResourceSpecification) IsPathExists() bool {
	return !utils.IsEmpty(s.Path)
}

func (s *ResourceSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		ResourceIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {
		s.ResourceIdentifier = tempIdentifierObject.ResourceIdentifier
	}
	var tempObject struct {
		Path       string        `yaml:"Path"`
		Set        string        `yaml:"Set"`
		Package    interface{}   `yaml:"Package"`
		Labels     []label.Label `yaml:"Labels"`
		Layers     []Layer       `yaml:"Layers"`
		Attributes []Attribute   `yaml:"Attributes"`
		Methods    []Method      `yaml:"Methods"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Path = tempObject.Path
		s.Set = tempObject.Set
		s.Labels = tempObject.Labels
		s.Layers = tempObject.Layers

		s.Attributes = tempObject.Attributes
		s.Methods = tempObject.Methods

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
		return &errors.ErrFieldRequired{FieldName: "Resource.Name"}
	}

	if utils.IsEmpty(s.Set) {
		return &errors.ErrFieldRequired{FieldName: "Resource.Set"}
	}

	return nil
}
