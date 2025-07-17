package layer

import (
	"gopkg.in/yaml.v3"
	sectionPkg "parsdevkit.net/application/models/section"
)

type Layer struct {
	LayerIdentifier
	Sections []sectionPkg.Section `yaml:"Sections"`
}

func NewLayer(id int, name string, sections []sectionPkg.Section) Layer {
	return Layer{
		LayerIdentifier: NewLayerIdentifier(id, name),
		Sections:        sections,
	}
}

func (s *Layer) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempIdentifierObject struct {
		LayerIdentifier
	}

	if err := unmarshal(&tempIdentifierObject); err != nil {
		return err
	} else {
		s.LayerIdentifier = tempIdentifierObject.LayerIdentifier
	}

	var tempObject struct {
		Sections []sectionPkg.Section `yaml:"Sections"`
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
