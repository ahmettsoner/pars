package codetemplate

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"gopkg.in/yaml.v3"
)

type Output struct {
	File string
}

func NewOutput(file string) Output {
	return Output{
		File: file,
	}
}
func (e Output) Validate() error {
	return v.ValidateStruct(&e,
		v.Field(&e.File, v.Required),
	)
}

func (s *Output) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				File string `yaml:"File"`
			}

			if err := unmarshal(&tempObject); err != nil {
				// if _, ok := err.(*yaml.TypeError); !ok {
				// 	return err
				// }
				return err

			} else {
				s.File = tempObject.File
			}

		} else {
			return err
		}

	} else {
		s.File = value
	}

	return nil
}
