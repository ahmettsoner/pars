package sharedtemplate

import (
	"parsdevkit.net/structs/workspace"

	v "github.com/go-ozzo/ozzo-validation/v4"
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
func (e TemplateSpecification) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.TemplateIdentifier.Name, v.Required),
		v.Field(&e.Template, v.Required),
	)
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
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Template = tempObject.Template
	}

	return nil
}
