package filetemplate

import (
	"fmt"
	"strings"

	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/workspace"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"parsdevkit.net/core/utils"
)

type TemplateSpecification struct {
	TemplateIdentifier
	Set             string
	Path            string
	Package         []string
	Output          Output
	Labels          []label.Label
	Layers          []Layer
	WorkspaceObject workspace.WorkspaceSpecification
	Template        Template
}

func NewTemplateSpecification(id int, name, workspace, set string, path string, output Output, _package []string, labels []label.Label, layers []Layer, template Template, workspaceObject workspace.WorkspaceSpecification) TemplateSpecification {
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
	}
}
func (e TemplateSpecification) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.TemplateIdentifier.Name, v.Required),
		v.Field(&e.Set, v.Required),
		v.Field(&e.Output, v.Required),
		v.Field(&e.Template, v.Required),
	)
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
	return !utils.IsEmpty(s.Path)
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
