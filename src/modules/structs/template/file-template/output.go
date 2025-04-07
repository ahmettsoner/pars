package filetemplate

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

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

func (s *Output) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				File string `yaml:"File"`
			}

			if err := unmarshal(&tempObject); err != nil {
				if _, ok := err.(*yaml.TypeError); !ok {
					return err
				}
			} else {
				s.File = tempObject.File
			}

		} else {
			return err
		}

	} else {
		s.File = value
	}

	if utils.IsEmpty(string(s.File)) {
		return &errors.ErrFieldRequired{FieldName: "File"}
	}

	return nil
}
