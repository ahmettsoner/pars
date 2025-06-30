package applicationproject

import (
	"parsdevkit.net/application/models/label"
	applicationGroup "parsdevkit.net/application/structs/group"
	applicationProject "parsdevkit.net/application/structs/project"
	applicationWorkspace "parsdevkit.net/application/structs/workspace"
	"parsdevkit.net/core/schemas"

	"parsdevkit.net/structs/project"
)

type Configuration struct {
	Layers       []applicationProject.Layer
	Dependencies []applicationProject.Package
	References   []ProjectBaseStruct
	Options      []string
	Modules      []string
	Components   []string
	Patterns     []string
}

func NewConfiguration(layers []applicationProject.Layer, dependencies []applicationProject.Package, references []ProjectBaseStruct, options, modules, components, patterns []string) Configuration {
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
		Layers:       []applicationProject.Layer(nil),
		Dependencies: []applicationProject.Package(nil),
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

func (s *Configuration) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		Layers       []applicationProject.Layer             `yaml:"Layers"`       //Burda inline defination eklenmeli, "Persistence:Data:Repository, Persistence:Data:Entity, Persistence:Data:Migration" gibi
		Dependencies []applicationProject.Package           `yaml:"Dependencies"` //Burda inline defination eklenmeli, "gopkg.in/yaml.v3@v3.0.1, gopkg.in/gorm" gibi
		References   []applicationProject.ProjectIdentifier `yaml:"References"`   //Burda inline defination eklenmeli, Workspace::Group/Name formatında "pars::core/utils, pars::service/project" gibi
		Options      []string                               `yaml:"Options"`
		Modules      []string                               `yaml:"Modules"`
		Components   []string                               `yaml:"Components"`
		Patterns     []string                               `yaml:"Patterns"`
	}

	if err := unmarshal(&tempObject); err != nil {
		// if _, ok := err.(*yaml.TypeError); !ok {
		// 	return err
		// }
		return err

	} else {

		s.Layers = tempObject.Layers
		s.Dependencies = tempObject.Dependencies
		for _, ref := range tempObject.References {
			projSpec := NewProjectBaseStruct(
				schemas.NewSchemaHeader(schemas.StructTypes.Project, string(project.ProjectKinds.Application), ref.Name, schemas.Metadata{}),
				NewProjectSpecification(
					0,
					"",
					ref.Group,
					ref.Workspace,
					"",
					applicationGroup.GroupIdentifier{},
					"",
					[]string(nil),
					[]label.Label(nil),
					[]string(nil),
					applicationWorkspace.WorkspaceIdentifier{},
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
