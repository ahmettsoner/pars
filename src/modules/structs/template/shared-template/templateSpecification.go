package sharedtemplate

import (
	"parsdevkit.net/core/errors"
	"parsdevkit.net/structs/workspace"

	"gopkg.in/yaml.v3"
)

type TemplateSpecification struct {
	TemplateIdentifier
	Template        Template
	WorkspaceObject workspace.WorkspaceSpecification
}

func NewTemplateSpecification(id int, name, workspace string, template Template, workspaceObject workspace.WorkspaceSpecification) TemplateSpecification {
	return TemplateSpecification{
		TemplateIdentifier: NewTemplateIdentifier(name, workspace),
		WorkspaceObject:    workspaceObject,
		Template:           template,
	}
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
		Template Template `yaml:"Template"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Template = tempObject.Template
	}

	if (s.Template == Template{}) {
		return &errors.ErrFieldRequired{FieldName: "Template"}
	}

	return nil
}
