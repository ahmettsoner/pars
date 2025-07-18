package resource

import (
	"fmt"
	"strings"

	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs"
	_string "parsdevkit.net/pkg/utilities/string"

	layerPkg "parsdevkit.net/application/models/layer"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/pkg/errors"
)

type ResourceSpecification struct {
	ResourceIdentifier
	schemas.SchemaSpecification
	Path            string
	Set             string
	Package         []string
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	Labels          []label.Label
	Layers          []layerPkg.Layer
	Generate        structs.ChangeTracker
}

func NewResourceSpecification(id int, name, workspace, path, set string, _package []string, labels []label.Label, layers []layerPkg.Layer, workspaceObject applicationWorkspace.WorkspaceIdentifier, generate structs.ChangeTracker) ResourceSpecification {
	return ResourceSpecification{
		ResourceIdentifier: NewResourceIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Path:               path,
		Set:                set,
		Package:            _package,
		Labels:             labels,
		Layers:             layers,
		Generate:           generate,
	}
}

func (e ResourceSpecification) Validate() error {
	if _string.IsEmpty(e.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}
	if _string.IsEmpty(e.Set) {
		return &errors.ErrFieldRequired{FieldName: "Set"}
	}
	return nil
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
	return !_string.IsEmpty(s.Path)
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
		Path     string                `yaml:"Path"`
		Set      string                `yaml:"Set"`
		Package  interface{}           `yaml:"Package"`
		Labels   []label.Label         `yaml:"Labels"`
		Layers   []layerPkg.Layer      `yaml:"Layers"`
		Generate structs.ChangeTracker `yaml:"Generate"`
	}

	if err := unmarshal(&tempObject); err != nil {
		return err

	} else {
		s.Path = tempObject.Path
		s.Set = tempObject.Set
		s.Labels = tempObject.Labels
		s.Layers = tempObject.Layers
		s.Generate = tempObject.Generate

		switch packages := tempObject.Package.(type) {
		case string:
			s.SetPackageFromString(packages)
		case []interface{}:
			for _, _package := range packages {
				s.AppendPackage(fmt.Sprint(_package))
			}
		}

	}

	if _string.IsEmpty(string(s.Generate)) {
		s.Generate = structs.ChangeTrackers.OnChange
	}
	return nil
}
