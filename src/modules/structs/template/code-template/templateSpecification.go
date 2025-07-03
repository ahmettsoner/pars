package codetemplate

import (
	"fmt"
	"strings"

	"parsdevkit.net/application/models/label"
	applicationTemplate "parsdevkit.net/application/structs/template"
	_string "parsdevkit.net/pkg/utilities/string"

	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/pkg/errors"
)

type TemplateSpecification struct {
	applicationTemplate.TemplateIdentifier
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	Set             string
	Path            string
	Package         []string
	Output          Output
	Labels          []label.Label
	Layers          []Layer
	Template        Template
}

func NewTemplateSpecification(id int, name, workspace, set string, path string, output Output, _package []string, labels []label.Label, layers []Layer, template Template, workspaceObject applicationWorkspace.WorkspaceIdentifier) TemplateSpecification {
	return TemplateSpecification{
		TemplateIdentifier: applicationTemplate.NewTemplateIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Set:                set,
		Path:               path,
		Output:             output,
		Package:            _package,
		Labels:             labels,
		Layers:             layers,
		Template:           template,
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
		applicationTemplate.TemplateIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {

		s.TemplateIdentifier = tempIdentifierObject.TemplateIdentifier
	}

	var tempObject struct {
		Set      string        `yaml:"Set"`
		Path     string        `yaml:"Path"`
		Output   Output        `yaml:"Output"`
		Package  interface{}   `yaml:"Package"`
		Labels   []label.Label `yaml:"Labels"`
		Layers   []Layer       `yaml:"Layers"`
		Template Template      `yaml:"Template"`
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

		switch packages := tempObject.Package.(type) {
		case string:
			s.SetPackageFromString(packages)
		case []interface{}:
			for _, _package := range packages {
				s.AppendPackage(fmt.Sprint(_package))
			}
		}
	}

	return nil
}
