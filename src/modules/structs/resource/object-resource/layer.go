package objectresource

import (
	"gopkg.in/yaml.v3"
	layerPkg "parsdevkit.net/structs/layer"
)

type Layer struct {
	layerPkg.LayerIdentifier
	Sections []Section `yaml:"Sections"`
}

func NewLayer(id int, name string, sections []Section) Layer {
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

	// var parsedData map[string]interface{}

	// if err := unmarshal(&parsedData); err != nil {
	// 	return err
	// } else {
	// 	sectionsData := parsedData["sections"]
	// 	var sections []Section
	// 	if sectionsArray, ok := sectionsData.([]interface{}); ok {
	// 		sections = make([]Section, len(sectionsArray))
	// 		for i, sectionData := range sectionsArray {
	// 			if err := yaml.Unmarshal([]byte(sectionData.(string)), &sections[i]); err != nil {
	// 				return err
	// 			}
	// 		}
	// 	} else {
	// 		sections = make([]Section, 1)
	// 		if err := yaml.Unmarshal([]byte(sectionsData.(string)), &sections[0]); err != nil {
	// 			return err
	// 		}
	// 	}
	// 	s.Sections = sections
	// }

	var tempObject struct {
		Sections []Section `yaml:"Sections"`
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
