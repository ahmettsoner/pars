package template

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

type TemplateSpecification struct {
	TemplateIdentifier
	schemas.SchemaSpecification
	Set             string
	Path            string
	Package         []string
	Output          Output
	Labels          []label.Label
	Layers          []layerPkg.Layer
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	Template        Template
	Generate        structs.ChangeTracker
}

func NewTemplateSpecification(id int, name, workspace, set string, path string, output Output, _package []string, labels []label.Label, layers []layerPkg.Layer, template Template, workspaceObject applicationWorkspace.WorkspaceIdentifier, generate structs.ChangeTracker) TemplateSpecification {
	return TemplateSpecification{
		TemplateIdentifier: NewTemplateIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Set:                set,
		Path:               path,
		Output:             output,
		Package:            _package,
		Labels:             labels,
		Layers:             layers,
		Template:           template,
		Generate:           generate,
	}
}
func (s TemplateSpecification) Validate() error {

	if _string.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	if _string.IsEmpty(s.Set) {
		return &errors.ErrFieldRequired{FieldName: "Set"}
	}

	if (s.Output == Output{}) {
		return &errors.ErrFieldRequired{FieldName: "Output"}
	}

	if (s.Template == Template{}) {
		return &errors.ErrFieldRequired{FieldName: "Template"}
	}
	return nil
}

func (s *TemplateSpecification) GetPackageString() string {
	return strings.Join(s.Package, "/")
}
func (s *TemplateSpecification) SetPackageFromString(_package string) {
	s.Package = strings.Split(_package, "/")
}
func (s *TemplateSpecification) AppendPackage(_package ...string) {
	s.Package = append(s.Package, _package...)
}
func (s *TemplateSpecification) IsPackageExists() bool {
	return len(s.Package) > 0
}
func (s *TemplateSpecification) IsPathExists() bool {
	return !_string.IsEmpty(s.Path)
}

func (s *TemplateSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		TemplateIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.TemplateIdentifier = tempIdentifierObject.TemplateIdentifier
	}

	var tempObject struct {
		Set      string                `yaml:"Set"`
		Path     string                `yaml:"Path"`
		Output   Output                `yaml:"Output"`
		Package  interface{}           `yaml:"Package"`
		Labels   []label.Label         `yaml:"Labels"`
		Layers   []layerPkg.Layer      `yaml:"Layers"`
		Template Template              `yaml:"Template"`
		Generate structs.ChangeTracker `yaml:"Generate"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Set = tempObject.Set
		s.Path = tempObject.Path
		s.Output = tempObject.Output
		s.Labels = tempObject.Labels
		s.Layers = tempObject.Layers
		s.Template = tempObject.Template
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
