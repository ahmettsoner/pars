package structs

import (
	applicationTemplate "parsdevkit.net/application/structs/template"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/pkg/errors"
)

type TemplateSpecification struct {
	applicationTemplate.TemplateIdentifier
	Template        Template
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
}

func NewTemplateSpecification(id int, name, workspace string, template Template, workspaceObject applicationWorkspace.WorkspaceIdentifier) TemplateSpecification {
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
