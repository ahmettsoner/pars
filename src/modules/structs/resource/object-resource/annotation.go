package objectresource

import (
	"fmt"
	"strings"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type Annotation struct {
	Type      string
	Arguments []MethodArgument
}

func NewAnnotation(_type string, arguments []MethodArgument) Annotation {
	return Annotation{
		Type:      _type,
		Arguments: arguments,
	}
}

func (s *Annotation) IsTypeExists() bool {
	return !utils.IsEmpty(s.Type)
}

func (s *Annotation) GetAllArguments() []MethodArgument {

	projectArguments := []MethodArgument{}
	projectArguments = append(projectArguments, s.Arguments...)

	return projectArguments
}
func (s *Annotation) AppendArguments(arguments ...MethodArgument) {
	s.Arguments = append(s.Arguments, arguments...)
}
func (s *Annotation) SetArgumentFromString(values string) {
	for _, value := range strings.Split(values, ",") {
		s.AppendArguments(NewMethodArgument("", value))
	}
}

func (s *Annotation) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var parameters string
	if err := unmarshal(&parameters); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Type      string      `yaml:"Type"`
				Arguments interface{} `yaml:"Arguments"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.Type = tempObject.Type

				switch arguments := tempObject.Arguments.(type) {
				case string:
					s.SetArgumentFromString(arguments)
				case []interface{}:
					for _, argument := range arguments {
						switch arg := argument.(type) {
						case string:
							var arg2 MethodArgument
							err := yaml.Unmarshal([]byte(string(arg)), &arg2)
							if err != nil {
								return err
							}
							s.AppendArguments(arg2)
						case map[string]interface{}:
							name, _ := arg["Name"].(string)
							value, valueExists := arg["Value"].(string)

							if valueExists {
								newArg := NewMethodArgument(name, value)
								s.AppendArguments(newArg)
							} else {
								return fmt.Errorf("Invalid value for method argument")
							}
						}
					}
				}
			}

		} else {
			return err
		}
	}

	if utils.IsEmpty(s.Type) {
		return &errors.ErrFieldRequired{FieldName: "Annotation.Type"}
	}

	return nil
}
