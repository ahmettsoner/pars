package structs

import (
	layerPkg "parsdevkit.net/application/models/layer"
	templateStruct "parsdevkit.net/application/structs/template"
)

type Layer struct {
	layerPkg.LayerIdentifier
	Sections []templateStruct.Section `yaml:"Sections"`
}

func NewLayer(id int, name string, sections []templateStruct.Section) Layer {
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
		Sections []templateStruct.Section `yaml:"Sections"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {
		s.Sections = tempObject.Sections
	}

	return nil
}
