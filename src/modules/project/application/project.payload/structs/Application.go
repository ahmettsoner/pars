package structs

import (
	"parsdevkit.net/models"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/pkg/errors"
)

type Application struct {
	ProjectType models.ProjectType
	Runtime     Runtime
	Language    Language
	Schema      Schema
	Options     []string
	Modules     []string
	Components  []string
	Patterns    []string
	EntryPoint  string
	Ports       []string
}

/*
Framework: React
HealthCheck:

	Path: /health
	Interval: 30s
*/
func (s Application) Validate() error {

	if _string.IsEmpty(string(s.ProjectType)) || s.ProjectType.String() == "Unknown" {
		return &errors.ErrFieldRequired{FieldName: "ProjectType"}
	}
	return nil
}

func NewApplication(
	projectType models.ProjectType,
	runtime Runtime,
	schema Schema,
	options,
	modules,
	components,
	patterns []string,
	entryPoint string,
	ports []string,
) Application {
	return Application{
		ProjectType: projectType,
		Runtime:     runtime,
		Schema:      schema,
		Options:     options,
		Modules:     modules,
		Components:  components,
		Patterns:    patterns,
		EntryPoint:  entryPoint,
		Ports:       ports,
	}
}

func (s *Application) UnmarshalYAML(unmarshal func(interface{}) error) error {

	var tempObject struct {
		ProjectType models.ProjectType `yaml:"ProjectType"`
		Runtime     Runtime            `yaml:"Runtime"`
		Language    Language           `yaml:"Language"`
		Schema      Schema             `yaml:"Schema"`
		Options     []string           `yaml:"Options"`
		Modules     []string           `yaml:"Modules"`
		Components  []string           `yaml:"Components"`
		Patterns    []string           `yaml:"Patterns"`
		EntryPoint  string             `yaml:"EntryPoint"`
		Ports       []string           `yaml:"Ports"`
	}

	err := unmarshal(&tempObject)
	if err != nil {
		return err
	}

	s.ProjectType = tempObject.ProjectType
	s.Runtime = tempObject.Runtime
	s.Language = tempObject.Language
	s.Schema = tempObject.Schema
	s.Options = tempObject.Options
	s.Modules = tempObject.Modules
	s.Components = tempObject.Components
	s.Patterns = tempObject.Patterns
	s.EntryPoint = tempObject.EntryPoint
	s.Ports = tempObject.Ports

	return nil
}
