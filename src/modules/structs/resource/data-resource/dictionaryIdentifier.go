package dataresource

import (
	"parsdevkit.net/core/utils"

	"parsdevkit.net/core/errors"

	"gopkg.in/yaml.v3"
)

type DictionaryIdentifier struct {
	Key string
}

func NewDictionaryIdentifier(key string) DictionaryIdentifier {
	return DictionaryIdentifier{
		Key: key,
	}
}

func (s *DictionaryIdentifier) IsKeyExists() bool {
	return !utils.IsEmpty(s.Key)
}

func (s *DictionaryIdentifier) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		if _, ok := err.(*yaml.TypeError); ok {
			var tempObject struct {
				Key string `yaml:"Key"`
			}

			err := unmarshal(&tempObject)
			if err != nil {
				return err
			}

			s.Key = tempObject.Key
		} else {
			return err
		}

	} else {
		s.Key = value
	}

	if utils.IsEmpty(s.Key) {
		return &errors.ErrFieldRequired{FieldName: "Dictionary.Key"}
	}

	return nil
}
