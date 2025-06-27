package dataresource

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/yaml.v3"
)

type ResourceIdentifier struct {
	ID        int
	Name      string
	Workspace string
}

func NewResourceIdentifier(id int, name string, workspace string) ResourceIdentifier {
	return ResourceIdentifier{
		ID:        id,
		Name:      name,
		Workspace: workspace,
	}
}
func (e ResourceIdentifier) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.Name, v.Required),
	)
}

func (s *ResourceIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name      string `yaml:"Name"`
				Workspace string `yaml:"Workspace"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Name = tempObject.Name
			s.Workspace = tempObject.Workspace
		} else {
			return err
		}

	} else {
		s.Name = value
	}

	return nil
}
