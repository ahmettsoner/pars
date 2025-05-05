package codetemplate

import (
	"gopkg.in/yaml.v3"
	layerPkg "parsdevkit.net/structs/layer"
	"parsdevkit.net/structs/template"
)

type Layer struct {
	layerPkg.LayerIdentifier
	Sections []template.Section `yaml:"Sections"`
}

func NewLayer(id int, name string, sections []template.Section) Layer {
	return Layer{
		LayerIdentifier: layerPkg.NewLayerIdentifier(id, name),
		Sections:        sections,
	}
}

func (s *Layer) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		layerPkg.LayerIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {
		s.LayerIdentifier = tempIdentifierObject.LayerIdentifier
	}

	var tempObject struct {
		Sections []template.Section `yaml:"Sections"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {
		s.Sections = tempObject.Sections
	}

	return nil
}
