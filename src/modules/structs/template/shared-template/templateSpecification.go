package sharedtemplate

import (
	applicationTemplate "parsdevkit.net/application/structs/template"
	"parsdevkit.net/core/errors"
	"parsdevkit.net/structs/workspace"
)

type TemplateSpecification struct {
	applicationTemplate.TemplateIdentifier
	Template        Template
	WorkspaceObject workspace.WorkspaceSpecification
}

func NewTemplateSpecification(id int, name, workspace string, template Template, workspaceObject workspace.WorkspaceSpecification) TemplateSpecification {
	return TemplateSpecification{
		TemplateIdentifier: applicationTemplate.NewTemplateIdentifier(0, name, workspace),
		WorkspaceObject:    workspaceObject,
		Template:           template,
	}
}
func (s TemplateSpecification) Validate() error {
	if (s.Template == Template{}) {
		return &errors.ErrFieldRequired{FieldName: "Template"}
	}
	return nil
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
		Template Template `yaml:"Template"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Template = tempObject.Template
	}

	return nil
}
