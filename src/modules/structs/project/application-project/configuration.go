package applicationproject

import (
	"encoding/json"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"
	"parsdevkit.net/structs/label"
	"parsdevkit.net/structs/project"
	"parsdevkit.net/structs/workspace"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Layers       []Layer
	Dependencies []Package
	References   []ProjectBaseStruct
	Options      []string
	Modules      []string
	Components   []string
	Patterns     []string
}

func NewConfiguration(layers []Layer, dependencies []Package, references []ProjectBaseStruct, options, modules, components, patterns []string) Configuration {
	return Configuration{
		Layers:       layers,
		Dependencies: dependencies,
		References:   references,
		Options:      options,
		Modules:      modules,
		Components:   components,
		Patterns:     patterns,
	}
}

func NewConfiguration_Empty() Configuration {
	return Configuration{
		Layers:       []Layer(nil),
		Dependencies: []Package(nil),
		References:   []ProjectBaseStruct(nil),
		Options:      []string(nil),
		Modules:      []string(nil),
		Components:   []string(nil),
		Patterns:     []string(nil),
	}
}

func (s *Configuration) AppendReferences(references ...ProjectBaseStruct) {
	s.References = append(s.References, references...)
}

func (s *Configuration) UnmarshalJSON(data []byte) error {

	var tempObject struct {
		Layers       []Layer
		Dependencies []Package
		References   []ProjectBaseStruct
		Options      []string
		Modules      []string
		Components   []string
		Patterns     []string
	}

	err := json.Unmarshal(data, &tempObject)
	if err != nil {
		return err
	}

	s.Layers = tempObject.Layers
	s.Dependencies = tempObject.Dependencies
	s.References = tempObject.References
	s.Options = tempObject.Options
	s.Modules = tempObject.Modules
	s.Components = tempObject.Components
	s.Patterns = tempObject.Patterns

	return nil
}

func (s *Configuration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Layers       []Layer             `yaml:"Layers"`       //Burda inline defination eklenmeli, "Persistence:Data:Repository, Persistence:Data:Entity, Persistence:Data:Migration" gibi
		Dependencies []Package           `yaml:"Dependencies"` //Burda inline defination eklenmeli, "gopkg.in/yaml.v3@v3.0.1, gopkg.in/gorm" gibi
		References   []ProjectIdentifier `yaml:"References"`   //Burda inline defination eklenmeli, Workspace::Group/Name formatında "pars::core/utils, pars::service/project" gibi
		Options      []string            `yaml:"Options"`
		Modules      []string            `yaml:"Modules"`
		Components   []string            `yaml:"Components"`
		Patterns     []string            `yaml:"Patterns"`
	}

	if err := unmarshal(&tempObject); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	} else {

		s.Layers = tempObject.Layers
		s.Dependencies = tempObject.Dependencies
		for _, ref := range tempObject.References {
			projSpec := NewProjectBaseStruct(
				project.NewHeader(structs.StructTypes.Project, project.StructKinds.Application, ref.Name, structs.Metadata{}),
				NewProjectSpecification(
					0,
					"",
					ref.Group,
					ref.Workspace,
					"",
					group.GroupSpecification{},
					"",
					[]string(nil),
					[]label.Label(nil),
					[]string(nil),
					workspace.WorkspaceSpecification{},
					Platform{},
					Runtime{},
					Schema{},
					Configuration{},
				),
			)
			s.AppendReferences(projSpec)
		}

		s.Options = tempObject.Options
		s.Modules = tempObject.Modules
		s.Components = tempObject.Components
		s.Patterns = tempObject.Patterns
	}

	return nil
}
