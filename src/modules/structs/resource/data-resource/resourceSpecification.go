package dataresource

import (
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
	WorkspaceObject workspace.WorkspaceSpecification
	Labels          []label.Label
	Layers          []Layer
	Dictionary      []Dictionary
	Groups          []Group
	Data            any
}

func NewResourceSpecification(id int, name, workspace, path, set string, labels []label.Label, layers []Layer, data any, workspaceObject workspace.WorkspaceSpecification) ResourceSpecification {
	return ResourceSpecification{
		ResourceIdentifier: NewResourceIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Path:               path,
		Set:                set,
		Labels:             labels,
		Layers:             layers,
		Data:               data,
	}
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
		Path   string        `yaml:"Path"`
		Set    string        `yaml:"Set"`
		Labels []label.Label `yaml:"Labels"`
		Layers []interface{} `yaml:"Layers"`
		Data   any           `yaml:"Data"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Path = tempObject.Path
		s.Set = tempObject.Set
		s.Labels = tempObject.Labels

		for _, layer := range tempObject.Layers {
			switch layerType := layer.(type) {
			case string:
				s.Layers = append(s.Layers, NewLayer(0, layerType, []Section{}))
			case interface{}:
				var value Layer

				if err := unmarshal(&value); err != nil {
					if _, ok := err.(*yaml.TypeError); !ok {
						return err
					}
				}
				s.Layers = append(s.Layers, value)

			}
		}

		s.Data = tempObject.Data
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Resource.Name"}
	}

	if utils.IsEmpty(s.Set) {
		return &errors.ErrFieldRequired{FieldName: "Resource.Set"}
	}

	return nil
}
