package structs

import (
	"parsdevkit.net/application/models/label"
	applicationResource "parsdevkit.net/application/structs/resource"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"

	"parsdevkit.net/pkg/errors"
	_string "parsdevkit.net/pkg/utilities/string"
)

type ResourceSpecification struct {
	applicationResource.ResourceIdentifier
	Path            string
	Set             string
	WorkspaceObject applicationWorkspace.WorkspaceIdentifier
	Labels          []label.Label
	Layers          []Layer
	Dictionary      []Dictionary
	Groups          []Group
}

func NewResourceSpecification(id int, name, workspace, path, set string, labels []label.Label, layers []Layer, workspaceObject applicationWorkspace.WorkspaceIdentifier) ResourceSpecification {
	return ResourceSpecification{
		ResourceIdentifier: applicationResource.NewResourceIdentifier(id, name, workspace),
		WorkspaceObject:    workspaceObject,
		Path:               path,
		Set:                set,
		Labels:             labels,
		Layers:             layers,
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

func (s *ResourceSpecification) IsPathExists() bool {
	return !_string.IsEmpty(s.Path)
}

func (s *ResourceSpecification) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		applicationResource.ResourceIdentifier
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
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

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
					// if _, ok := err.(*yaml.TypeError); !ok {
					// 	return err
					// }
					return err

				}
				s.Layers = append(s.Layers, value)

			}
		}
	}

	return nil
}
