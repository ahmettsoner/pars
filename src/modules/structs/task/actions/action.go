package actions

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type ActionInterface interface {
	GetName() string
	GetType() string
	GetInputs() []string
	GetOutput() string
	GetSubtasks() map[string]ActionInterface
	GetData() map[string]interface{}
}

type Action struct {
	ActionInterface
	Type     string
	Inputs   []string
	Output   string
	Name     string
	SubTasks map[string]ActionInterface
	Retry    Retry
	Timeout  int
	Data     map[string]interface{}
}

func NewAction(name string) Action {
	return Action{
		Name:     name,
		SubTasks: make(map[string]ActionInterface),
	}
}

func (a Action) GetName() string                         { return a.Name }
func (a Action) GetType() string                         { return a.Type }
func (a Action) GetInputs() []string                     { return a.Inputs }
func (a Action) GetOutput() string                       { return a.Output }
func (a Action) GetSubtasks() map[string]ActionInterface { return a.SubTasks }
func (a Action) GetData() map[string]interface{}         { return a.Data }

func (s *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Name string `yaml:"Name"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Name = tempObject.Name
			}

		} else {
			return err
		}

	} else {
		s.Name = value
	}

	if utils.IsEmpty(s.Name) {
		return &errors.ErrFieldRequired{FieldName: "Name"}
	}

	return nil
}
